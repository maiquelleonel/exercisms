import re


def response(hey_bob):
    phrase = hey_bob.strip()
    if len(phrase) > 0:
        is_upper = phrase.isupper()
        is_question = phrase.endswith("?")
        if is_upper and is_question:
            return "Calm down, I know what I'm doing!"
        elif is_question:
            return "Sure."
        elif is_upper:
            return "Whoa, chill out!"
        elif re.findall(r"\w+", phrase):
            return "Whatever."

    return "Fine. Be that way!"
