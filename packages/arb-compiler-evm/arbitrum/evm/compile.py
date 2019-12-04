# Copyright 2019, Offchain Labs, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import pyevmasm

from ..annotation import modifies_stack
from ..std import stack_manip
from ..std import byterange, bitwise
from ..std import bst
from ..vm import AVMOp
from ..ast import AVMLabel
from ..ast import BlockStatement
from ..compiler import compile_block
from .. import value

from . import os, call_frame
from . import execution

import eth_utils


class EVMNotSupported(Exception):
    """VM tried to run opcode that blocks"""

    pass


def replace_self_balance(instrs):
    out = []
    i = 0
    while i < len(instrs):
        if (
            instrs[i].name == "ADDRESS"
            and instrs[i + 1].name == "PUSH20"
            and instrs[i + 1].operand == (2 ** 160 - 1)
            and instrs[i + 2].name == "AND"
            and instrs[i + 3].name == "BALANCE"
        ):
            out.append(AVMOp("SELF_BALANCE"))
            i += 4
        else:
            out.append(instrs[i])
            i += 1
    return out


def contains_endcode(val):
    if not isinstance(val, int):
        return False
    while val >= 0xA265:
        if (val & 0xFFFF == 0xA165) or (val & 0xFFFF == 0xA265):
            return True
        val = val // 256
    return False


def remove_metadata(instrs):
    for i in range(len(instrs) - 2, -1, -1):
        first_byte = instrs[i].opcode
        second_byte = instrs[i + 1].opcode
        if ((first_byte == 0xA1) or (first_byte == 0xA2)) and second_byte == 0x65:
            return instrs[:i]
        if hasattr(instrs[i], "operand") and contains_endcode(instrs[i].operand):
            # leave the boundary instruction in--it's probably valid and minimal testing suggests this reduces errors
            return instrs[: i + 1]
    return instrs


def generate_evm_code(raw_code, storage):
    contracts = {}
    for contract in raw_code:
        contracts[contract] = list(pyevmasm.disassemble_all(raw_code[contract]))

    code_tuples_data = {}
    for contract in raw_code:
        code_tuples_data[contract] = byterange.frombytes(
            bytes.fromhex(raw_code[contract].hex())
        )
    code_tuples_func = bst.make_static_lookup(code_tuples_data)

    @modifies_stack([value.IntType()], 1)
    def code_tuples(vm):
        code_tuples_func(vm)

    code_hashes_data = {}
    for contract in raw_code:
        code_hashes_data[contract] = int.from_bytes(
            eth_utils.crypto.keccak(raw_code[contract]), byteorder="big"
        )
    code_hashes_func = bst.make_static_lookup(code_hashes_data)

    @modifies_stack([value.IntType()], [value.ValueType()])
    def code_hashes(vm):
        code_hashes_func(vm)

    code_sizes = {}
    for contract in contracts:
        code_sizes[contract] = len(contracts[contract]) + sum(
            op.operand_size for op in contracts[contract]
        )

    # Give the interrupt contract address a nonzero size
    code_sizes[0x01] = 1
    code_size_func = bst.make_static_lookup(code_sizes)

    @modifies_stack([value.IntType()], 1)
    def code_size(vm):
        code_size_func(vm)

    impls = []
    contract_info = []
    for contract in sorted(contracts):
        if contract not in storage:
            storage[contract] = {}
        impls.append(
            generate_contract_code(
                AVMLabel("contract_entry_" + str(contract)),
                contracts[contract],
                code_tuples_data[contract],
                contract,
                code_size,
                code_tuples,
                code_hashes,
            )
        )
        contract_info.append(
            {
                "code_point": AVMLabel("contract_entry_" + str(contract)),
                "contractID": contract,
                "storage": storage[contract],
            }
        )

    def initialization(vm):
        os.initialize(vm, contract_info)
        vm.jump_direct(AVMLabel("run_loop_start"))

    def run_loop_start(vm):
        vm.set_label(AVMLabel("run_loop_start"))
        os.get_next_message(vm)
        execution.setup_initial_call(vm)
        vm.push(AVMLabel("run_loop_start"))
        vm.jump()

    main_code = []
    main_code.append(compile_block(run_loop_start))
    main_code += impls
    return compile_block(initialization), BlockStatement(main_code)


