class Matrix:
    def __init__(self, matrix_string):
        self.matrix = [
            list(map(lambda x: int(x), line.split(" ")))
            for line in matrix_string.splitlines()
        ]

    def row(self, index):
        return self.matrix[index - 1]

    def column(self, index):
        return [lines[index - 1] for lines in self.matrix]
