STUDENTS = [
    "Alice",
    "Bob",
    "Charlie",
    "David",
    "Eve",
    "Fred",
    "Ginny",
    "Harriet",
    "Ileana",
    "Joseph",
    "Kincaid",
    "Larry",
]


class Garden:
    PLANTS = {"V": "Violets", "R": "Radishes", "G": "Grass", "C": "Clover"}
    _plants = []

    def __init__(self, diagram, students=STUDENTS):
        self._plants = []
        self.size = 2
        self.diagram = diagram
        self.students = sorted(students)

    def _draw(self):
        for line in list(self.diagram.splitlines()):
            self._plants.append(
                [line[i : i + self.size] for i in range(0, len(line), self.size)]
            )

    def plants(self, student):
        self._draw()
        student_plants = "".join(
            self._plants[0][self.students.index(student)]
            + self._plants[1][self.students.index(student)]
        )
        return [self.PLANTS[letter] for letter in student_plants]
