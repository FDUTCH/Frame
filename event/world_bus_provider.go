package event

import "github.com/df-mc/dragonfly/server/world"

type WorldBusProvider interface {
	WorldBus(w *world.World) (*Bus, bool)
}
