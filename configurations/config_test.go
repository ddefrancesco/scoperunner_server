package configurations

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestInitEnvConfigString(t *testing.T) {
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

func TestInitEnvConfigBool(t *testing.T) {
	// Set test environment variables
	os.Setenv("SCOPE_TEST_BOOL_ENV_VAR", "false")
	defer os.Unsetenv("SCOPE_TEST_ENV_VAR")

	// Call the function being tested
	err := InitEnvConfig()
	if err != nil {
		t.Errorf("InitEnvConfig() returned an unexpected error: %v", err)
	}

	// Check if the environment variable was loaded correctly
	value := viper.GetBool("test.bool.env.var")
	if value != false {
		t.Errorf("Expected TEST_BOOL_ENV_VARto be 'false', but got '%t'", value)
	}
}

func TestInitTestConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test config test ok",
			wantErr: false,
		},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitTestConfig(); (err != nil) != tt.wantErr {
				t.Errorf("InitTestConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "config test ok",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitConfig(); (err != nil) != tt.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
