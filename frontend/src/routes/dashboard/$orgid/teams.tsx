import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dashboard/$orgid/teams')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/dashboard/$orgid/teams"!</div>
}
