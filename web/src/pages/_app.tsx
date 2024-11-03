import { QUERY_CLIENT } from "@/libs/react-router";
import "@/styles/globals.css";
import { QueryClientProvider } from "@tanstack/react-query";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <QueryClientProvider client={QUERY_CLIENT}>
      <Component {...pageProps} />
    </QueryClientProvider>
  );
}
