package main

import (
    "fmt"
    s "github.com/Boian/gominion/setup"
)

func Play() {
    fmt.Println("Time to Play the Game")
    game := s.Setup()
    fmt.Println(len(game.Trash))
}

func main() {
    Play()
}
