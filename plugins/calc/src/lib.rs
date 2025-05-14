use std::collections::HashMap;

use extism_pdk::*;
use serde::{Deserialize, Serialize};
use serde_json::{json};

#[derive(Serialize, Deserialize)]
struct Tool {
    #[serde(rename = "name")]
    name: String,
    
    #[serde(rename = "description")]
    description: String,
    
    #[serde(rename = "inputSchema")]
    input_schema: InputSchema,
}

#[derive(Serialize, Deserialize)]
struct InputSchema {
    #[serde(rename = "type")]
    type_field: String,
    
    #[serde(rename = "required")]
    required: Vec<String>,
    
    #[serde(rename = "properties")]
    properties: HashMap<String, serde_json::Value>,
}

#[plugin_fn]
pub fn tools_information() -> FnResult<String> {

    let add = Tool {
        name: "add".to_string(),
        description: "a function to add numbers".to_string(),
        input_schema: InputSchema {
            type_field: "object".to_string(),
            required: vec!["a".to_string(), "b".to_string()],
            properties: HashMap::from([
                ("a".to_string(), json!({
                    "type": "number",
                    "description": "first number to add"
                })),
                ("b".to_string(), json!({
                    "type": "number",
                    "description": "second number to add"
                }))
            ]),
        },
    };

    let tools = vec![add];
    let json_data = serde_json::to_string(&tools)?;
    Ok(json_data)
}

// use json data for inputs and outputs
#[derive(FromBytes, Deserialize, PartialEq, Debug)]
#[encoding(Json)]
struct Add {
    a: i32,
    b: i32,
}

#[derive(ToBytes, Serialize, PartialEq, Debug)]
#[encoding(Json)]
struct Sum {
    value: i32,
}

#[plugin_fn]
pub fn add(input: Add) -> FnResult<Sum> {
    Ok(Sum {
        value: input.a + input.b,
    })
}
