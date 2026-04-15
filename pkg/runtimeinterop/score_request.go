package runtimeinterop

import (
	"fmt"

	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/analysis"
	"github.com/iantybo/fifteen-thirty-one-go-utils/pkg/cards"
)

// ScoreRequest carries runtime-provided card codes for counting verification.
type ScoreRequest struct {
	Hand   []string `json:"hand"`
	Cut    string   `json:"cut"`
	IsCrib bool     `json:"is_crib"`
}

// ScoreResponse contains a hand score and component breakdown.
type ScoreResponse struct {
	Total    int `json:"total"`
	Fifteens int `json:"fifteens"`
	Pairs    int `json:"pairs"`
	Runs     int `json:"runs"`
	Flush    int `json:"flush"`
	Nobs     int `json:"nobs"`
}

// ScoreFromCodes parses request card codes and computes cribbage score.
func ScoreFromCodes(req ScoreRequest) (ScoreResponse, error) {
	if len(req.Hand) != 4 {
		return ScoreResponse{}, fmt.Errorf("hand must have exactly 4 cards, got %d", len(req.Hand))
	}

	hand := make([]cards.Card, len(req.Hand))
	for i, code := range req.Hand {
		c, err := cards.Parse(code)
		if err != nil {
			return ScoreResponse{}, fmt.Errorf("invalid hand card %q: %w", code, err)
		}
		hand[i] = c
	}

	cut, err := cards.Parse(req.Cut)
	if err != nil {
		return ScoreResponse{}, fmt.Errorf("invalid cut card %q: %w", req.Cut, err)
	}

	return ScoreFromCards(hand, cut, req.IsCrib), nil
}

// ScoreFromCards scores parsed cards and applies crib-specific flush rules.
func ScoreFromCards(hand []cards.Card, cut cards.Card, isCrib bool) ScoreResponse {
	base := analysis.ScoreHand(hand, cut)
	flush := base.Flush
	total := base.Total

	// utils analysis uses hand flush rules; adjust for crib here.
	if isCrib && flush == 4 {
		total -= 4
		flush = 0
	}

	return ScoreResponse{
		Total:    total,
		Fifteens: base.Fifteens,
		Pairs:    base.Pairs,
		Runs:     base.Runs,
		Flush:    flush,
		Nobs:     base.Nobs,
	}
}
