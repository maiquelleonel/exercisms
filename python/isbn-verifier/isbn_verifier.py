import re


def is_valid(isbn):
    clean_isbn = re.sub(r"-|[^0-9X]", "", isbn.upper())
    if len(clean_isbn) != 10:
        return False
    multipliers = range(10, 0, -1)
    calc = [
        (int(char) * multipliers[i]) if char != "X" else 10
        for i, char in enumerate(clean_isbn)
    ]
    return sum(calc) % 11 == 0
