import string


def is_pangram(sentence):
    alphabet = string.ascii_lowercase
    return all(letter in sentence.lower() for letter in alphabet)