def evm_div(vm):
    vm.dup1()
    vm.iszero()
    vm.ifelse(lambda vm: vm.pop(), lambda vm: vm.div())


def evm_sdiv(vm):
    vm.dup1()
    vm.iszero()
    vm.ifelse(lambda vm: vm.pop(), lambda vm: vm.sdiv())


def evm_mod(vm):
    vm.dup1()
    vm.iszero()
    vm.ifelse(lambda vm: vm.pop(), lambda vm: vm.mod())


def evm_smod(vm):
    vm.dup1()
    vm.iszero()
    vm.ifelse(lambda vm: vm.pop(), lambda vm: vm.smod())


def evm_addmod(vm):
    vm.dup2()
    vm.iszero()
    vm.ifelse(lambda vm: [vm.pop(), vm.pop()], lambda vm: vm.addmod())


def evm_mulmod(vm):
    vm.dup2()
    vm.iszero()
    vm.ifelse(lambda vm: [vm.pop(), vm.pop()], lambda vm: vm.mulmod())


def not_supported_op(name):
    raise EVMNotSupported(name)


EVM_STATIC_OPS = {
    "SELF_BALANCE": lambda vm: [vm.push(0), os.balance_get(vm)],
    # 0s: Stop and Arithmetic Operations
    "STOP": execution.stop,
    "ADD": lambda vm: vm.add(),
    "MUL": lambda vm: vm.mul(),
    "SUB": lambda vm: vm.sub(),
    "DIV": evm_div,
    "SDIV": evm_sdiv,
    "MOD": evm_mod,
    "SMOD": evm_smod,
    "ADDMOD": evm_addmod,
    "MULMOD": evm_mulmod,
    "EXP": lambda vm: vm.exp(),
    "SIGNEXTEND": lambda vm: vm.signextend(),
    # 10s: Comparison & Bitwise Logic Operations
    "LT": lambda vm: vm.lt(),
    "GT": lambda vm: vm.gt(),
    "SLT": lambda vm: vm.slt(),
    "SGT": lambda vm: vm.sgt(),
    "EQ": lambda vm: vm.eq(),
    "ISZERO": lambda vm: vm.iszero(),
    "AND": lambda vm: vm.bitwise_and(),
    "OR": lambda vm: vm.bitwise_or(),
    "XOR": lambda vm: vm.bitwise_xor(),
    "NOT": lambda vm: vm.bitwise_not(),
    "BYTE": lambda vm: vm.byte(),
    "SHL": lambda vm: [vm.swap1(), bitwise.shift_left(vm)],
    "SHR": lambda vm: [vm.swap1(), bitwise.shift_right(vm)],
    "SAR": lambda vm: [vm.swap1(), bitwise.arithmetic_shift_right(vm)],
    # 20s: SHA3
    "SHA3": os.evm_sha3,
    # 30s: Environmental Information
    "ADDRESS": lambda vm: [
        os.get_call_frame(vm),
        call_frame.call_frame.get("contractID")(vm),
    ],
    "ORIGIN": os.message_origin,
    "CALLER": os.message_caller,
    "CALLVALUE": os.message_value,
    "CALLDATALOAD": os.message_data_load,
    "CALLDATASIZE": os.message_data_size,
    "CALLDATACOPY": os.message_data_copy,
    # TODO: Arbitrary value
    "GASPRICE": lambda vm: vm.push(1),
    "RETURNDATASIZE": os.return_data_size,
    "RETURNDATACOPY": os.return_data_copy,
    # 40s: Block Information
    "BLOCKHASH": lambda vm: not_supported_op("BLOCKHASH"),
    "COINBASE": lambda vm: not_supported_op("COINBASE"),
    "TIMESTAMP": os.get_timestamp,
    "NUMBER": os.get_block_number,
    "DIFFICULTY": lambda vm: not_supported_op("DIFFICULTY"),
    "GASLIMIT": lambda vm: vm.push(10000000000),
    # 50s: Stack, Memory, Storage and Flow Operations
    "POP": lambda vm: vm.pop(),
    "MLOAD": os.memory_load,
    "MSTORE": os.memory_store,
    "MSTORE8": os.memory_store8,
    "SLOAD": os.storage_load,
    "SSTORE": os.storage_store,
    "GETPC": lambda vm: not_supported_op("GETPC"),
    "MSIZE": os.memory_length,
    # TODO: Fill in here
    "GAS": lambda vm: vm.push(9999999999),
    "DUP1": stack_manip.dup_n(0),
    "DUP2": stack_manip.dup_n(1),
    "DUP3": stack_manip.dup_n(2),
    "DUP4": stack_manip.dup_n(3),
    "DUP5": stack_manip.dup_n(4),
    "DUP6": stack_manip.dup_n(5),
    "DUP7": stack_manip.dup_n(6),
    "DUP8": stack_manip.dup_n(7),
    "DUP9": stack_manip.dup_n(8),
    "DUP10": stack_manip.dup_n(9),
    "DUP11": stack_manip.dup_n(10),
    "DUP12": stack_manip.dup_n(11),
    "DUP13": stack_manip.dup_n(12),
    "DUP14": stack_manip.dup_n(13),
    "DUP15": stack_manip.dup_n(14),
    "DUP16": stack_manip.dup_n(15),
    "SWAP1": stack_manip.swap_n(1),
    "SWAP2": stack_manip.swap_n(2),
    "SWAP3": stack_manip.swap_n(3),
    "SWAP4": stack_manip.swap_n(4),
    "SWAP5": stack_manip.swap_n(5),
    "SWAP6": stack_manip.swap_n(6),
    "SWAP7": stack_manip.swap_n(7),
    "SWAP8": stack_manip.swap_n(8),
    "SWAP9": stack_manip.swap_n(9),
    "SWAP10": stack_manip.swap_n(10),
    "SWAP11": stack_manip.swap_n(11),
    "SWAP12": stack_manip.swap_n(12),
    "SWAP13": stack_manip.swap_n(13),
    "SWAP14": stack_manip.swap_n(14),
    "SWAP15": stack_manip.swap_n(15),
    "SWAP16": stack_manip.swap_n(16),
    # a0s: Logging Operations
    "LOG1": os.evm_log1,
    "LOG2": os.evm_log2,
    "LOG3": os.evm_log3,
    "LOG4": os.evm_log4,
    # f0s: System operations
    "CREATE": lambda vm: not_supported_op("CREATE"),
    "CREATE2": lambda vm: not_supported_op("CREATE2"),
    "RETURN": execution.ret,
    "REVERT": execution.revert,
    "SELFDESTRUCT": execution.selfdestruct,
    "BALANCE": lambda vm: [
        print("Warning: BALANCE was used which may lead to an error"),
        vm.push(0),
        vm.swap1(),
        os.ext_balance(vm),
    ],
}

