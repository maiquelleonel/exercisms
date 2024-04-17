package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
	seen map[string]bool
}

var max int = 26 * 26 * 10 * 10 * 10
var Seen = make(map[string]bool)

func (r *Robot) Name() (string, error) {
	var err error
	if r.seen == nil {
		r.seen = Seen
		if err = r.genName(); err != nil {
			return "", err
		}
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	if r.seen[r.name] {
		r.seen[r.name] = false
		r.genName()
	}
}

func (r *Robot) genName() error {
	for len(r.seen)+1 <= max {
		candidate := r.makeName()
		if _, seen := r.seen[candidate]; !seen {
			r.seen[candidate] = true
			r.name = candidate
			return nil
		}
	}
	return errors.New("no new name is possible")
}

func (r Robot) makeName() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	first := rand.Intn(len(letters))
	second := rand.Intn(len(letters))
	return fmt.Sprintf("%s%s%03d", string(letters[first]), string(letters[second]), rand.Intn(1000))
}
