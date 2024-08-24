import React from 'react';
import { Report, Row } from './interface';

type BalanceSheetReportProps = {
  report: Report;
};

/**
 * BalanceSheetReport Component displays balance sheet report with
 * headers, section titles, and rows of financial data.
 */
const BalanceSheetReport: React.FC<BalanceSheetReportProps> = ({ report }) => {
  const { rows, titles } = report;

  const headerRow = rows?.find((item) => item.type === 'Header');
  const rowSections = rows?.filter((item) => item.type === 'Section');

  return (
    <div className="content-center pb-10">
      {titles.map((title, index) => (
        <h1 data-testid="report-title" key={index}>
          {title}
        </h1>
      ))}
      {rows?.length && (
        <table className="table-auto border-collapse mt-10">
          <thead>
            <tr data-testid="header-row">
              {headerRow?.cells?.map((item, index) => (
                <th key={`${item.value}-${index}`} className={'font-normal p-2 bg-blue-100'}>
                  {item.value}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {rowSections?.map((row, index) => (
              <RowSection key={index} row={row} />
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

type RowSectionProps = {
  row: Row;
};
const RowSection = ({ row }: RowSectionProps) => {
  /**
   * Some row sections are heading for that section that contains only title without rows
   * We need to make those section headers a bit bigger
   */
  if (!row?.rows?.length && row.title) {
    return (
      <tr data-testid={'row-section'}>
        <td className="text-2xl pt-4">{row.title}</td>
      </tr>
    );
  }

  /**
   * Rows can have rows within them. Recursively call the same component in such cases.
   */
  if (!!row.rows?.length) {
    return (
      <>
        {row.title && (
          <tr className="pt-2">
            <td className="font-semibold pt-2">{row.title}</td>
          </tr>
        )}
        {row.rows.map((row, index) => <RowSection key={index} row={row} />)}
      </>
    );
  }

  return (
    <tr
      data-testid={'row-section'}
      className={row.type === 'SummaryRow' ? 'font-semibold' : 'px-4 py-2 text-blue-700'}
    >
      {row?.cells?.map((cell, index) => (
        <td key={index} className="border px-8 text-sm">{cell.value}</td>
      ))}
    </tr>
  );
};

export default BalanceSheetReport;