UNHANDLED_OPCODE = {
    0x3F: "EXTCODEHASH",
    0xF5: "CREATE2",
    0x1B: "SHL",
    0x1C: "SHR",
    0x1D: "SAR",
}


def get_opcode_name(instr):
    if instr.opcode in UNHANDLED_OPCODE:
        return UNHANDLED_OPCODE[instr.opcode]
    if instr.name[:4] == "PUSH":
        return "PUSH"
    return instr.name


def generate_contract_code(
    label, code, code_tuple, contract_id, code_size, code_tuples, code_hashes
):
    code = remove_metadata(code)
    code = replace_self_balance(code)

    jump_table = {}
    for insn in code:
        if insn.name == "JUMPDEST":
            jump_table[insn.pc] = AVMLabel(
                "jumpdest_{}_{}".format(contract_id, insn.pc)
            )
    dispatch_func = bst.make_static_lookup(jump_table)

    @modifies_stack([value.IntType()], [value.ValueType()], contract_id)
    def dispatch(vm):
        dispatch_func(vm)

    @modifies_stack(0, 1, contract_id)
    def get_contract_code(vm):
        vm.push(code_tuple)

    def evm_extcodesize(vm):
        print("Warning: EXTCODESIZE was used which may lead to an error")
        code_size(vm)
        vm.dup0()
        vm.tnewn(0)
        vm.eq()
        vm.ifelse(lambda vm: [vm.error()])

    def evm_extcodecopy(vm):
        print("Warning: EXTCODECOPY was used which may lead to an error")
        code_tuples(vm)
        vm.dup0()
        vm.tnewn(0)
        vm.eq()
        vm.ifelse(lambda vm: [vm.error()])
        os.set_scratch(vm)
        os.evm_copy_to_memory(vm, os.get_scratch)

    def evm_extcodehash(vm):
        print("Warning: EXTCODEHASH was used which may lead to an error")
        code_hashes(vm)
        vm.dup0()
        vm.tnewn(0)
        vm.eq()
        vm.ifelse(lambda vm: [vm.error()])

    EVM_CONTRACT_OPS = {
        "JUMP": lambda vm: [
            dispatch(vm),
            vm.dup0(),
            vm.tnewn(0),
            vm.eq(),
            vm.ifelse(lambda vm: [vm.error()], lambda vm: [vm.jump()]),
        ],
        "JUMPI": lambda vm: [
            dispatch(vm),
            vm.dup0(),
            vm.tnewn(0),
            vm.eq(),
            vm.ifelse(lambda vm: [vm.error()], lambda vm: [vm.cjump()]),
        ],
        "CODESIZE": lambda vm: vm.push(len(code)),
        "CODECOPY": lambda vm: os.evm_copy_to_memory(vm, get_contract_code),
        "EXTCODESIZE": evm_extcodesize,
        "EXTCODECOPY": evm_extcodecopy,
        "EXTCODEHASH": evm_extcodehash,
    }

    def run_op(instr):
        def evm_invalid_op(vm):
            if instr.opcode != 0xFE:
                print(
                    "Warning: Source code contained nonstandard invalid opcode {}".format(
                        hex(instr.opcode)
                    )
                )
            execution.revert(vm)

        evm_instr_ops = {
            "PUSH": lambda vm: vm.push(instr.operand),
            "JUMPDEST": lambda vm: vm.set_label(
                AVMLabel("jumpdest_{}_{}".format(contract_id, instr.pc))
            ),
            "CALL": lambda vm: execution.call(vm, instr.pc, contract_id),
            "CALLCODE": lambda vm: execution.call(vm, instr.pc, contract_id),
            "DELEGATECALL": lambda vm: execution.delegatecall(
                vm, instr.pc, contract_id
            ),
            "STATICCALL": lambda vm: execution.staticcall(vm, instr.pc, contract_id),
            "INVALID": evm_invalid_op,
        }

        instr_name = get_opcode_name(instr)

        if instr_name in EVM_STATIC_OPS:
            return EVM_STATIC_OPS[instr_name]

        if instr_name in EVM_CONTRACT_OPS:
            return EVM_CONTRACT_OPS[instr_name]

        if instr_name in evm_instr_ops:
            return evm_instr_ops[instr_name]

        raise Exception("Unhandled instruction {}".format(instr))

    contract_code = [label]
    for insn in code:
        block = compile_block(run_op(insn))
        block.add_node("EthOp({}, {})".format(insn, insn.pc))
        contract_code.append(block)

    return BlockStatement(contract_code)
