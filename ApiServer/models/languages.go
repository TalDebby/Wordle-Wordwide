package models

type Words []string

type Languages struct {
	Languages []Language `json:"languages"`
}

type Language struct {
	Language  string `json:"language"`
	Direction string `json:"writing_direction"`
	Words     Words  `json:"words"`
}

type LanguageItem struct {
	Name             string `json:"language"`
	WritingDirection string `json:"writing_direction"`
}

type LanguagesList struct {
	Languages []LanguageItem `json:"languages"`
}
