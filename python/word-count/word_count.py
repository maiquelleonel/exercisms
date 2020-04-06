import re


def count_words(sentence):
    unapostrify = re.sub(r"(\w)(')(\w)", "\\1RR\\3", sentence.lower())
    special_chars_striped = re.sub(r"\W\W?|_+", "*", unapostrify)
    reapostrify = special_chars_striped.replace("RR", "'")
    words = reapostrify.split("*")
    return {word: words.count(word) for word in words if word}
