#!/bin/bash
# Test script that demonstrates the proper MCP initialization sequence

# Create a temporary input file with the correct sequence
cat > /tmp/mcp_test_input.jsonl << 'EOF'
{
    "jsonrpc": "2.0", 
    "id": 0, 
    "method": "initialize", 
    "params": {"protocolVersion": "2024-11-05", "capabilities": {}, "clientInfo": {"name": "test", "version": "1.0.0"}}
}
{
    "jsonrpc": "2.0", 
    "method": "notifications/initialized"
}
{
    "jsonrpc": "2.0", 
    "id": 2, 
    "method": "resources/read", 
    "params": {
        "uri": "message:///about"
    }
}
EOF

# Run the server with the input file
echo "---------------------------------------------------------"
echo "Running MCP server with proper initialization sequence..."
echo "---------------------------------------------------------"

# Pipe the input to the server and process output with jq
cat /tmp/mcp_test_input.jsonl | docker run --rm -i sea-flea:demo --debug --plugins ./plugins | jq -c '.' | jq -s '.'

# Clean up
rm /tmp/mcp_test_input.jsonl

echo "---------------------------------------------------------"
echo "Test complete!"
echo "---------------------------------------------------------"
