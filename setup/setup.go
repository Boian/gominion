package setup

import (
    c "github.com/Boian/gominion/cards"
    p "github.com/Boian/gominion/players"
)

func Setup() Game {
    // Function to setup the game
    game := new(Game)
    return *game
}

type Trash struct {
    Cards []c.Card
}

type Pile struct {
    Cards []c.Card
}

type Supply struct {
    Piles []Pile  // Map Pile
}

type Game struct {
    CardSupply Supply
    Trash []c.Card
    Players []p.Player
}
