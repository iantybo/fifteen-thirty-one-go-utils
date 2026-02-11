package analysis

import (
	"sort"

	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/cards"
)

// DiscardOption represents a possible discard choice and its expected value.
type DiscardOption struct {
	Keep          []cards.Card `json:"keep"`
	Discard       []cards.Card `json:"discard"`
	AvgHandScore  float64      `json:"avg_hand_score"`
	MinHandScore  int          `json:"min_hand_score"`
	MaxHandScore  int          `json:"max_hand_score"`
}

// ScoreHandFunc is a function that scores a 4-card hand given a cut card.
// Returns the total point value of the hand.
type ScoreHandFunc func(hand []cards.Card, cut cards.Card) int

// BestDiscard evaluates all ways to keep 4 cards from a 6-card dealt hand,
// computing the average hand score across all possible cut cards.
// Returns discard options sorted by average score descending.
func BestDiscard(dealt []cards.Card, scoreFn ScoreHandFunc) []DiscardOption {
	if len(dealt) != 6 || scoreFn == nil {
		return nil
	}

	deck := cards.FullDeck()
	remaining := cards.Remove(deck, dealt)

	discardCombos := cards.Combinations(dealt, 2)
	options := make([]DiscardOption, 0, len(discardCombos))

	for _, discard := range discardCombos {
		keep := cards.Remove(dealt, discard)

		totalScore := 0
		minScore := 999
		maxScore := 0

		for _, cut := range remaining {
			score := scoreFn(keep, cut)
			totalScore += score
			if score < minScore {
				minScore = score
			}
			if score > maxScore {
				maxScore = score
			}
		}

		avgScore := float64(totalScore) / float64(len(remaining))
		options = append(options, DiscardOption{
			Keep:         keep,
			Discard:      discard,
			AvgHandScore: avgScore,
			MinHandScore: minScore,
			MaxHandScore: maxScore,
		})
	}

	sort.Slice(options, func(i, j int) bool {
		return options[i].AvgHandScore > options[j].AvgHandScore
	})

	return options
}

// TopDiscard returns the top N discard options, or all if n > total options.
func TopDiscard(dealt []cards.Card, scoreFn ScoreHandFunc, n int) []DiscardOption {
	all := BestDiscard(dealt, scoreFn)
	if n >= len(all) {
		return all
	}
	return all[:n]
}
