package filter

type sort struct {
	ColId string `json:"colId"`
	Sort  string `json:"sort"`
}

type Filter struct {
	Type string `json:"type"`
	From string `json:"from"`
	To   string `json:"to"`

	FilterType string `json:"filterType"`
}

type DynamicFilter struct {
	Sort   *[]sort           `json:"sort"`
	Filter map[string]Filter `json:"filters"`
}
