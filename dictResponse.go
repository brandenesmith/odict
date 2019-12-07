package main

// DictResponse is the response expected from the call to the oxford api.
type DictResponse struct {
	Results []DictResult `json:"results"`
}

// DictResult is a query result from the api.
type DictResult struct {
	ID             string         `json:"id"`
	Language       string         `json:"language"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
}

// LexicalEntry is the entry for the specified word resolving to a slice of Entry objects.
type LexicalEntry struct {
	Entries []Entry `json:"entries"`
}

// Entry is the actual dictionary entry for the word.
type Entry struct {
	Etymologies         []string             `json:"etymologies"`
	GrammaticalFeatures []GrammaticalFeature `json:"grammaticalFeatures"`
	HomographNumber     string               `json:"homographNumber"`
	Senses              []Sense              `json:"senses"`
	Language            string               `json:"language"`
	LexicalCategory     string               `json:"lexicalCategory"`
}

// GrammaticalFeature is the actual grammatical features of the word.
type GrammaticalFeature struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// Sense is the slice of definitions, examples and other definitions.
type Sense struct {
	Definitions []string `json:"definitions"`
	Example     []string `json:"examples"`
	Domains     []string `json:"domains"`
	Subsenses   []Sense  `json:"subsenses"`
}
