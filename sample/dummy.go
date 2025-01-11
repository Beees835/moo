package sample

import (
	"fmt"
	"os"
	"strings"

	"github.com/speecan/moo/game"
)

// EstimateHuman is played by human
func EstimateHuman(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		var input string
		fmt.Print("?: ")
		fmt.Fscanln(os.Stdin, &input)
		guess := game.Str2Int(strings.Split(input, ""))
		fn(guess)
		return guess
	}
}

// EstimateWithRandom is idiot algo.
// returns estimate number with simply random
func EstimateWithRandom(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}

// EstimateWithRandom2 is idiot algo.
// exclude duplicate queries
func EstimateWithRandom2(difficulty int) game.Estimate {
	query := make([][]int, 0)
	isDuplicated := func(i []int) bool {
		for _, v := range query {
			if game.Equals(v, i) {
				return true
			}
		}
		return false
	}
	return func(fn game.Question) (res []int) {
		var r []int
		for {
			r = game.GetMooNum(difficulty)
			if !isDuplicated(r) {
				break
			}
		}
		fn(r)
		query = append(query, r)
		return r
	}
}

func Est(difficulty int) game.Estimate {
	candidates := game.GetAllCandidates(difficulty)

	return func(q game.Question) []int {
		guess := candidates[0]

		curHits, curBlows := q(guess)
		if curHits == difficulty {
			return guess
		}

		// Else we here proceed to filter out all candidates that doesn't match the current curHits and curBlows
		// As we know if it doesn't match that aren't possible candidates
		newCandidates := candidates[:0]
		for _, c := range candidates {
			if game.GetHit(c, guess) == curHits && game.GetBlow(c, guess) == curBlows {
				newCandidates = append(newCandidates, c)
			}
		}
		candidates = newCandidates

		return guess
	}
}
