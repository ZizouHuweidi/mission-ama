package middleware

import (
	"os"
	"testing"

	"github.com/zizouhuweidi/mission-ama/config"
	"github.com/zizouhuweidi/mission-ama/ent"
	"github.com/zizouhuweidi/mission-ama/pkg/services"
	"github.com/zizouhuweidi/mission-ama/pkg/tests"
)

var (
	c   *services.Container
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = services.NewContainer()

	// Create a user
	var err error
	if usr, err = tests.CreateUser(c.ORM); err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err = c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
