package config

// ---------------- config: config.go ------------------------------------------
// # Package Documentation
// - Package Name: config
// - Description:  sets configuration variables using a single YAML file.
//   Values in key names can be type values or OS environment variables
//   using ${value} notation. On startup, if _APP_ENV OS environment
//   variable exist, it will select the config.{env}.yaml, otherwise it loads
//   config.yaml.
//
// ## Package File(s)
// - config.go: contains config logic
// - bootsplash.go: bootup logo with config/debug info
//
// ### Current File
// - config.go
//
// #### Global VARIABLE(s)/CONSTANT(s)
// omit this section unless there is a reason to bring up attention to detail
// - {{var1}}: {{var1_describe}}
// - {{const1}}: {{const1_description}}
//
// #### Data Structure(s)
// - {{datastruct1}}: {{datastruct1_describe}}
// - {{datastruct2}}: {{datastruct2_describe}}
//
// #### Constructor(s)
// - NewConfig: Loads the configuration from the YAML file and environment
//   variables.
//
// #### Method(s)
// - interpolateEnvVars(): replaces placeholders in the `FileConfig` struct
//   with env variables.
//
// #### Function(s)
// - ReplaceEnvVars:
//
// **Developer Note(s)**
// **Developer Note(s)**
// - Sample YAML config: config.yaml
/*
---
app:
  name: __MODULE_NAME__
  version: "0.1.0"
  commit_sha: ${_APP_COMMIT_SHA}
  build_id: ${_APP_BUILD_ID}
  build_date: ${_APP_BUILD_DATE}
  env: ${_APP_ENV}
  debug: false

features:
  enable_auth: false
  enable_cache: false
  feature_flag1: false

database:
  driver: "postgres"
  dsn: "postgres://user:pass@localhost:5432/dbname"
  max_connections: 10

server:
  host: 0.0.0.0
  port: 8080
  health_endpoint: /healthz
  info_endpoint: /info
---
*/
// -----------------------------------------------------------------------------

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

// {{blank_line}}
// ================ GLOBAL(s) / CONSTANT(s) ====================================
// {{blank_line}}
var configFileName = determineConfigFile()
var ErrFileNotFound = errors.New("config file not found")
var ErrFileIsEmpty = errors.New("config file is empty")

// determineConfigFile(): selects the appropriate configuration file based
// on _APP_ENV
func determineConfigFile() string {
	if env, exists := os.LookupEnv("_APP_ENV"); exists && env != "" {
		return fmt.Sprintf("./config/config.%s.yaml", env)
	}
	return "./config/config.yaml"
}

// {{blank_line}}
// ================ DATA STRUCTURE(s) / INTERFACE(s) ===========================
// {{blank_line}}
// Config{}: stores application-wide configuration loaded from the YAML file
type Config struct {
	DebugFlag      bool
	Hostname       string
	AppConfig      AppConfig      `yaml:"app"`
	FeatureConfig  FeatureConfig  `yaml:"features"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
	ServerConfig   ServerConfig   `yaml:"server"`
}

type AppConfig struct {
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	CommitSHA string `yaml:"commit_sha"`
	BuildID   string `yaml:"build_id"`
	BuildDate string `yaml:"build_date"`
	Env       string `yaml:"env"`
	Debug     bool   `yaml:"debug"`
}

type FeatureConfig struct {
	EnableAuth   bool `yaml:"enable_auth"`
	EnableCache  bool `yaml:"enable_cache"`
	FeatureFlag1 bool `yaml:"feature_flag1"`
}

type DatabaseConfig struct {
	Driver         string `yaml:"driver"`
	DSN            string `yaml:"dsn"`
	MaxConnections int    `yaml:"max_connections"`
}

type ServerConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	HealthEndpoint string `yaml:"health_endpoint"`
	InfoEndpoint   string `yaml:"info_endpoint"`
}

// {{blank_line}}
// ================ CONSTRUCTOR(s) =============================================
// {{blank_line}}

// NewConfig(): loads the configuration from the YAML file and
// environment variables, and returns a pointer to a Config
func NewConfig() (*Config, error) {
	cfg := &Config{}

	data, err := readFile(configFileName)
	if err != nil {
		return nil, err
	}

	// unmarshal the data into the Config struct
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	// replace environment variables if placeholders are found
	cfg.interpolateEnvVars()

	return cfg, nil
}

// {{blank_line}}
// ================ PUBLIC METHOD(s) ===========================================
// {{blank_line}}
// DebugFlagSet(): returns the string value of the DebugFlag setting from args[]
func (c *Config) DebugFlagSet() {
	c.DebugFlag = true
}

// {{blank_line}}
// ================ PRIVATE METHOD(s) ==========================================
// {{blank_line}}
// interpolateEnvVars(): replaces placeholders within the `Config` struct
// with actual environment variable values
func (c *Config) interpolateEnvVars() {
	v := reflect.ValueOf(c).Elem()
	replaceEnvVars(v)
}

// {{blank_line}}
// ================ PUBLIC FUNCTION(s) =========================================
// {{blank_line}}
// readFile(): reads in file for constructor
func readFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("readFile: cannot read '%s': %w", fileName, ErrFileNotFound)
	}

	if len(data) == 0 {
		return nil, ErrFileIsEmpty
	}

	return data, nil
}

// replaceEnvVars(): recursively replaces placeholder strings in
// struct fields with corresponding environment variable values
func replaceEnvVars(v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.Struct {
			replaceEnvVars(field)
		} else if field.Kind() == reflect.String {
			fieldValue := field.String()
			if strings.HasPrefix(fieldValue, "${") && strings.HasSuffix(fieldValue, "}") {
				envVarName := fieldValue[2 : len(fieldValue)-1]
				envVarValue := os.Getenv(envVarName)
				if envVarValue != "" {
					field.SetString(envVarValue)
				} else {
					log.Fatalf("🟥 STARTUP ERROR: Missing environment variable: %s\n", envVarName)
				}
			}
		}
	}
}

// CheckConfig(): verifies the presence of required configuration values.
// If the value is empty or "missing", sets the error flag and returns "missing"
func CheckConfig(value string, errorFlag *bool) string {
	if value == "" {
		*errorFlag = true
		return "missing"
	}
	if value == "missing" {
		*errorFlag = true
		return "missing"
	}
	return value
}
