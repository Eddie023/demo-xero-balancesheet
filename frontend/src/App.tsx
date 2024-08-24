import React, { useEffect } from 'react';

import BalanceSheetReport from './components/BalanceSheetReport';
import { Report } from './components/BalanceSheetReport/interface';

function App() {
  const [report, setReport] = React.useState<Report>();
  const [error, setError] = React.useState<string>('');

  const fetchBalanceSheetReport = async (): Promise<Report | undefined> => {
    try {
      // for the sake of this demo and simplicity I have hardcoded base URL.
      const response = await fetch('http://localhost:8080/api/reports/balance-sheet');
      const data = response.json();

      return data;
    } catch (e: unknown) {
      setError('failed to fetch reports from API');
    }
  };

  useEffect(() => {
    fetchBalanceSheetReport().then((data) => setReport(data));
  }, []);

  if (error) {
    return (
      <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
        <strong className="font-bold">Holy smokes! {'  '}</strong>
        <span className="block sm:inline">{error}</span>
      </div>
    );
  }

  return (
    <div className="container">
      <h1 data-testid="report-name" className="text-2xl font-bold underline">
        {report?.name || ''}
      </h1>
      {!!report && <BalanceSheetReport report={report} />}
    </div>
  );
}

export default App;
