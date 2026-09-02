package models

type JSONValueIn struct {
	Value  string `json:"value"`
	Source string `json:"source"`
	Type   string `json:"type"`
	Label  string `json:"label"`
}

type JSONValueOut struct {
	UUID   string `json:"uuid"`
	Value  string `json:"value"`
	Source string `json:"source"`
	Type   string `json:"type"`
	Label  string `json:"label"`
}

type JSONRelation struct {
	Relation string   `json:"relation"`
	From     []string `json:"from"`
	To       []string `json:"to"`
}

type NEO4JRelation struct {
	Relation string
	From     []string
	To       []string
}
