package cli

import (
	"flag"
	"fmt"
)

// Config holds all command line arguments
type Config struct {
	Transport   string
	HTTPPort    int
	Debug       bool
	PluginsPath string
	Filter      string
	DemoTools bool
	DemoResources bool
	DemoPrompts bool
	// cert and key for TLS
	CertFile string
	KeyFile  string
	// Settings for the plugins
	Settings string
}

/*
```bash
mkcert \
-cert-file mcp.amphipod.local.crt \
-key-file mcp.amphipod.local.key \
amphipod.local "*.amphipod.local" localhost 127.0.0.1 ::1
```

*/



// ParseFlags parses command line arguments and returns a Config
func ParseFlags() (*Config, error) {
	cfg := &Config{}

	// Define flags with default values
	flag.StringVar(&cfg.Transport, "transport", "stdio", "Transport type (stdio, streamable-http)")
	flag.IntVar(&cfg.HTTPPort, "http-port", 5050, "HTTP port for streamable-http transport")
	flag.BoolVar(&cfg.Debug, "debug", false, "Enable debug mode")
	flag.StringVar(&cfg.PluginsPath, "plugins", ".", "Path to plugins directory")
	flag.StringVar(&cfg.Filter, "filter", "*.*", "Filter for plugins (e.g., *.js, *.sh, filename.ext)")

	flag.BoolVar(&cfg.DemoTools, "demo-tools", false, "Load demo tools")
	flag.BoolVar(&cfg.DemoResources, "demo-resources", false, "Load demo resources")
	flag.BoolVar(&cfg.DemoPrompts, "demo-prompts", false, "Load demo prompts")

	// TLS flags
	flag.StringVar(&cfg.CertFile, "cert", "", "Path to TLS certificate file")
	flag.StringVar(&cfg.KeyFile, "key", "", "Path to TLS key file")

	// Settings for the plugins
	flag.StringVar(&cfg.Settings, "settings", "", "Settings for the plugins (JSON format)")


	// Parse flags
	flag.Parse()

	// Validate transport
	validTransports := map[string]bool{"stdio": true, "streamable-http": true}
	if !validTransports[cfg.Transport] {
		return nil, fmt.Errorf("invalid transport value: %s (valid values: stdio, streamable-http)", cfg.Transport)
	}

	// Validate filter (basic validation, could be enhanced)
	// This is a simple validation that ensures filter has some non-empty value
	if cfg.Filter == "" {
		return nil, fmt.Errorf("filter cannot be empty")
	}

	return cfg, nil
}
