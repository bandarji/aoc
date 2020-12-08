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
    accumulator = 0
    executed = set()  # which commands already executed
    ins_ptr = 0
    while True:
        if ins_ptr in executed:
            break
        executed.add(ins_ptr)
        operation, argument = instructions[ins_ptr]
        if operation == 'acc':
            accumulator += argument
        elif operation == 'jmp':
            ins_ptr += argument - 1
        else:
            pass
        print(f'ins_ptr={ins_ptr}, operation={operation}, argument={argument}, accumulator={accumulator}, executed={executed}')
        ins_ptr += 1
    return accumulator

tape = assemble_tape(TEST_COMMANDS)
accumulator = run(tape)
print(f'ACCUMULATOR = {accumulator}\n')

tape = assemble_tape(read_input())
accumulator = run(tape)
print(f'ACCUMULATOR = {accumulator}\n')
