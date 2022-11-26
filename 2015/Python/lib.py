def get_input(filename: str) -> str:
    with open(filename) as file:
        return file.read()
