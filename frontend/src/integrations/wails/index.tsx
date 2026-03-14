// src/WailsProvider.tsx
import { useState, useEffect } from "react";
import { EventsOn } from "@wails/runtime";

export function WailsProvider({ children }: { children: React.ReactNode }) {
  const [ready, setReady] = useState(false);

  useEffect(() => {
    EventsOn("wails:loaded", () => {
      setReady(true);
    });

    // fallback por si el evento ya se disparó
    const timeout = setTimeout(() => {
      setReady(true);
    }, 500);

    return () => clearTimeout(timeout);
  }, []);

  if (!ready) return null;

  return <>{children}</>;
}
