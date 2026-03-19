import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/dashboard/$orgid/inventory/')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/dashboard/$orgid/inventory/"!</div>
}
