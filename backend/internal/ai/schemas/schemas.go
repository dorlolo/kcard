package schemas

var KnowledgeClassification = []byte(`{"type":"object","properties":{"knowledgePoints":{"type":"array"}}}`)
var DeckGeneration = []byte(`{"type":"object","properties":{"decks":{"type":"array"},"cards":{"type":"array"}}}`)
var ReviewPlan = []byte(`{"type":"object","properties":{"days":{"type":"array"}}}`)
var PlanOptimization = []byte(`{"type":"object","properties":{"changes":{"type":"array"}}}`)
