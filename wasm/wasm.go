package wasm

import (
	"context"
	"fmt"
	"os"
	"sea-flea/mcp"
	"sea-flea/utils"
	"strings"

	"maps"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero"
)

// TODO : check if we need a mutex for the plugins
// TODO : check errors handling and logs

func GetEnvVariableStartingWith(prefix string) map[string]string {
	envVars := map[string]string{}
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, prefix) {
			parts := strings.SplitN(env, "=", 2)
			if len(parts) == 2 {
				envVars[parts[0]] = parts[1]
			}
		}
	}
	return envVars
}

func LoadPlugins(server *mcp.MCPServer) {

	//fmt.Println("🔥", GetEnvVariableStartingWith("WASM_"))

	ctx := context.Background()

	// Load plugins from the specified path
	pluginConfig := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
	}

	// *-------------------
	// * 001-FILTER
	// *-------------------
	//! use it to filter the plugins: wasmFiles
	utils.Log(func() string {
		if len(server.FilterItems()) == 0 {
			return "🧩 no filter applied"
		}
		items := "🧩"
		for _, item := range server.FilterItems() {
			items += " " + item + "/"
		}
		return items
	}, server.LogOutput())
	//! to be removed

	// List all  wasm files in the cfg.PluginsPath path
	wasmFiles, err := os.ReadDir(server.PluginsPath())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading plugins directory: %v\n", err)
		os.Exit(1)
	}

	config := map[string]string{}

	// add the content of the env vars to the config variable
	wasmEnvVars := GetEnvVariableStartingWith("WASM_")
	if len(wasmEnvVars) > 0 {
		maps.Copy(config, wasmEnvVars)
	}

	// add the content of the settings to the config variable
	settings, err := server.PluginsSettings()
	if err != nil {
		// Log the output
		utils.Log(func() string {
			return "😡 when reading the plugins settings:" + err.Error()
		}, server.LogOutput())
	}
	if settings != nil {
		maps.Copy(config, settings)
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
				Config:       config,
				//Config:       GetEnvVariableStartingWith("WASM_"),
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
