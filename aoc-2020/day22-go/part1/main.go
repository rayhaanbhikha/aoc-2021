package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Player struct {
	cards []int
}

func NewPlayer(playerDetails string) *Player {
	cards := strings.Split(strings.TrimSpace(playerDetails), "\n")[1:]
	cardsN := make([]int, 0)
	for _, card := range cards {
		cardN, _ := strconv.Atoi(card)
		cardsN = append(cardsN, cardN)
	}
	return &Player{cards: cardsN}
}

func (p *Player) hasCards() bool {
	return len(p.cards) > 0
}

func (p *Player) playCard() (int, bool) {
	if len(p.cards) == 0 {
		return 0, false
	}
	cardToPlay := p.cards[0]
	p.cards = p.cards[1:]
	return cardToPlay, true
}

func (p *Player) addCards(card1, card2 int) {
	p.cards = append(p.cards, card1, card2)
}

func (p *Player) computeScore() (int, bool) {
	n := len(p.cards)
	if n == 0 {
		return 0, false
	}

	sum := 0
	for i, card := range p.cards {
		sum += ((n - i) * card)
	}
	return sum, true
}

func main() {
	data, _ := ioutil.ReadFile("../input")
	input := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	player1 := NewPlayer(input[0])
	player2 := NewPlayer(input[1])

	for player1.hasCards() && player2.hasCards() {
		player1Card, _ := player1.playCard()
		player2Card, _ := player2.playCard()

		if player1Card > player2Card {
			player1.addCards(player1Card, player2Card)
		} else {
			player2.addCards(player2Card, player1Card)
		}
	}

	fmt.Println(player1.computeScore())
	fmt.Println(player2.computeScore())
}
