package es_queries

type ESQueries struct {
	Equals []FieldValue `json::"equals"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
