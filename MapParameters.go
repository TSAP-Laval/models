package models

import "github.com/jinzhu/gorm"

// MapParameters est la modélisation de la taille des zones
// de la HeatMap
type MapParameters struct {
	gorm.Model
	Longeur  int
	Largeur  int
	Equipe   Equipe
	EquipeID int
}
