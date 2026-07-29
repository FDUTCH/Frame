package targets

import (
	"iter"

	"github.com/FDUTCH/Frame/translations"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

func OriginPlayer(o *cmd.Output, src cmd.Source) *player.Player {
	pl, ok := src.(*player.Player)
	if !ok {
		o.Errort(translations.InvalidOrigin)
	}
	return pl
}

func CallForAllOptional[T any](o *cmd.Output, targets cmd.Optional[[]cmd.Target]) (iter.Seq[T], bool) {
	t, ok := targets.Load()
	if !ok {
		return nil, false
	}
	return CallForAll[T](o, t)
}

func CallForAll[T any](o *cmd.Output, targets []cmd.Target) (iter.Seq[T], bool) {
	return func(yield func(T) bool) {
		var called bool
		for _, target := range targets {
			val, ok := target.(T)
			if !ok {
				continue
			}
			called = true
			if !yield(val) {
				return
			}
		}
		if !called {
			o.Errort(translations.NoTargetMatch)
		}
	}, len(targets) > 0
}

func CallByName(srv *server.Server, o *cmd.Output, str string, fn func(tx *world.Tx, player *player.Player)) bool {
	h, ok := srv.PlayerByName(str)
	if !ok {
		o.Errort(translations.PlayerNotFound)
		return false
	}
	err := player.NewRef(h).Do(fn).Err()
	if err != nil {
		o.Error(err)
	}
	return err == nil
}

func SinglePlayerOptional(o *cmd.Output, targets cmd.Optional[[]cmd.Target]) *player.Player {
	val, ok := targets.Load()
	if !ok {
		return nil
	}
	return SinglePlayer(o, val)
}

func SinglePlayer(o *cmd.Output, targets []cmd.Target) *player.Player {
	if len(targets) > 1 {
		o.Errort(translations.ToManyTargets)
		return nil
	} else if len(targets) == 0 {
		o.Errort(translations.NoTargetMatch)
		return nil
	}
	pl, ok := targets[0].(*player.Player)
	if !ok {
		o.Errort(translations.TargetNotPlayer)
		return nil
	}
	return pl
}
