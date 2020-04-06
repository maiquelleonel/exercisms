import re


def abbreviate(words):
    initials = [word[0:1] for word in re.split("\\W\\W?", words) if word]
    return "".join(initials).upper()
