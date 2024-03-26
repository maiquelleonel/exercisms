package protein

import (
	"errors"
)

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("invalid base")

func getCodons(rna string) *[]string {
	var cod string
	var ret []string
	for _, char := range rna {
		cod += string(char)
		if len(cod) == 3 {
			ret = append(ret, cod)
			cod = ""
		}
	}
	return &ret
}

func FromRNA(rna string) ([]string, error) {
	var ret []string
	codons := *getCodons(rna)
	for _, codon := range codons {
		if protein, err := FromCodon(codon); err == nil {
			ret = append(ret, protein)
		} else if err == ErrInvalidBase {
			return nil, err
		} else {
			break
		}
	}
	return ret, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase

}
