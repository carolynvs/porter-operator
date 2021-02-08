package mage

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

const (
	ProductionRegistry = "ghcr.io/getporter"
	TestRegistry       = "localhost:5000"
)

var (
	Env = getAmbientEnvironment()
)

type Environment struct {
	Name            string
	Registry        string
	ControllerImage string
	OlmBundleImage  string
	AgentImage      string
}

func UseTestEnvironment() {
	os.Setenv("ENV", "test")
	Env = GetTestEnvironment()
}

func UseProductionEnvironment() {
	os.Setenv("ENV", "prod")
	Env = GetProductionEnvironment()
}

func getAmbientEnvironment() Environment {
	name := os.Getenv("ENV")
	switch strings.ToLower(name) {
	case "prod", "production":
		return GetProductionEnvironment()
	case "test", "":
		return GetTestEnvironment()
	default:
		Must(errors.Errorf("ENV=%q is not a valid environment name", name))
	}
	return Environment{}
}

func GetTestEnvironment() Environment {
	return buildEnvironment("test", TestRegistry)
}

func GetProductionEnvironment() Environment {
	return buildEnvironment("production", ProductionRegistry)
}

func buildEnvironment(name string, registry string) Environment {
	return Environment{
		Name:            name,
		Registry:        registry,
		ControllerImage: path.Join(registry, "porterops-controller:canary"),
		AgentImage:      path.Join(registry, "porter:kubernetes-canary"),
		OlmBundleImage:  path.Join(registry, "porter-operator-olm-bundle:v0.1.0"),
	}
}
