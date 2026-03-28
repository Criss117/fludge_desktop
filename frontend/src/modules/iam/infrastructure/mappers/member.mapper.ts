import type { responses } from "@wails/go/models";
import type { Member } from "@iam/domain/entities/member.entity";

export function mapMemberToDomain(member: responses.Member): Member {
  return {
    id: member.id,
    role: member.role,
    operatorId: member.operatorId,
    organizationId: member.organizationId,
    createdAt: member.createdAt,
    updatedAt: member.updatedAt,
    deletedAt: member.deletedAt,
  };
}
