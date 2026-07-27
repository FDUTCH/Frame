package main

import (
	"log/slog"

	Frame "github.com/FDUTCH/Frame"
	"github.com/FDUTCH/Frame/event"
	"github.com/df-mc/dragonfly/server"
)

func main() {
	cfg, _ := server.DefaultConfig().Config(slog.Default())
	f := Frame.NewFrame(cfg)
	f.Server().CloseOnProgramEnd()
	defer f.Close()
	event.Subscribe(f.GeneralBus(), event.CancelExplosions)
	event.Subscribe(f.GeneralBus(), func(ev event.PlayerJoinEvent) {
		ev.Player.Message("Hi!")
	})

	f.Run()
}
