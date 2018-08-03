package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)
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
		newPlayername, _ = reader.ReadString('\n')
		newPlayer := &Player{
			Name: newPlayername,
		}
		players = append(players, newPlayer)
	}
	println("Great! Lets Play Dudo with:")
	for _, player := range players {
		println(player.Name)
	}
	round := 0
	println("Round: " + strconv.Itoa(round))
}
