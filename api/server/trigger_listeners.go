package server

import (
	"context"

	"github.com/dukhyungkim/fn/api/models"
	"github.com/dukhyungkim/fn/fnext"
)

type triggerListeners []fnext.TriggerListener

var _ fnext.TriggerListener = new(triggerListeners)

func (t *triggerListeners) BeforeTriggerCreate(ctx context.Context, trigger *models.Trigger) error {
	for _, l := range *t {
		err := l.BeforeTriggerCreate(ctx, trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *triggerListeners) AfterTriggerCreate(ctx context.Context, trigger *models.Trigger) error {
	for _, l := range *t {
		err := l.AfterTriggerCreate(ctx, trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *triggerListeners) BeforeTriggerUpdate(ctx context.Context, trigger *models.Trigger) error {
	for _, l := range *t {
		err := l.BeforeTriggerUpdate(ctx, trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *triggerListeners) AfterTriggerUpdate(ctx context.Context, trigger *models.Trigger) error {
	for _, l := range *t {
		err := l.AfterTriggerUpdate(ctx, trigger)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *triggerListeners) BeforeTriggerDelete(ctx context.Context, triggerID string) error {
	for _, l := range *t {
		err := l.BeforeTriggerDelete(ctx, triggerID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *triggerListeners) AfterTriggerDelete(ctx context.Context, triggerID string) error {
	for _, l := range *t {
		err := l.AfterTriggerDelete(ctx, triggerID)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddTriggerListener adds an TriggerListener for the server to use.
func (s *Server) AddTriggerListener(listener fnext.TriggerListener) {
	*s.triggerListeners = append(*s.triggerListeners, listener)
}
