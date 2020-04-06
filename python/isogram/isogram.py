import re


def is_isogram(string):
    letters = re.sub(r"\W", "", string).lower()
    ret = True
    for letter in letters:
        if letters.count(letter) > 1:
            ret = False
            break
    return ret
