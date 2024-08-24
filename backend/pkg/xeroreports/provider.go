package xeroreports

import (
	"context"
	"errors"
	"fmt"

	"github.com/eddie023/accounting/pkg/server"
)

type Client interface {
	GetBalanceSheet(ctx context.Context, query map[string]string) (Report, error)
}

type XeroReportsProvider struct {
	client Client
}

func NewXeroReportsProvider(client Client) *XeroReportsProvider {
	return &XeroReportsProvider{
		client: client,
	}
}

func (xr *XeroReportsProvider) BuildReport(ctx context.Context, query map[string]string) (server.Report, error) {
	report, err := xr.client.GetBalanceSheet(ctx, nil)
	if err != nil {
		return server.Report{}, fmt.Errorf("getting balance sheet report from client: %w", err)
	}

	if err = validate(report); err != nil {
		return server.Report{}, fmt.Errorf("failed validation: %w", err)
	}

	return transform(report)
}

// validate runs a sanity check for data obtained by third-party sources as per business requirement.
func validate(report Report) error {
	if report.ReportType != "BalanceSheet" {
		return errors.New("report must be of type 'BalanceSheet'")
	}

	return nil
}

// transform transforms third party data to our internal platform's expected struct type.
// for the purpose of this demonstration, I have removed the redundant 'Report' prefix from each field.
func transform(r Report) (server.Report, error) {
	report := server.Report{
		ID:   r.ReportID,
		Name: r.ReportName,
		Type: r.ReportType,
		Date: r.ReportDate,
	}

	if r.ReportTitles != nil {
		report.Titles = *r.ReportTitles
	}

	if (r.Rows != nil) && (len(*r.Rows) > 0) {
		requiredRows := make([]server.Row, len(*r.Rows))
		for i, row := range *r.Rows {
			requiredRows[i].Title = row.Title
			requiredRows[i].Type = row.RowType
			if row.Rows != nil && (len(*row.Rows) > 0) {
				tr := transformRows(*row.Rows)
				requiredRows[i].Rows = tr
			}

			if row.Cells != nil && (len(*row.Cells) > 0) {
				tc := transformCells(*row.Cells)
				requiredRows[i].Cells = tc
			}
		}

		report.Rows = &requiredRows
	}

	return report, nil
}

func transformRows(rawRows []Row) *[]server.Row {
	rows := make([]server.Row, len(rawRows))
	for i, r := range rawRows {
		rows[i] = server.Row{
			Type:  r.RowType,
			Title: r.Title,
		}

		if r.Rows != nil {
			rs := transformRows(*r.Rows)
			rows[i].Rows = rs
		}

		if r.Cells != nil {
			cs := transformCells(*r.Cells)
			rows[i].Cells = cs
		}
	}

	return &rows
}

func transformCells(rawCells []Cell) *[]server.Cell {
	cells := make([]server.Cell, len(rawCells))

	for i, c := range rawCells {
		cells[i] = server.Cell{
			Value: c.Value,
		}
	}

	return &cells
}
