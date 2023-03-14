package datastore

import (
	"testing"

	"github.com/dukhyungkim/fn/api/datastore/datastoretest"
	"github.com/dukhyungkim/fn/api/models"
)

func TestDatastore(t *testing.T) {
	f := func(t *testing.T) models.Datastore {
		return NewMock()
	}
	datastoretest.RunAllTests(t, f, datastoretest.NewBasicResourceProvider())
}
