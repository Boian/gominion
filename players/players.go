package players

import (
    c "github.com/Boian/gominion/cards"
)

type Player struct {
    Position    int // Position of player in start order
    Hand        Hand // the players hand
    Deck        Deck // teh player's deck 
}

func (p *Player) Turn() {
    var actions = 1
    var buys = 1
    // Player turn
    // Optional Play Card
    // Buy card
    p.DiscardHand()
    p.DrawHand()
}

func (p *Player) DrawCard() {
    // Draw a card
    // Take top deck card
        // Reshuffle deck if needed
    // Add card to hand
}

func (p *Player) PlayCard() {
    // Play a card
}

func (p *Player) DiscardCard() {
    // Discard a card
}

func (p *Player) TrashCard() {
    // Trash Card
}

func (p *Player) DiscardHand() {
    // Discard Hand
}

func (p *Player) DrawHand() {
    // Draw Hand
}

func (p *Player) GainCard() {
    // Gain card from supply
}

type Hand struct {
    Cards []c.Card // the cards in a players hand
}

type Deck struct {
    Cards []c.Card // the cards in a players deck
}

type DiscardPile struct {
    Cards []c.Card // the cards in a players discard pile
}

type InPlay struct {
    Cards []c.Card // cards in play
}
