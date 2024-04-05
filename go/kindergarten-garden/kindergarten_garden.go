package kindergarten

import (
	"errors"
	"slices"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	diagram  string
	students []string
	plants   [][]string
}

var PLANTS map[string]string = map[string]string{
	"V": "violets",
	"R": "radishes",
	"G": "grass",
	"C": "clover",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	var plants [][]string = make([][]string, 2)
	var students []string = slices.Clone(children)
	slices.Sort(students)

	if diagram[0] != '\n' {
		return &Garden{}, errors.New("invalid diagram")
	}

	lines := slices.Delete(strings.Split(diagram, "\n"), 0, 1)

	if len(lines[0]) > len(lines[1]) {
		return &Garden{}, errors.New("mismatch row size")
	}

	if len(children) > len(slices.Compact(children)) {
		return &Garden{}, errors.New("duplicate name")
	}

	for l, line := range lines {
		if strings.IndexAny(line, "VRGC") == -1 {
			return &Garden{}, errors.New("invalid cup code")
		}

		if !(len(line)%2 == 0) {
			return &Garden{}, errors.New("odd number of cups")
		}

		for i := 0; i <= len(line)-2; i += 2 {
			plants[l] = append(plants[l], string(line[i:i+2]))
		}
	}

	return &Garden{
		diagram:  diagram,
		students: students,
		plants:   plants,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if slices.Index(g.students, child) == -1 {
		return []string{}, false
	}
	var a, b string
	var ret []string
	a = g.plants[0][slices.Index(g.students, child)]
	b = g.plants[1][slices.Index(g.students, child)]
	for _, letter := range a + b {
		ret = append(ret, PLANTS[string(letter)])
	}
	return ret, true
}
