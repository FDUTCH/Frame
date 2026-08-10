package Frame

import (
	"github.com/FDUTCH/Frame/event"
	"github.com/df-mc/dragonfly/server/player"
)

// playerHandler stores frame instance inside player's handler.
type playerHandler struct {
	*event.PlayerHandler
	f *Frame
}

// FromPlayer returns Frame instance from player.
func FromPlayer(pl *player.Player) *Frame {
	h, ok := pl.Handler().(playerHandler)
	if !ok {
		return nil
	}
	return h.f
}
