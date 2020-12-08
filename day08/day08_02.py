TEST_COMMANDS = """nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6"""


def read_input():
    with open('input.txt') as input_file:
        return input_file.read()


def assemble_tape(instructions):
    commands = []
    for command in instructions.split('\n'):
        if command:
            operation, argument = command.split()
            commands.append((operation, int(argument)))
    return commands


def run(instructions):
    ins_count = len(instructions)
    accumulator = 0
    executed = set()  # which commands already executed
    ins_ptr = 0
    completed = False
    while True:
        if ins_ptr in executed:
            break
        if ins_ptr == ins_count:
            completed = True
            break
        executed.add(ins_ptr)
        operation, argument = instructions[ins_ptr]
        if operation == 'acc':
            accumulator += argument
        elif operation == 'jmp':
            ins_ptr += argument - 1
        else:
            pass
        ins_ptr += 1
    return completed, accumulator
    

def sed(instructions, index=None):
    if not index:
        return instructions
    operation, argument = instructions[index]
    instructions[index] = ('jmp' if operation == 'nop' else 'nop', argument)
    return instructions


def test_tape(tape):
    opcode_indices = [i for i, (opcode, arg) in enumerate(tape) if opcode in ('nop', 'jmp')]
    for index in opcode_indices:
        tape_ = tape[:]
        opcode, _ = tape_[index]
        print(f'Modifying {opcode} at index {index}')
        tape_ = sed(tape_, index=index)
        completed, accumulator = run(tape_)
        if completed:
            print(f'Finished with ACCUMULATOR={accumulator}')
            break


tape = assemble_tape(TEST_COMMANDS)
test_tape(tape)

tape = assemble_tape(read_input())
test_tape(tape)
