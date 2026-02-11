package cards

import (
	"fmt"
	"strings"
)

// Suit represents a card suit.
type Suit string

const (
	Spades   Suit = "S"
	Hearts   Suit = "H"
	Diamonds Suit = "D"
	Clubs    Suit = "C"
)

// AllSuits returns all four suits.
func AllSuits() []Suit {
	return []Suit{Spades, Hearts, Diamonds, Clubs}
}

// Rank represents a card rank (1=Ace through 13=King).
type Rank int

const (
	Ace   Rank = 1
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
)

// AllRanks returns all thirteen ranks in order.
func AllRanks() []Rank {
	return []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
}

// Card represents a playing card with a rank and suit.
type Card struct {
	Rank Rank `json:"rank"`
	Suit Suit `json:"suit"`
}

// String returns a human-readable representation like "AS", "10H", "KD".
func (c Card) String() string {
	var r string
	switch c.Rank {
	case Ace:
		r = "A"
	case Jack:
		r = "J"
	case Queen:
		r = "Q"
	case King:
		r = "K"
	default:
		r = fmt.Sprintf("%d", int(c.Rank))
	}
	return r + string(c.Suit)
}

// Value15 returns the card's value for computing fifteens.
// Face cards are 10, ace is 1.
func (c Card) Value15() int {
	if c.Rank >= 10 {
		return 10
	}
	return int(c.Rank)
}

// Parse parses a card string like "AS", "10H", "KD" into a Card.
func Parse(s string) (Card, error) {
	s = strings.TrimSpace(strings.ToUpper(s))
	if len(s) < 2 {
		return Card{}, fmt.Errorf("invalid card string: %q", s)
	}
	suit := Suit(s[len(s)-1:])
	rankStr := s[:len(s)-1]
	var r Rank
	switch rankStr {
	case "A":
		r = Ace
	case "J":
		r = Jack
	case "Q":
		r = Queen
	case "K":
		r = King
	default:
		var v int
		_, err := fmt.Sscanf(rankStr, "%d", &v)
		if err != nil || v < 2 || v > 10 {
			return Card{}, fmt.Errorf("invalid rank: %q", rankStr)
		}
		r = Rank(v)
	}
	switch suit {
	case Spades, Hearts, Diamonds, Clubs:
	default:
		return Card{}, fmt.Errorf("invalid suit: %q", string(suit))
	}
	return Card{Rank: r, Suit: suit}, nil
}

// FullDeck returns all 52 cards in a standard deck.
func FullDeck() []Card {
	deck := make([]Card, 0, 52)
	for _, s := range AllSuits() {
		for _, r := range AllRanks() {
			deck = append(deck, Card{Rank: r, Suit: s})
		}
	}
	return deck
}
