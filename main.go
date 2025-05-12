package main

import (
	"fmt"
	"os"
	"sea-flea/cli"
	"sea-flea/demo"
	"sea-flea/mcp"
	"sea-flea/transport"
	"strings"
)

func main() {

	cfg, err := cli.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Handle filters
	filters := strings.SplitSeq(cfg.Filter, ",")
	for filter := range filters {
		filter = strings.TrimSpace(filter)
		//fmt.Printf("Processing filter: %s\n", filter)
		// Add your filter processing logic here
	}

	/*
	fmt.Printf("Configuration:\n")
	fmt.Printf("  Transport: %s\n", cfg.Transport)
	fmt.Printf("  HTTP Port: %d\n", cfg.HTTPPort)
	fmt.Printf("  Debug: %t\n", cfg.Debug)
	fmt.Printf("  Plugins Path: %s\n", cfg.PluginsPath)
	fmt.Printf("  Filter: %s\n", cfg.Filter)
	fmt.Printf("  Load Demo Tools: %t\n", cfg.DemoTools)
	fmt.Printf("  Load Demo Resources: %t\n", cfg.DemoResources)
	fmt.Printf("  Load Demo Prompts: %t\n", cfg.DemoPrompts)
	*/

	// Create server instance
	server := mcp.NewMCPServer(cfg.Debug)

	if cfg.DemoTools {
		demo.LoadTools(server)
	}

	if cfg.DemoResources {
		demo.LoadResources(server)
	}

	if cfg.DemoPrompts {
		demo.LoadPrompts(server)
	}

	// Run the appropriate transport based on the config
	switch cfg.Transport {
	case "stdio":
		transport.STDIO(server)
	case "streamable-http":
		transport.StreamableHTTP(server, cfg.CertFile, cfg.KeyFile)
	default:
		panic("Invalid mcp transport.")
	}

}
