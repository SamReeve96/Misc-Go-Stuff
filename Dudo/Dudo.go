package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Player struct {
	Name      string
	diceAlive []Die
}

type Die struct {
	value int
}

func (die *Die) roll() int {
	return rand.Intn(6)
}

func main() {
	roundNumber := 1
	players := gameStartup()
	roundRunner(roundNumber, players)

}

func gameStartup() []*Player {
	var numPlayersString string
	var players []*Player

	println("Let's Play Dudo! How many players?")
	fmt.Scanln(&numPlayersString)
	numPlayers, err := strconv.Atoi(numPlayersString)
	if err != nil {
		println("something went wrong")
	}
	for i := 0; i < numPlayers; i++ {

		newPlayer := &Player{
			Name:      assignName(),
			diceAlive: rollDice(),
		}
		players = append(players, newPlayer)
	}
	println("Great! Lets Play Dudo with:")
	for _, player := range players {
		println(player.Name)
	}
	return players
}

func assignName() string {
	reader := bufio.NewReader(os.Stdin)
	var newPlayername string
	println("Enter Player Name")
	newPlayername, _ = reader.ReadString('\n')
	return newPlayername
}

func rollDice() []Die {
	Dice := []Die{}
	for i := 0; i < 5; i++ {
		newDie := Die{
			value: rand.Intn(6),
		}
		Dice = append(Dice, newDie)
	}
	return Dice
}

func roundRunner(roundNumber int, players []*Player) {
	roundNumber++
	println("Round: " + strconv.Itoa(roundNumber))
	for _, player := range players {
		println(player.Name + "'s Turn")
	}
}
