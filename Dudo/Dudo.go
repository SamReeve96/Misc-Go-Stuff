package main

import (
	"fmt"
	"strconv"
)

type Player struct {
	Name      string
	diceAlive []Dice
}

type Dice struct {
	value int
}

func main() {
	var numPlayersString string
	var players []*Player

	println("Let's Play Dudo! How many players?")
	fmt.Scanln(&numPlayersString)
	numPlayers, err := strconv.Atoi(numPlayersString)
	if err != nil {
		println("something went wrong")
	}
	for i := 0; i < numPlayers; i++ {
		var newPlayername string
		println("Enter Player Name")
		fmt.Scanln(&newPlayername)
		newPlayer := &Player{
			Name: newPlayername,
		}
		players = append(players, newPlayer)
	}
	println("Great! Lets Play Dudo with:")
	for i := 0; i < numPlayers; i++ {
		println(numPlayers[i])
	}
}

func showName(Player player) {
	player.showName()
}
