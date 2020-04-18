def decode(string):
    import re

    if len(string) == 0:
        return string
    parse_string = re.sub(r"(\d\d?)([a-zA-Z\s])", "\\1*\\2", string)
    formula = re.findall(r"\d{1,2}\*.|.", parse_string)
    secret = ""
    for item in formula:
        partial = item.split("*")
        secret += partial[1] * int(partial[0]) if len(partial) > 1 else partial[0]
    return secret


def encode(string):
    if len(string) == 0:
        return string

    count = 0
    last_letter = string[0]
    result = []
    for letter in string + "$":
        if last_letter is letter:
            count += 1
        else:
            result.append((count, last_letter))
            count = 1
        last_letter = letter
    secret = ""
    for tuple in result:
        secret += f"{tuple[0]}{tuple[1]}" if tuple[0] > 1 else tuple[1]
    return secret
