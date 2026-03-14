declare interface Window {
  go: Record<string, Record<string, (...args: unknown[]) => Promise<unknown>>>;
  wails: unknown;
}
