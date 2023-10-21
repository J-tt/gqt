package server

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Creature struct {
	gorm.Model
	Name  string
	Image []byte
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Creature{})
	return db, err
}

func (wss *WSServer) GetState() ([]Creature, error) {
	var creatures []Creature
	err := wss.db.Find(&creatures).Error
	return creatures, err
}
