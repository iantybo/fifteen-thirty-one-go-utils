package analysis

import (
	"sort"

	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/cards"
)

// ScoreBreakdown provides a detailed breakdown of a cribbage hand score.
type ScoreBreakdown struct {
	Total    int `json:"total"`
	Fifteens int `json:"fifteens"`
	Pairs    int `json:"pairs"`
	Runs     int `json:"runs"`
	Flush    int `json:"flush"`
	Nobs     int `json:"nobs"`
}

// ScoreHand scores a 4-card cribbage hand with a cut card.
func ScoreHand(hand []cards.Card, cut cards.Card) ScoreBreakdown {
	all := make([]cards.Card, 0, 5)
	all = append(all, hand...)
	all = append(all, cut)

	sb := ScoreBreakdown{}
	sb.Fifteens = scoreFifteens(all)
	sb.Pairs = scorePairs(all)
	sb.Runs = scoreRuns(all)
	sb.Flush = scoreFlush(hand, cut)
	sb.Nobs = scoreNobs(hand, cut)
	sb.Total = sb.Fifteens + sb.Pairs + sb.Runs + sb.Flush + sb.Nobs
	return sb
}

// ScoreHandTotal returns just the total score, compatible with ScoreHandFunc.
func ScoreHandTotal(hand []cards.Card, cut cards.Card) int {
	return ScoreHand(hand, cut).Total
}

func scoreFifteens(all []cards.Card) int {
	n := len(all)
	points := 0
	for mask := 1; mask < (1 << n); mask++ {
		sum := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += all[i].Value15()
			}
		}
		if sum == 15 {
			points += 2
		}
	}
	return points
}

func scorePairs(all []cards.Card) int {
	count := map[cards.Rank]int{}
	for _, c := range all {
		count[c.Rank]++
	}
	points := 0
	for _, n := range count {
		if n >= 2 {
			points += (n * (n - 1) / 2) * 2
		}
	}
	return points
}

func scoreRuns(all []cards.Card) int {
	count := map[int]int{}
	var ranks []int
	for _, c := range all {
		r := int(c.Rank)
		if count[r] == 0 {
			ranks = append(ranks, r)
		}
		count[r]++
	}
	sort.Ints(ranks)

	bestLen := 0
	bestMult := 0
	for start := 0; start < len(ranks); start++ {
		for end := start; end < len(ranks); end++ {
			runLen := end - start + 1
			if runLen < 3 {
				continue
			}
			if ranks[end]-ranks[start] != runLen-1 {
				continue
			}
			mult := 1
			for i := start; i <= end; i++ {
				mult *= count[ranks[i]]
			}
			if runLen > bestLen {
				bestLen = runLen
				bestMult = mult
			} else if runLen == bestLen {
				bestMult += mult
			}
		}
	}
	if bestLen == 0 {
		return 0
	}
	return bestLen * bestMult
}

func scoreFlush(hand []cards.Card, cut cards.Card) int {
	if len(hand) != 4 {
		return 0
	}
	s := hand[0].Suit
	for i := 1; i < 4; i++ {
		if hand[i].Suit != s {
			return 0
		}
	}
	if cut.Suit == s {
		return 5
	}
	return 4
}

func scoreNobs(hand []cards.Card, cut cards.Card) int {
	for _, c := range hand {
		if c.Rank == cards.Jack && c.Suit == cut.Suit {
			return 1
		}
	}
	return 0
}
