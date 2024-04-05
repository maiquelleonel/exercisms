package tournament

import (
	"bufio"
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
	team                                 string
	matches, wins, draws, losses, points int
}

type ResultsTable map[string]*Result

func (rt ResultsTable) setResult(match Match) (ok bool) {

	home, away := &Result{team: match.home}, &Result{team: match.away}

	if rt[match.home] == nil {
		rt[match.home] = home
	}

	if rt[match.away] == nil {
		rt[match.away] = away
	}

	home = rt[match.home]
	away = rt[match.away]

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

	rt[home.team] = home
	rt[away.team] = away

	return true
}

func (r *ResultsTable) sort() []Result {
	var resultSorted []Result
	for _, team := range *r {
		resultSorted = append(resultSorted, *team)
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
	ResultsTable := &ResultsTable{}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		match := strings.Split(line, ";")

		if len(match) != 3 {
			return errors.New("invalid match")
		}

		Match := &Match{
			home:   strings.TrimSpace(match[0]),
			away:   strings.TrimSpace(match[1]),
			result: strings.TrimSpace(match[2]),
		}

		if slices.Index([]string{"win", "draw", "loss"}, Match.result) == -1 {
			return errors.New("invalid result")
		}

		ResultsTable.setResult(*Match)

	}

	io.WriteString(writer, fmt.Sprintf("%-30s | MP |  W |  D |  L |  P\n", "Team"))
	for _, result := range ResultsTable.sort() {
		io.WriteString(writer,
			fmt.Sprintf("%-30s |  %v |  %v |  %v |  %v |  %v\n",
				result.team, result.matches, result.wins, result.draws, result.losses, result.points))
	}

	return nil
}
