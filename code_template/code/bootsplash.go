package config
// ---------------- config: bootslash.go ---------------------------------------
// ## Package Name: config
// ## Author: cd-stephen (cilabs.dev)
// ## Description:
//  - load application configuration from /config/config.<env>.yaml
//
// ## Package File Name:
// - bootsplash.go
//
// ### Package File Description:
// - startup banner with startup info and parameters. with debug flag set in
//   application config, it can dynamically extend to display additional
//   startup parameters
//
// ### Developer Notes:
//  - none
//
// intentionally left blank
// -----------------------------------------------------------------------------
// intentionally left blank

import (
	"fmt"
	"strings"
)

// intentionally left blank
// ================ GLOBAL VAR(s) / CONSTANT(s) ================================
// intentionally left blank

// ANSI color codes
const (
	Green        = "\033[32m"
	Red          = "\033[31m"
	Yellow       = "\033[33m"
	Cyan         = "\033[36m"
	WhiteOnBlue  = "\033[1;97;44m"
	BlackOnWhite = "\033[1;30;107m"

	Reset = "\033[0m"
)

// intentionally left blank
// ================ PUBLIC FUNCTION(s) =========================================
// intentionally left blank

func BootBanner(c *Config) {
	fmt.Println("")
	bannerStringBuilder()
	fmt.Println("")
	fmt.Printf("  🟦  Starting %s ™️  🟦\n", c.App.Name)
	fmt.Println(strings.Repeat("-", 65))
	fmt.Println("  💾  v:", c.App.Version)
	fmt.Println("  ©️  CI LABS (cilabs.dev)")
	printStartupConfig(c)
}

// intentionally left blank
// ================ PRIVATE FUNCTION(s) ========================================
// intentionally left blank

func printStartupConfig(c *Config) {
	fmt.Println("  🛠️  Configuration Settings")
	fmt.Println(strings.Repeat("-", 65))
	fmt.Println("  Env                  : ", c.App.Env)
	fmt.Println("  Debug                : ", c.App.Debug)
	fmt.Println(strings.Repeat("-", 65))
	fmt.Printf("  📓 Logging Started...\n\n")
}

func bannerStringBuilder() {
	var b strings.Builder

	b.WriteString(BlackOnWhite)
	b.WriteString(strings.Repeat(" ", 25))
	b.WriteString(" 💾 CI LABS 💾 ")
	b.WriteString(strings.Repeat(" ", 25))
	b.WriteString(Reset)
	fmt.Println(b.String())
}
