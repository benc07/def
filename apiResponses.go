package main

type Definition []struct {
	Word      string `json:"word"`
	Phonetic  string `json:"phonetic,omitempty"`
	Phonetics []struct {
		Text      string `json:"text"`
		Audio     string `json:"audio"`
		SourceURL string `json:"sourceUrl,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Synonyms   []any  `json:"synonyms"`
			Antonyms   []any  `json:"antonyms"`
		} `json:"definitions"`
		Synonyms []any `json:"synonyms"`
		Antonyms []any `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}
