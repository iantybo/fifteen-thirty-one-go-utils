package contract

import "github.com/iantybo/fifteen-thirty-one-go-utils/pkg/cards"

// HandAnalyzer is the shared contract for cribbage hand scoring used by analysis
// and checked against the fifteen-thirty-one-go server via utilscompat tests.
type HandAnalyzer interface {
	ScoreTotal(hand []cards.Card, cut cards.Card) int
}
