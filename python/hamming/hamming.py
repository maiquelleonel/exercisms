def distance(strand_a, strand_b):
    diff = 0
    if len(strand_a) != len(strand_b):
        raise ValueError("Sequence length must be equal")
    sequence_a = list(strand_a)
    sequence_b = list(strand_b)
    for index, value in enumerate(sequence_a):
        if value != sequence_b[index]:
            diff += 1
    return diff
