package model

type Texts []Localized

type Localized struct {
    Lang    string  `json:"lang"`
    Text    string  `json:"text"`
}
