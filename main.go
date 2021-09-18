package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
)

type position uint

const (
	VERTICAL_UP position = iota
	VERTICAL_DOWN
	HORIZONTAL_LEFT
	HORIZONTAL_RIGHT
)

type STATE uint

const (
	DESTROYED STATE = iota
	ALIVE
)

type board struct {
	cells [10][10]*boat
}

type player struct {
	name    string
	address string
	*board
}

type coord struct {
	x uint
	y uint
}

type boatCharacteristics struct {
	name string
	size uint
}

type boat struct {
	boatCharacteristics
	player *player
	hits   map[coord]player
}

// Initialisation direct avec mot "var" car on se trouve à l'extérieur d'une fonction.
// S'applique pour n'importe quelle structure de données

var shipBlueprints = map[string]boatCharacteristics{
	"Destroyer": {
		name: "Destroyer",
		size: 6,
	},
	"Cruiser": {
		name: "Cruiser",
		size: 5,
	},
	"Hunter": {
		name: "Hunter",
		size: 3,
	},
	"Aircraft": {
		name: "AircraftCarrier",
		size: 10,
	},
}

func (g *board) String() string {

	var result bytes.Buffer // Utilisation d'un buffer pour écononmiser de la mémoire
	// L'opération de concaténation avec + génère une nouvelle variable string à chaque fois
	// On utilisera une mémoire tampon pour ne créer qu'une string au final

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if g.cells[x][y] != nil {
				result.WriteString(string(g.cells[x][y].name[0]))
			} else {
				result.WriteString("-")
			}
		}
		result.WriteString("\n")
	}

	return result.String()
}

func determineState(b *boat) STATE {
	switch b.size {
	case 0:
		return DESTROYED
	default:
		return ALIVE
	}
}

func (g *board) addBoat(b *boat, POSITION position, x uint, y uint) error {

	// Are coordinates valid ?
	if x > 10 || y > 10 {
		return fmt.Errorf("position (%v,%v) is not valid", x, y)
	}

	switch POSITION {
	case VERTICAL_DOWN:

		if y+b.size > 10 {
			return fmt.Errorf("%v cannot be placed here", b.name)
		}

		for i := y; i < y+b.size; i++ {
			if g.cells[x][i] != nil {
				return fmt.Errorf("%v cannot be placed here", b.name)
			}
		}

		for i := y; i < y+b.size; i++ {
			g.cells[x][i] = b
		}
		fmt.Printf("Boat '%v' placed at position (%v,%v)\n", b.name, x, y)

	case VERTICAL_UP:

		if y+b.size > 10 {
			return fmt.Errorf("%v cannot be placed here", b.name)
		}

		for i := y; i > y-b.size; i-- {
			if g.cells[x][i] != nil {
				return errors.New("boat cannot be placed here")
			}
		}

		for i := y; i > y-b.size; i-- {
			g.cells[x][i] = b
		}

		fmt.Printf("Boat '%v' placed at position (%v,%v)\n", b.name, x, y)

	case HORIZONTAL_LEFT:
		if int(x-b.size) < 0 {
			return fmt.Errorf("%v cannot be placed here", b.name)
		}

		for i := x; i > x-b.size; i-- {
			if g.cells[i][y] != nil {
				return fmt.Errorf("%v cannot be placed here", b.name)
			}
		}

		for i := x; i > x-b.size; i-- {
			g.cells[i][y] = b
		}

		fmt.Printf("Boat '%v' placed at position (%v,%v)\n", b.name, x, y)

	case HORIZONTAL_RIGHT:

		if x+b.size > 10 {
			return fmt.Errorf("%v cannot be placed here", b.name)
		}

		for i := x; i < x+b.size; i++ {
			if g.cells[i][y] != nil {
				return errors.New("boat cannot be placed here")
			}
		}

		for i := x; i < x+b.size; i++ {
			g.cells[i][y] = b
		}

		fmt.Printf("Boat '%v' placed at position (%v,%v)\n", b.name, x, y)

	}

	return nil
}

func (p *player) attack(opponent *player, x uint, y uint) {
	if opponentBoat := opponent.board.cells[x][y]; opponentBoat != nil {
		fmt.Printf("Well done %v ! %v's %v has been hit !\n", p.name, opponent.name, opponentBoat.name)
		opponentBoat.size--
		opponent.board.cells[x][y] = nil
	} else {
		fmt.Println("Missed !")
	}
}

func (p *player) shipFactory(choice string) (*boat, error) {

	if blueprint, inMap := shipBlueprints[choice]; inMap {

		return &boat{
			boatCharacteristics: blueprint,
			player:              p,
		}, nil
	}

	return nil, fmt.Errorf("desired ship named '%v' could not be found in our blueprints", choice)
}

func main() {

	player1 := &player{
		name:    "ii02735",
		address: "localhost",
		board:   &board{},
	}

	player2 := &player{
		name:    "rahim",
		address: "localhost",
		board:   &board{},
	}
	for _, _player := range []*player{player1, player2} {
		fmt.Printf("%v's board preparations...", _player.name)
		ship, err := _player.shipFactory("Destroyer")
		handleError(err)
		err = _player.board.addBoat(ship, VERTICAL_DOWN, 0, 3)
		handleError(err)
		ship, err = _player.shipFactory("Hunter")
		handleError(err)
		err = _player.board.addBoat(ship, HORIZONTAL_RIGHT, 6, 7)
		handleError(err)
		err = _player.board.addBoat(ship, HORIZONTAL_LEFT, 6, 6)
		handleError(err)
		err = _player.board.addBoat(ship, VERTICAL_UP, 9, 6)
		handleError(err)
	}

	player1.attack(player2, 6, 6)
	player1.attack(player2, 6, 6)

}

func handleError(err error) {

	if err != nil {
		fmt.Printf("Error : %v\n", err.Error())
		os.Exit(1)
	}
}
