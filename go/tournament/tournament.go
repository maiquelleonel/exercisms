package tournament

import (
	"cmp"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

type Match struct {
	home   string
	away   string
	result string
}

type Result struct {
	team    string
	matches int
	wins    int
	draws   int
	losses  int
	points  int
}

type ResultTable map[string]Result

func setResult(match Match, ResultTable ResultTable) (ok bool) {

	home, away := Result{team: match.home}, Result{team: match.away}

	if _, ok := ResultTable[match.home]; !ok {
		ResultTable[match.home] = home
	}

	if _, ok := ResultTable[match.away]; !ok {
		ResultTable[match.away] = away
	}

	home = ResultTable[match.home]
	away = ResultTable[match.away]

	home.matches += 1
	away.matches += 1
	switch match.result {
	case "win":
		home.wins += 1
		away.losses += 1
		home.points += 3
	case "draw":
		home.draws += 1
		away.draws += 1
		home.points += 1
		away.points += 1
	case "loss":
		home.losses += 1
		away.wins += 1
		away.points += 3
	}

	ResultTable[home.team] = home
	ResultTable[away.team] = away

	return true
}

func sort(ResultTable ResultTable) []Result {
	var resultSorted []Result
	for _, team := range ResultTable {
		resultSorted = append(resultSorted, team)
	}

	slices.SortStableFunc(resultSorted, func(a, b Result) int {
		// inverse order                      // 0 is equal
		if n := cmp.Compare(b.points, a.points); n != 0 {
			return n
		}
		return cmp.Compare(a.team, b.team)
	})

	return resultSorted
}

func Tally(reader io.Reader, writer io.Writer) error {
	var ResultTable ResultTable = make(ResultTable, 0)
	var data string
	if b, err := io.ReadAll(reader); err == nil {
		data = string(b)
	}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) < 2 {
			continue
		}

		if line[0] == '#' {
			continue
		}

		match := strings.Split(line, ";")

		if len(match) != 3 {
			return errors.New("invalid match")
		}

		Match := Match{
			home:   strings.TrimSpace(match[0]),
			away:   strings.TrimSpace(match[1]),
			result: strings.TrimSpace(match[2]),
		}

		if slices.Index([]string{"win", "draw", "loss"}, Match.result) == -1 {
			return errors.New("invalid result")
		}

		setResult(Match, ResultTable)

	}

	io.WriteString(writer, "Team                           | MP |  W |  D |  L |  P\n")
	for _, result := range sort(ResultTable) {
		result.team += strings.Repeat(" ", 31-len(result.team))
		line := fmt.Sprintf("%v|  %v |  %v |  %v |  %v |  %v\n",
			result.team, result.matches, result.wins, result.draws, result.losses, result.points)
		io.WriteString(writer, line)
	}

	return nil
}
