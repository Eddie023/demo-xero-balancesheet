package server

type Report struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Titles []string `json:"titles"`
	Date   string   `json:"date"`
	Rows   *[]Row   `json:"rows,omitempty"`
}

// Row is a Row on a Report
type Row struct {
	Type  string  `json:"type"`
	Title string  `json:"title"`
	Rows  *[]Row  `json:"rows,omitempty"`
	Cells *[]Cell `json:"cells,omitempty"`
}

type Cell struct {
	Value string `json:"value"`
}
