package stats

import (
	"math"

	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/analysis"
	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/cards"
)

// HandDistribution describes the score distribution of a hand across all possible cuts.
type HandDistribution struct {
	Hand       []cards.Card   `json:"hand"`
	Scores     map[int]int    `json:"scores"`
	AvgScore   float64        `json:"avg_score"`
	MedianScore float64       `json:"median_score"`
	StdDev     float64        `json:"std_dev"`
	MinScore   int            `json:"min_score"`
	MaxScore   int            `json:"max_score"`
	TotalCuts  int            `json:"total_cuts"`
}

// AnalyzeHand computes the full score distribution for a 4-card hand
// across all possible cut cards from the remaining deck.
func AnalyzeHand(hand []cards.Card) HandDistribution {
	deck := cards.FullDeck()
	remaining := cards.Remove(deck, hand)

	scores := make(map[int]int)
	var allScores []int
	totalScore := 0
	minScore := math.MaxInt32
	maxScore := math.MinInt32

	for _, cut := range remaining {
		sb := analysis.ScoreHand(hand, cut)
		s := sb.Total
		scores[s]++
		allScores = append(allScores, s)
		totalScore += s
		if s < minScore {
			minScore = s
		}
		if s > maxScore {
			maxScore = s
		}
	}

	n := len(remaining)
	avg := float64(totalScore) / float64(n)

	// Standard deviation
	variance := 0.0
	for _, s := range allScores {
		diff := float64(s) - avg
		variance += diff * diff
	}
	variance /= float64(n)
	stdDev := math.Sqrt(variance)

	// Median
	sorted := make([]int, len(allScores))
	copy(sorted, allScores)
	sortInts(sorted)
	var median float64
	if n%2 == 0 {
		median = float64(sorted[n/2-1]+sorted[n/2]) / 2.0
	} else {
		median = float64(sorted[n/2])
	}

	return HandDistribution{
		Hand:        hand,
		Scores:      scores,
		AvgScore:    avg,
		MedianScore: median,
		StdDev:      stdDev,
		MinScore:    minScore,
		MaxScore:    maxScore,
		TotalCuts:   n,
	}
}

func sortInts(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

// CompareHands compares two 4-card hands by their average expected score.
// Returns a positive number if hand1 is better, negative if hand2 is better.
func CompareHands(hand1, hand2 []cards.Card) float64 {
	d1 := AnalyzeHand(hand1)
	d2 := AnalyzeHand(hand2)
	return d1.AvgScore - d2.AvgScore
}

// RankHands sorts multiple hands by average expected score (best first).
func RankHands(hands [][]cards.Card) []HandDistribution {
	dists := make([]HandDistribution, len(hands))
	for i, h := range hands {
		dists[i] = AnalyzeHand(h)
	}
	for i := 1; i < len(dists); i++ {
		key := dists[i]
		j := i - 1
		for j >= 0 && dists[j].AvgScore < key.AvgScore {
			dists[j+1] = dists[j]
			j--
		}
		dists[j+1] = key
	}
	return dists
}
