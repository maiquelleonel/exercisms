def to_rna(dna_strand):
    dna_transcript = {"G": "C", "C": "G", "T": "A", "A": "U"}
    return "".join([dna_transcript[enzim] for enzim in dna_strand])
