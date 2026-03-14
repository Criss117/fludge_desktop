import {
  QueryCache,
  QueryClient,
  QueryClientProvider,
  MutationCache,
} from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: Infinity, // Los datos locales no se "vuelven obsoletos" solos
      gcTime: Infinity, // Mantener caché mientras la app esté abierta
      refetchOnWindowFocus: false, // No aplica en mobile
      refetchOnReconnect: false, // No hay servidor al que reconectar
      refetchOnMount: false, // Evitar lecturas innecesarias de SQLite
      retry: false, // Errores de SQLite no se resuelven reintentando
      networkMode: "always", // No depende de conexión a internet
    },
    mutations: {
      retry: false,
      networkMode: "always",
    },
  },

  queryCache: new QueryCache({
    onError: (error: unknown, query) => {
      console.error(
        `[Query error] key: ${JSON.stringify(query.queryKey)}`,
        error,
      );
    },
  }),

  mutationCache: new MutationCache({
    onError: (error: unknown, _variables, _context, mutation) => {
      console.error(
        `[Mutation error] key: ${JSON.stringify(mutation.options.mutationKey)}`,
        error,
      );
    },
  }),
});

export function QueryProvider({ children }: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={queryClient}>
      {children}
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}
