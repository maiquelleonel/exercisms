class School:
    def __init__(self):
        self.grades = {}

    def _validates_presence_of(self, name):
        if name in self.roster():
            raise ValueError("You cannot change the grade lunched")

    def add_student(self, name, grade):
        self._validates_presence_of(name)
        if grade in self.grades.keys():
            self.grades[grade].append(name)
            self.grades[grade].sort()
        else:
            self.grades[grade] = [name]

    def roster(self):
        if len(self.grades):
            return [
                student
                for grade in sorted(self.grades.keys())
                for student in self.grades[grade]
            ]

        return []

    def grade(self, grade_number=None):
        if not grade_number is None:
            return self.grades.get(grade_number, [])
        return self.roster()
