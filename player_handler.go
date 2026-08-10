package Frame

import (
	"github.com/FDUTCH/Frame/event"
	"github.com/FDUTCH/Frame/storage"
	"github.com/df-mc/dragonfly/server/player"
)

// playerHandler stores frame instance inside player's handler.
type playerHandler struct {
	*event.PlayerHandler
	f       *Frame
	storage *storage.Storage
}

func newPlayerHandler(f *Frame) *playerHandler {
	return &playerHandler{PlayerHandler: event.NewPlayerHandler(f, f.generalBus), f: f, storage: storage.NewStorage()}
}

// FromPlayer returns Frame instance from player.
func FromPlayer(pl *player.Player) *Frame {
	h, ok := pl.Handler().(*playerHandler)
	if !ok {
		return nil
	}
	return h.f
}

// GetPlayerValue ...
func GetPlayerValue[T any](pl *player.Player) (T, bool) {
	var val T
	h, ok := pl.Handler().(*playerHandler)
	if !ok {
		return val, false
	}
	return storage.Get[T](h.storage)
}

// SetPlayerValue ...
func SetPlayerValue[T any](pl *player.Player, val T) {
	h, ok := pl.Handler().(*playerHandler)
	if ok {
		storage.Set(h.storage, val)
	}
}
