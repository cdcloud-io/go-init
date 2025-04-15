package config

// ---------------- config: bootslash.go ---------------------------------------
// # Package Documentation
// - Package Name: config
// - Description:  config
//
// ## Package File(s)
// - config.go: core structs, contructors, method to load a Config
// - bootsplash.go: bootsplash on startup
//
// ### Current File
// - bootsplash.go: (current_file)
//
// -----------------------------------------------------------------------------

import (
	"fmt"
	"strings"
)

// {{blank_line}}
// ================ GLOBAL(s) / CONSTANT(s) ====================================
// {{blank_line}}

// ANSI color codes
const (
	Green  = "\033[32m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"

	WhiteOnBlue = "\033[44;97m"

	Reset = "\033[0m"
)

// {{blank_line}}
// ================ PUBLIC FUNCTION(s) =========================================
// {{blank_line}}
func PrintStartupBanner(version string) {
	fmt.Println("")
	fmt.Println(WhiteOnBlue + "--- ☾ci LABS 🔬 ---" + Reset)
	fmt.Println("")
	fmt.Println("  🟦  Starting {{app_name}} ™️  🟦")
	fmt.Println(strings.Repeat("-", 79))
	fmt.Println("  💾  v:", version)
	fmt.Println("  ©️  CI LABS (cilabs.dev)")
}

func PrintStartupConfig(c *Config) {
	fmt.Println("  📃  Configuration Settings")
	fmt.Println(strings.Repeat("-", 79))
	fmt.Println("  Env                  : ", c.AppConfig.Env)
	fmt.Println("  Debug                : ", c.AppConfig.Debug)
	fmt.Println(strings.Repeat("-", 79))
}
