package configurations

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestInitEnvConfig(t *testing.T) {
	// Set test environment variables
	os.Setenv("SCOPE_TEST_ENV_VAR", "test_value")
	defer os.Unsetenv("SCOPE_TEST_ENV_VAR")

	// Call the function being tested
	err := InitEnvConfig()
	if err != nil {
		t.Errorf("InitEnvConfig() returned an unexpected error: %v", err)
	}

	// Check if the environment variable was loaded correctly
	value := viper.GetString("test.env.var")
	if value != "test_value" {
		t.Errorf("Expected TEST_ENV_VAR to be 'test_value', but got '%s'", value)
	}
}
