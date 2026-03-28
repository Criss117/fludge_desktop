import type { Operator } from "@iam/domain/entities/operator.entity";
import type { responses } from "@wails/go/models";

export function mapOperatorToDomain(operator: responses.Operator): Operator {
  return {
    id: operator.id,
    name: operator.name,
    email: operator.email,
    username: operator.username,
    operatorType: operator.operatorType as "ROOT" | "EMPLOYEE",
    createdAt: operator.createdAt,
    updatedAt: operator.updatedAt,
    deletedAt: operator.deletedAt,
  };
}
