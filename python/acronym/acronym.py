import re


def abbreviate(words):
    initials = [word[0:1] for word in re.split(r'\s_?|\s-?\s?|-', words) if word]
    return "".join(initials).upper()
