import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/select-organization')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/select-organization"!</div>
}
