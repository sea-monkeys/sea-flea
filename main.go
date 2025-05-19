package main

import (
	"fmt"
	"os"
	"sea-flea/cli"
	"sea-flea/demo"
	"sea-flea/mcp"
	"sea-flea/transport"
	"sea-flea/utils"
	"sea-flea/wasm"
)

func main() {

	cfg, err := cli.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	//*-------------------
	//* 002-SK-PLUGINS
	//*-------------------
	if cfg.Generate {
		// Generate a new source code plugin
		utils.Log(func() string {
			return "💀 " + cfg.Name +" / " + cfg.Language
		}, cfg.Debug)
		//! to be removed
		os.Exit(0)
	}

	// Create server instance
	server := mcp.NewMCPServer(cfg)

	if cfg.DemoTools {
		demo.LoadTools(server)
	}

	if cfg.DemoResources {
		demo.LoadResources(server)
	}

	if cfg.DemoPrompts {
		demo.LoadPrompts(server)
	}

	// Load WASM plugins if the path is provided
	if cfg.PluginsPath != "" {
		wasm.LoadPlugins(server)
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
