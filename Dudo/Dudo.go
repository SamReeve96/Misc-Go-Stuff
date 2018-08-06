package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	currentVal := 0
	currentQuant := 0
	reader := bufio.NewReader(os.Stdin)
	println("Round: " + strconv.Itoa(roundNumber))
	for i := 0; i < len(players)+1; i++ {
		if i == len(players) {
			i = 0
		} else {
			println(players[i].Name + "'s Turn")
			if currentVal == 0 {
				currentVal, currentQuant = startingBet()
			} else {
				println("Current bet is: " + strconv.Itoa(currentQuant) + " " + strconv.Itoa(currentVal) + "'s")
				inputValid := false
				for inputValid == false {
					println("Bet(1) or Call BS(2)")
					stringChoice, _ := reader.ReadString('\n')
					if stringChoice == "1\n" {
						currentVal, currentQuant = bet(currentVal, currentQuant)
					} else if stringChoice == "2\n" {
						callBS()
						currentVal = 0
						currentQuant = 0
						i = 0
						roundNumber++
					} else {
						println("Invalid input")
					}
				}
			}
		}
	}
} ///////////////////////////////////////////////Notes Broken Menu and reversed Quant and val

func startingBet() (int, int) {
	reader := bufio.NewReader(os.Stdin)
	println("Enter Starting bet in the format [Quantity][Space][value]")
	betMadeString, _ := reader.ReadString('\n')
	betMade := strings.Fields(betMadeString)
	Quantity, _ := strconv.Atoi(betMade[0])
	Value, _ := strconv.Atoi(betMade[1])
	return Quantity, Value
}

func bet(currentVal int, currentQuant int) (int, int) {
	reader := bufio.NewReader(os.Stdin)
	println("Enter Starting bet in the format [Quantity][Space][value]")
	betMadeString, _ := reader.ReadString('\n')
	betMade := strings.Fields(betMadeString)
	Quantity, _ := strconv.Atoi(betMade[0])
	Value, _ := strconv.Atoi(betMade[1])
	return Quantity, Value
}

func callBS() {
	//Calling BS
}
