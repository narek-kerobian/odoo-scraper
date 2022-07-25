package model

type LocalizedList []Localized

type Localized struct {
    Lang    string  `json:"lang"`
    Text    string  `json:"text"`
}
