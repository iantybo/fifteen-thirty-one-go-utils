package contract

import (
	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/analysis"
	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/cards"
)

type analysisImpl struct{}

func (analysisImpl) ScoreTotal(hand []cards.Card, cut cards.Card) int {
	return analysis.ScoreHandTotal(hand, cut)
}

// DefaultHandAnalyzer returns the canonical HandAnalyzer backed by pkg/analysis.
func DefaultHandAnalyzer() HandAnalyzer {
	return analysisImpl{}
}

var _ HandAnalyzer = analysisImpl{}
