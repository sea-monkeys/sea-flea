# Tests if you are working with devcontainer

Start the MCP server:
```bash
docker compose up mcp-server --build --no-log-prefix
```

First, initialize the handshake:
```bash
docker compose up initialize --build --no-log-prefix
```

Then you can test the other JSON RPC requests:

```bash
docker compose up resources-list --build --no-log-prefix
```

```bash
docker compose up tools-list --build --no-log-prefix
```

```bash
docker compose up use-tool-add --build --no-log-prefix
```
