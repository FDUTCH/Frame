package scheduler

import (
	"bytes"
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/world"
)

type manager struct {
	tasks []*Task
	buff  *bytes.Buffer
}

func newManager(tasks ...*Task) *manager {
	return &manager{tasks: tasks, buff: bytes.NewBuffer(nil)}
}

func (s *manager) tick(_ *world.Tx, pl *player.Player) (bool, error) {
	var (
		now    = time.Now()
		offset int
	)

	for idx, t := range s.tasks {
		if t.Canceled() {
			offset++
			continue
		}
		if t.DoneBy(now) {
			offset++
			if t.onDone != nil {
				t.onDone(pl)
			}
			t.Cancel()
			continue
		}
		if t.onTick != nil {
			t.onTick(t, pl)
		}

		if t.Name() != "" {
			if s.buff.Len() != 0 {
				s.buff.WriteString("\n")
			}
			s.buff.WriteString(t.Name())
			s.writeProgress(t.ProgressAt(now))
		}

		if offset != 0 {
			s.tasks[idx-offset] = t
		}
	}

	if offset != 0 {
		s.tasks = s.tasks[:len(s.tasks)-offset]
	}
	s.display(pl)

	return len(s.tasks) > 0, nil
}

func (s *manager) addTask(t *Task) {
	s.tasks = append(s.tasks, t)
}

func (s *manager) display(pl *player.Player) {
	defer s.buff.Reset()
	t := title.New().WithActionText(s.buff.String())
	pl.SendTitle(t)
}

func (s *manager) writeProgress(progress float64) {
	s.buff.WriteString(": [")
	count := int(progress * steps)

	for range int(steps) - count {
		s.buff.WriteRune('#')
	}
	for range count {
		s.buff.WriteRune('_')
	}

	s.buff.WriteRune(']')
}

const steps = 20.0
