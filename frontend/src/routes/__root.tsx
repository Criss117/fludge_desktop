import type { QueryClient } from "@tanstack/react-query";
import { createRootRouteWithContext, Outlet } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { Toaster } from "@shared/components/ui/sonner";

import { TooltipProvider } from "@shared/components/ui/tooltip";

export interface RouterAppContext {
  queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<RouterAppContext>()({
  component: RootLayout,
  head: () => ({
    meta: [
      {
        title: "fludge",
      },
      {
        name: "description",
        content: "fludge is a web application",
      },
    ],
    links: [
      {
        rel: "icon",
        href: "/favicon.ico",
      },
    ],
    scripts: [
      {
        crossOrigin: "anonymous",
        src: "//unpkg.com/react-scan/dist/auto.global.js",
      },
    ],
  }),
});

function RootLayout() {
  return (
    <>
      <TooltipProvider>
        <Outlet />
        <Toaster />
      </TooltipProvider>
      <TanStackRouterDevtools />
    </>
  );
}
