//go:build experimental || couchdb

package couchdb

import (
	"testing"

	"github.com/lamassuiot/lamassuiot/core/v3/pkg/config"
	"github.com/lamassuiot/lamassuiot/core/v3/pkg/engines/storage"
	"github.com/lamassuiot/lamassuiot/shared/subsystems/v3/pkg/test/subsystems"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestBuildStorageEngineCouchDB(t *testing.T) {
	builder := subsystems.GetSubsystemBuilder[subsystems.StorageSubsystem](subsystems.CouchDB)
	backend, err := builder.Run()
	if err != nil {
		t.Fatalf("could not run storage subsystem: %s", err)
	}

	t.Cleanup(func() { backend.AfterSuite() })
	logger := log.WithField("test", "BuildStorageEngine_CouchDB")

	conf := backend.Config.(config.PluggableStorageEngine)

	// Call the BuildStorageEngine function
	storageBuilder := storage.GetEngineBuilder(conf.Provider)
	assert.NotNil(t, storageBuilder)

	storageEngine, err := storageBuilder(logger, conf)
	assert.NoError(t, err)

	assert.Equal(t, storageEngine.GetProvider(), config.CouchDB)
}
