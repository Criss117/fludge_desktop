import { StrictMode } from "react";
import "./global.css";

import ReactDOM from "react-dom/client";
import {
  RouterProvider,
  createRouter,
  createHashHistory,
} from "@tanstack/react-router";
// Import the generated route tree
import { routeTree } from "./routeTree.gen";
import { queryClient } from "@/integrations/ts-query/index";
import { Integrations } from "./integrations";
// import { useAuth } from "./integrations/auth";

const hashHistory = createHashHistory();

// Create a new router instance
const router = createRouter({
  routeTree,
  history: hashHistory,
  context: {
    queryClient,
  },
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function Main() {
  return <RouterProvider router={router} />;
}

// Render the app
const rootElement = document.getElementById("root")!;
if (!rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <StrictMode>
      <Integrations>
        <Main />
      </Integrations>
    </StrictMode>,
  );
}
