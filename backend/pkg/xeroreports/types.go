package xeroreports

type Report struct {
	ReportID       string    `json:"ReportID,omitempty"`
	ReportName     string    `json:"ReportName,omitempty"`
	ReportType     string    `json:"ReportType,omitempty"`
	ReportTitles   *[]string `json:"ReportTitles,omitempty"`
	ReportDate     string    `json:"ReportDate,omitempty"`
	UpdatedDateUTC string    `json:"UpdatedDateUTC,omitempty"`
	Rows           *[]Row    `json:"Rows,omitempty"`
}

// Row is a Row on a Report
type Row struct {
	RowType string  `json:"RowType,omitempty"`
	Title   string  `json:"Title,omitempty"`
	Rows    *[]Row  `json:"Rows,omitempty"`
	Cells   *[]Cell `json:"Cells,omitempty"`
}

type Cell struct {
	Value      string       `json:"Value,omitempty"`
	Attributes *[]Attribute `json:"Attributes,omitempty"`
}

type Attribute struct {
	Value string `json:"Value,omitempty"`
	ID    string `json:"Id,omitempty"`
}
