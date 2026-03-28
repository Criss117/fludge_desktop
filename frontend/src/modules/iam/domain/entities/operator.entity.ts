export interface Operator {
  id: string;
  name: string;
  email: string;
  username: string;
  operatorType: "ROOT" | "EMPLOYEE";
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
}
