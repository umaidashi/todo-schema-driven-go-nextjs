package model

type Priority struct {
	Name    string `json:"name"` // DB に保存される値
	Label   string `json:"label"`
	Color   string `json:"color"`
	BgColor string `json:"bgColor"`
}
