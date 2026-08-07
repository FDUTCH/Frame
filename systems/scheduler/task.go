package scheduler

import (
	"sync/atomic"
	"time"

	"github.com/df-mc/dragonfly/server/player"
)

type Task struct {
	name string

	canceled atomic.Bool

	date, start time.Time

	onTick func(t *Task, pl *player.Player)
	onDone func(pl *player.Player)
}

// End returns deadline for the task.
func (s *Task) End() time.Time {
	return s.date
}

// Cancel cancels the task.
func (s *Task) Cancel() {
	s.canceled.Store(true)
}

// Canceled returns whether the task is canceled.
func (s *Task) Canceled() bool {
	return s.canceled.Load()
}

// ProgressAt returns progress at the moment passed.
func (s *Task) ProgressAt(moment time.Time) float64 {
	diff := s.date.Sub(s.start)
	rel := s.date.Sub(moment)
	return rel.Seconds() / diff.Seconds()
}

// DoneBy returns if the task will be done by the moment passed.
func (s *Task) DoneBy(moment time.Time) bool {
	return s.date.Before(moment)
}

// Name returns the name of the task.
func (s *Task) Name() string {
	return s.name
}

type Option = func(task *Task)

// OptionName sets the name for a Task.
// Tasks with the name will be shown in the action bar of the player.
func OptionName(name string) Option {
	return func(task *Task) {
		task.name = name
	}
}

// OptionTick calls onTick every time task ticks.
func OptionTick(onTick func(t *Task, pl *player.Player)) Option {
	return func(task *Task) {
		task.onTick = onTick
	}
}

// OnDone calls onDone when task is done.
func OnDone(onDone func(pl *player.Player)) Option {
	return func(task *Task) {
		task.onDone = onDone
	}
}
