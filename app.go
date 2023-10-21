package main

import (
	"context"
	"encoding/json"

	"github.com/j-tt/character-todo/server"
)

// App struct
type App struct {
	ctx context.Context
	*server.WSServer
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.WSServer = server.InitWeb()
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) PushCreature(creatures []server.Creature) {
	//fmt.Println(creatureStr)
	// TODO: intermediate state lib

	a.Publish(creatures)
}

func (a *App) LoadState() string {
	// TODO: intermediate state lib

	//a.Publish(name)
	creatures, _ := a.GetState()
	bytes, _ := json.Marshal(creatures)
	return string(bytes)
}
