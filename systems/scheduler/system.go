package scheduler

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/FDUTCH/Frame"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

// System manages tasks.
// New instance can be created with new(scheduler.System).
type System struct {
	mu      sync.Mutex
	players map[*world.EntityHandle]*manager

	done func()
}

func (s *System) Close() error {
	if s.done != nil {
		s.done()
	}
	return nil
}

func (s *System) Init(f *Frame.Frame) {
	s.players = make(map[*world.EntityHandle]*manager)

	ctx, cancel := context.WithCancel(context.Background())
	s.done = cancel
	go s.startTicking(ctx)
}

func (s *System) startTicking(ctx context.Context) {
	ticker := time.NewTicker(time.Second / 20)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.tick(ctx)
		}
	}
}

func (s *System) tick(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for h, m := range s.players {
		ok, err := player.Call(ctx, h, m.tick)
		if (err != nil && !errors.Is(err, world.ErrEntityNotInWorld)) || !ok {
			delete(s.players, h)
		}
	}
}

// Schedule returns Task, that can be configured with Option functions.
func (s *System) Schedule(target *world.EntityHandle, duration time.Duration, options ...Option) *Task {
	now := time.Now()
	t := &Task{
		start: now,
		date:  now.Add(duration),
	}

	for _, opt := range options {
		opt(t)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.players[target]
	if !ok {
		m = newManager()
		s.players[target] = m
	}
	m.addTask(t)
	return t
}
