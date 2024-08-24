export interface Report {
  id: string;
  name: string;
  type: string;
  titles: string[];
  date: string;
  rows?: Row[];
}

export interface Row {
  type: string;
  title?: string;
  rows?: Row[];
  cells?: Cell[];
}

export interface Cell {
  value: string;
  attributes?: Attribute[];
}

export interface Attribute {
  value?: string;
  id?: string;
}