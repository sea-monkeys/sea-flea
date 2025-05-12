package wasm

import (
	"context"
	"fmt"
	"os"
	"sea-flea/mcp"
	"strings"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero"
)

// TODO : check if we need a mutex for the plugins
// TODO : check errors handling and logs

func LoadPlugins(server *mcp.MCPServer) {
	ctx := context.Background()

	// Load plugins from the specified path
	pluginConfig := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
	}

	// List all  wasm files in the cfg.PluginsPath path
	wasmFiles, err := os.ReadDir(server.PluginsPath())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading plugins directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range wasmFiles {
		if strings.HasSuffix(file.Name(), ".wasm") {
			wasmFilePath := fmt.Sprintf("%s/%s", server.PluginsPath(), file.Name())

			manifest := extism.Manifest{
				Wasm: []extism.Wasm{
					extism.WasmFile{
						Path: wasmFilePath,
					},
				},
				AllowedHosts: []string{"*"},
				Config:       map[string]string{},
			}

			pluginInst, err := extism.NewPlugin(ctx, manifest, pluginConfig, nil) // new
			if err != nil {
				// Handle error case
				fmt.Fprintf(os.Stderr, "Error loading plugin: %v\n", err)
				return
			}

			if pluginInst.FunctionExists("tools_information") {
				// TODO : return error
				registerToolsOfThePlugin(server, pluginInst)
			}

			if pluginInst.FunctionExists("resources_information") {
				// TODO : return error
				registerResourcesOfThePlugin(server, pluginInst)
			}
			
			if pluginInst.FunctionExists("prompts_information") {
				// TODO : return error
				registerPromptsOfThePlugin(server, pluginInst)
			}

		}
	}

}
