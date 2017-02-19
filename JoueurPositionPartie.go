package models

import "github.com/jinzhu/gorm"

// JoueurPositionPartie est une modélisation de la Position
// d'un joueur dans une partie
type JoueurPositionPartie struct {
	gorm.Model
	Joueur     Joueur
	JoueurID   int
	Partie     Partie
	PartieID   int
	Position   Position
	PositionID int
	Equipe     Equipe
	EquipeID   int
}

// Expand effectue un fetch de tous les children de la
// JoueurPositionPartie
func (j *JoueurPositionPartie) Expand(db *gorm.DB) {
	db.Model(j).Related(&(j.Position))
}
