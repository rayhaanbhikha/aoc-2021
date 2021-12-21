package main

import "fmt"

type Die struct {
	numTimesRolled int
	currentVal     int
}

func (d *Die) roll() int {
	d.numTimesRolled++

	d.currentVal++
	if d.currentVal > 100 {
		d.currentVal %= 100
	}
	return d.currentVal
}

func NewDie() *Die {
	return &Die{
		currentVal: 0,
	}
}

type Player struct {
	label        string
	position     int
	currentScore int
}

func (p *Player) rollDie(die *Die) bool {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += die.roll()
	}

	p.position += sum
	if p.position > 10 {
		p.position %= 10
		if p.position == 0 {
			p.position = 10
		}
	}

	p.currentScore += p.position

	return p.currentScore >= 1000
}

func main() {
	die := NewDie()

	p1 := &Player{label: "player 1", position: 7}
	p2 := &Player{label: "player 2", position: 2}

	for {
		if p1Won := p1.rollDie(die); p1Won {
			fmt.Println(die.numTimesRolled * p2.currentScore)
			break
		}

		if p2Won := p2.rollDie(die); p2Won {
			fmt.Println(die.numTimesRolled * p1.currentScore)
			break
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)

}
