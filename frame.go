package Frame

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/FDUTCH/Frame/event"
	"github.com/FDUTCH/Frame/storage"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/world"
)

// Frame manages your server & systems lifecycle.
type Frame struct {
	logger *slog.Logger

	srv            *server.Server
	generalBus     *event.Bus
	generalStorage *storage.Storage

	closeBeforeServer []io.Closer
	closeAfterServer  []io.Closer

	once sync.Once

	worlds sync.Map
}

// GeneralStorage returns general storage.
func (f *Frame) GeneralStorage() *storage.Storage {
	return f.generalStorage
}

// NewFrame creates new Frame instance.
func NewFrame(config server.Config) *Frame {
	srv := config.New()
	f := &Frame{srv: srv, logger: config.Log.With("src", "FRAME"), generalBus: event.NewBus(), generalStorage: storage.NewStorage()}
	f.AddWorld(srv.World())
	f.AddWorld(srv.Nether())
	f.AddWorld(srv.End())
	return f
}

// GeneralBus returns general *event.Bus.
func (f *Frame) GeneralBus() *event.Bus {
	return f.generalBus
}

// AddWorld adds world to the internal map.
// World added to the internal map will be automatically closed when ever current Frame will be closed.
// Added world receives it's onw *event.Bus, and new world.Handler is set.
func (f *Frame) AddWorld(w *world.World) *event.Bus {
	ownBus := event.NewBus()
	w.Handle(event.NewWorldHandler(ownBus, f.generalBus))
	f.worlds.Store(w, ownBus)
	return ownBus
}

// CreateNewWorld creates new world and calls AddWorld.
func (f *Frame) CreateNewWorld(config world.Config) (*world.World, *event.Bus) {
	w := config.New()
	return w, f.AddWorld(w)
}

// RemoveWorld removes world.
func (f *Frame) RemoveWorld(w *world.World) {
	f.worlds.Delete(w)
}

// WorldBus returns world's *event.Bus.
func (f *Frame) WorldBus(w *world.World) (*event.Bus, bool) {
	bus, ok := f.worlds.Load(w)
	if ok {
		return bus.(*event.Bus), true
	}
	return nil, false
}

// CloseBeforeServer will close closer before the server.
// Should be called for systems that work with players or worlds.
func (f *Frame) CloseBeforeServer(closer io.Closer) {
	f.closeBeforeServer = append(f.closeBeforeServer, closer)
}

// CloseAfterServer will close closer after the server.
// Should be called for systems that work with sensitive data, like databases etc...
func (f *Frame) CloseAfterServer(closer io.Closer) {
	f.closeAfterServer = append(f.closeAfterServer, closer)
}

// Run setups and runs the server.
func (f *Frame) Run() {
	f.srv.Listen()
	for pl := range f.srv.Accept() {
		ev := event.PlayerJoinEvent{Player: pl}
		if bus, ok := f.WorldBus(pl.Tx().World()); ok {
			event.Publish(bus, ev)
		}
		event.Publish(f.generalBus, ev)
		pl.Handle(newPlayerHandler(f))
	}
}

// CloseOnProgramEnd closes current Frame instance.
func (f *Frame) CloseOnProgramEnd() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		_ = f.Close()
	}()
}

// Server returns server.
func (f *Frame) Server() *server.Server {
	return f.srv
}

// Close closes all closers, server and worlds.
func (f *Frame) Close() error {
	f.once.Do(func() {
		for _, closer := range f.closeBeforeServer {
			if err := closer.Close(); err != nil {
				f.logger.Error(fmt.Sprintf("error closing %T", closer), "err", err)
			}
		}

		if err := f.srv.Close(); err != nil {
			f.logger.Error("error closing server", "err", err)
		}

		for _, closer := range f.closeAfterServer {
			if err := closer.Close(); err != nil {
				f.logger.Error(fmt.Sprintf("error closing %T", closer), "err", err)
			}
		}

		f.worlds.Range(func(key, _ any) bool {
			_ = key.(*world.World).Close()
			return true
		})
	})
	return nil
}
