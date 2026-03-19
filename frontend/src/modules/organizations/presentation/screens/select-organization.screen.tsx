import type { responses } from "@wails/go/models";

interface Props {
  organizations: responses.OperatorOrganizationResponse[];
}

export function SelectOrganizationScreen({ organizations }: Props) {
  return (
    <pre>
      <code>{JSON.stringify(organizations, null, 2)}</code>
    </pre>
  );
}
