
use extism_pdk::*;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Resource {
    uri: String,
    name: String,
    description: String,
    #[serde(rename = "mimeType")]
    mime_type: String,
    text: Option<String>,
    blob: Option<String>,
}

#[derive(Serialize, Deserialize)]
struct ResourceContent {
    uri: String,
    #[serde(rename = "mimeType")]
    mime_type: String,
    text: Option<String>,
    blob: Option<String>,
}

#[plugin_fn]
pub fn resources_information() -> FnResult<String> {
    let resources = vec![
        Resource {
            uri: "sea-flea:///readme".to_string(),
            name: "readme".to_string(),
            description: "Sea Flea documentation".to_string(),
            mime_type: "text/markdown".to_string(),
            text: None,
            blob: None,
        },
    ];

    let json_data = serde_json::to_string(&resources)?;
    Ok(json_data)
}


#[plugin_fn]
pub fn readme() -> FnResult<String> {
    let content = ResourceContent {
        uri: "sea-flea:///readme".to_string(),
        mime_type: "text/markdown".to_string(),
        text: Some(r#"# Sea Flea - MCP WASM Runner

Sea Flea is an MCP (Model Context Protocol) server that supports WebAssembly (WASM) plugins. Plugins can provide three types of capabilities:

- **Tools**: Functions that can be called with arguments
- **Resources**: Static content accessible via URIs  (dynamic content is not yet implemented)
- **Prompts**: Templates for generating conversation prompts        

"#.to_string()),
        blob: None,
    };

    let json_data = serde_json::to_string(&content)?;
    Ok(json_data)
}
