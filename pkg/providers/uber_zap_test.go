package providers

import (
	"github.com/ereb-or-od/kenobi-logger/pkg/options"
	"testing"
)

func TestNewUberZapLoggerShouldReturnDefaultUberZapLoggerWhenConfigurationDoesNotSelected(t *testing.T) {
	uberZapLogger, err := New()
	if err != nil {
		t.Error("error does not expected")
	}

	if uberZapLogger == nil {
		t.Error("default-logger must be initialized")
	}
}

func TestNewUberZapLoggerWithOptionsShouldReturnUberZapLoggerWhenConfigurationSelected(t *testing.T) {
	uberZapLogger, err := New(options.New())
	if err != nil {
		t.Error("error does not expected")
	}

	if uberZapLogger == nil {
		t.Error("default-logger must be initialized")
	}
}

func TestNewZapConfigShouldReturnDefaultZapConfigWhenConfigurationSelectedAsDefault(t *testing.T) {
	config := newZapConfig(options.New())
	if config == nil {
		t.Error("config should be initialized")
	}
}
