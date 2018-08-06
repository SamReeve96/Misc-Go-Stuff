package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Player struct {
	Name string
	dice []Die
}

type Die struct {
	value int
}

func (die *Die) Roll() int {
	return rand.Intn(6)
}

func (player *Player) LoseDie() {
	playersDice := player.dice
	//Dont care what die goes just remove
	playersDice = append(playersDice[:1], playersDice[2:]...)
}

func main() {
	roundNumber := 1
	players := gameStartup()
	RoundRunner(roundNumber, players)

}

func gameStartup() []*Player {
	var players []*Player
	var numPlayersString string
	println("Let's Play Dudo! How many players?")
	_, err := fmt.Scan(&numPlayersString)
	numPlayers, err := strconv.Atoi(numPlayersString)
	if err != nil {
		println("something went wrong")
	}
	for i := 0; i < numPlayers; i++ {
		newPlayer := &Player{
			Name: AssignName(),
			dice: RollDice(6),
		}
		players = append(players, newPlayer)
	}
	println("Great! Lets Play Dudo with:")
	for _, player := range players {
		println(player.Name)
	}
	return players
}

func AssignName() string {
	var newPlayername string
	println("Enter Player Name")
	_, err := fmt.Scan(&newPlayername)
	if err != nil {
		println("something went wrong")
	}
	return newPlayername
}

func RollDice(numOfDice int) []Die {
	Dice := []Die{}
	for i := 0; i < numOfDice; i++ {
		newDie := Die{
			value: rand.Intn(6),
		}
		Dice = append(Dice, newDie)
	}
	return Dice
}

func RoundRunner(roundNumber int, players []*Player) {
	currentVal := 0
	currentQuant := 0
	var prevPlayer *Player
	println("Round: " + strconv.Itoa(roundNumber))
	for i := 0; i < len(players)+1; i++ {
		if i == len(players) {
			i = -1
		} else {
			println("---- " + players[i].Name + "'s Turn ----")
			if currentVal == 0 {
				currentQuant, currentVal = startingBet()
				prevPlayer = players[i]
			} else {
				println("Current bet is: " + strconv.Itoa(currentQuant) + " " + strconv.Itoa(currentVal) + "'s")
				inputValid := false
				for inputValid == false {
					println("Bet(1) or Call BS(2)")
					var choice string
					_, err := fmt.Scan(&choice)
					if err != nil {
						println("something went wrong")
					}
					if choice == "1" {
						inputValid = true
						currentQuant, currentVal = bet(currentQuant, currentVal)
						prevPlayer = players[i]
					} else if choice == "2" {
						inputValid = true
						println("Calling BS")
						revealDice(players)
						callBS(players, players[i], prevPlayer, currentQuant, currentVal)
						println("New Round")
						roundNumber++
						for i := 0; i < len(players); i++ {
							players[i].dice = RollDice(len(players[i].dice))
						}
						RoundRunner(roundNumber, players)
					} else {
						println("Invalid input")
					}
				}
			}
		}
	}
}

func startingBet() (int, int) {
	var quantBet string
	println("Enter Quantity:")
	_, err := fmt.Scan(&quantBet)
	if err != nil {
		println("something went wrong")
	}
	var valueBet string
	println("Enter Value: (1-6)")
	_, err = fmt.Scan(&valueBet)
	if err != nil {
		println("something went wrong")
	}
	Quantity, _ := strconv.Atoi(quantBet)
	Value, _ := strconv.Atoi(valueBet)
	return Quantity, Value
}

func bet(currentQuant int, currentVal int) (int, int) {
	var quantBet string
	println("Enter Quantity:")
	_, err := fmt.Scan(&quantBet)
	if err != nil {
		println("something went wrong")
	}
	var valueBet string
	println("Enter Value: (1-6)")
	_, err = fmt.Scan(&valueBet)
	if err != nil {
		println("something went wrong")
	}
	Quantity, _ := strconv.Atoi(quantBet)
	Value, _ := strconv.Atoi(valueBet)
	return Quantity, Value
}

func countDiceValues(players []*Player) []int {
	quantityOfValues := []int{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < len(players); i++ {
		currentPlayer := players[i]
		for j := 0; j < len(currentPlayer.dice); j++ {
			currentDie := currentPlayer.dice[i]
			switch currentDie.value {
			case 1:
				quantityOfValues[0]++
			case 2:
				quantityOfValues[1]++
			case 3:
				quantityOfValues[2]++
			case 4:
				quantityOfValues[3]++
			case 5:
				quantityOfValues[4]++
			case 6:
				quantityOfValues[5]++
			}
		}
	}
	return quantityOfValues
}

func revealDice(players []*Player) {
	var quantityOfValues []int
	quantityOfValues = countDiceValues(players)
	fmt.Println("There were:")
	for i := 0; i < len(quantityOfValues); i++ {
		fmt.Println(strconv.Itoa(quantityOfValues[i]) + " " + strconv.Itoa(i+1) + "'s")
	}
}

func callBS(players []*Player, caller *Player, accused *Player, currentQuant int, currentVal int) {
	//Calling BS
	quantityOfValues := countDiceValues(players)
	fmt.Println("BS Called on: " + strconv.Itoa(currentQuant) + " " + strconv.Itoa(currentVal) + "'s")
	if currentQuant <= quantityOfValues[currentVal-1] {
		println(caller.Name + " was Wrong! They lose a Die")
		caller.LoseDie()
	} else {
		println(caller.Name + " was Right! " + accused.Name + " lose a Die")
		accused.LoseDie()
	}
}
