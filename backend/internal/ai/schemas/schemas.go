// Package schemas 提供 AI 结构化输出所需的 JSON Schema 定义。
package schemas

// KnowledgeClassification 常量定义知识分类输出的 JSON Schema。
var KnowledgeClassification = []byte(`{
  "type": "object",
  "additionalProperties": false,
  "required": ["knowledgePoints", "warnings"],
  "properties": {
    "knowledgePoints": {
      "type": "array",
      "items": {
        "type": "object",
        "additionalProperties": false,
        "required": ["content", "summary", "tags", "confidence"],
        "properties": {
          "content": {"type": "string"},
          "summary": {"type": "string"},
          "tags": {"type": "array", "items": {"type": "string"}},
          "confidence": {"type": "number"}
        }
      }
    },
    "warnings": {"type": "array", "items": {"type": "string"}}
  }
}`)

// DeckGeneration 常量定义卡片组生成输出的 JSON Schema。
var DeckGeneration = []byte(`{
  "type": "object",
  "additionalProperties": false,
  "required": ["decks", "cards", "warnings"],
  "properties": {
    "decks": {"type": "array", "items": {"type": "object", "additionalProperties": true}},
    "cards": {"type": "array", "items": {"type": "object", "additionalProperties": true}},
    "warnings": {"type": "array", "items": {"type": "string"}}
  }
}`)

// ReviewPlan 常量定义复习计划生成输出的 JSON Schema。
var ReviewPlan = []byte(`{
  "type": "object",
  "additionalProperties": false,
  "required": ["days", "warnings"],
  "properties": {
    "days": {"type": "array", "items": {"type": "object", "additionalProperties": true}},
    "warnings": {"type": "array", "items": {"type": "string"}}
  }
}`)

// PlanOptimization 常量定义复习计划优化输出的 JSON Schema。
var PlanOptimization = []byte(`{
  "type": "object",
  "additionalProperties": false,
  "required": ["changes", "warnings"],
  "properties": {
    "changes": {"type": "array", "items": {"type": "object", "additionalProperties": true}},
    "warnings": {"type": "array", "items": {"type": "string"}}
  }
}`)
