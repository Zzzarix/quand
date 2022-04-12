package domain

type Question struct {
	Text   string `json:"text"`
	Kind   string `json:"type"`
	Author string `json:"author"`
	// Topic string `json:"topic"`

	// Answer   string            `json:"answer,omitempty"`
	// Variants map[string]string `json:"variants,omitempty"`
}

// IQuestion is intermediate representation
type IQuestion struct {
	Text   string `json:"text"`
	Kind   string `json:"type"`
	Author string `json:"author"`
	// Topic string `json:"topic"`
}

// Question's Kinds
const (
	JustAsk = "JustAsk"
)
