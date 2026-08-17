package Frame

import (
	"context"
	"iter"

	"github.com/df-mc/dragonfly/server/world"
)

// Entities returns iterator that yields entities from handles passed, batching world transactions for entities in the same world.
func Entities(tx *world.Tx, entities ...*world.EntityHandle) iter.Seq2[world.Entity, bool] {
	return func(yield func(world.Entity, bool) bool) {
		fetcher := &entFetcher{}
		fetcher.fetch(tx, entities, yield)
	}
}

type entFetcher struct {
	held map[*world.Tx]struct{}
}

func (f *entFetcher) fetch(original *world.Tx, handles []*world.EntityHandle, yield func(entity world.Entity, ok bool) bool) bool {
	if len(handles) == 0 {
		return false
	}

	h := handles[0]
	rest := handles[1:]

	if original != nil {
		if ent, ok := h.Entity(original); ok {
			return yield(ent, true) && f.fetch(original, rest, yield)
		}
	}

	for tx := range f.held {
		if ent, ok := h.Entity(tx); ok {
			return yield(ent, false) && f.fetch(original, rest, yield)
		}
	}

	continueExecution, _ := world.CallEntity(context.Background(), h, func(tx *world.Tx, e world.Entity) (bool, error) {
		if f.held == nil {
			f.held = make(map[*world.Tx]struct{})
		}
		f.held[tx] = struct{}{}

		return yield(e, false) && f.fetch(original, rest, yield), nil
	})

	return continueExecution
}
