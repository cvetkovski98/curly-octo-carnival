export interface Timestamp {
  seconds: number;
}

export interface Machine {
  id: number;
  name: string;
  serialNumber: string;
  lineId: number;
  lastServicedAt: Timestamp;
  installedAt: Timestamp;
}

export interface ProductionLine {
  name: string;
  factory: string;
  zone: string;
  machines: Array<Machine>;
}

export interface GetInventoryResponse {
  type: string
  lines: Array<ProductionLine>;
}
