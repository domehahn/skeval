package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/domehahn/skeval/pkg/generator"
)

const version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "version":
		fmt.Printf("skeval version %s\n", version)
	case "gen":
		handleGen(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`skeval — Synthetic Behavioral & Adversarial Evaluation Generator for AI Agent Skills

Usage:
  skeval gen --skill <name> --tools <tool1,tool2>
  skeval version`)
}

func handleGen(args []string) {
	fs := flag.NewFlagSet("gen", flag.ExitOnError)
	skillName := fs.String("skill", "sample-skill", "Skill name")
	toolsStr := fs.String("tools", "", "Comma-separated list of declared tools")
	_ = fs.Parse(args)

	var tools []string
	if *toolsStr != "" {
		for _, s := range strings.Split(*toolsStr, ",") {
			tools = append(tools, strings.TrimSpace(s))
		}
	}

	suite := generator.GenerateSuite(*skillName, tools)
	out, _ := json.MarshalIndent(suite, "", "  ")
	fmt.Println(string(out))
}
