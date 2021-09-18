package main

import (
	"testing"
)

/**
 * Test d'initialisation d'un plateau
 * On s'attend à ce que toutes les cases soient
 * vides
 */
func TestInitBoard(t *testing.T) {
	b := &board{}
	for y, cellArray := range b.cells {

		for x, cell := range cellArray {
			if cell != nil {
				t.Errorf("La case (%v,%v) devrait être vide. Obtenu : %v", y, x, cell)
			}
		}
	}
}

func TestAssignBoardToPlayer(t *testing.T) {

	b := &board{}
	p := &player{name: "test"}
	p2 := &player{name: "test2"}

	p.board = b

	if p.board == nil {
		t.Errorf("Le board ne devrait pas être nul")
	}

	if p2.board != nil {
		t.Errorf("Le joueur %v ne devrait pas avoir de board. Obtenu : %v", p2.name, p2.board)
	}
}

func TestPlacerNavire(t *testing.T) {

	p := &player{name: "test", board: &board{}}
	boat, err := p.shipFactory("Destroyer")

	if err != nil {
		t.Errorf("Erreur : %v", err)
	}

	if boat == nil {
		t.Errorf("Un %v devrait être créé. Obtenu : nil", boat.name)
	}

	err = p.addBoat(boat, VERTICAL_DOWN, 0, 0) //devrait se trouver tout en haut à gauche

	if err != nil {
		t.Errorf("Le %v devrait pouvoir être placé à (0,0). Erreur obtenue : %v", boat.name, err)
	}

	for _, _coord := range []coord{{0, 0}, {0, 1}, {0, 2}} {
		if p.board.cells[_coord.x][_coord.y] == nil {
			t.Errorf("Un %v destroyer devrait être ici. Obtenu : nil", boat.name)
		}
	}

	boat, err = p.shipFactory("Hunter")

	if err != nil {
		t.Errorf("Erreur : %v", err)
	}

	err = p.addBoat(boat, HORIZONTAL_LEFT, 2, 1)

	if err == nil /*cette fois-ci le placement devrait être impossible, on s'attend à une erreur*/ {
		t.Errorf("Un %v ne devrait pas pouvoir être placé à la position (2,1). Il est en conflit avec un %v placé à la position (0,0)", boat.name, p.cells[0][0])
	}
}

// TODO : écrire tests pour attaque
