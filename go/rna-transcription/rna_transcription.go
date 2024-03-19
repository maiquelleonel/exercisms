package strand

func ToRNA(dna string) string {
	var ret string
	for _, enzim := range dna {
		switch enzim {
		case 'G':
			ret += "C"
		case 'C':
			ret += "G"
		case 'T':
			ret += "A"
		case 'A':
			ret += "U"
		}
	}
	return ret
}
