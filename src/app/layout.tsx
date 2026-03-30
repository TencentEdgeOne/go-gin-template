import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Gin + EdgeOne Pages",
  description: "Go Functions allow you to run Go web frameworks like Gin on EdgeOne Pages. Build full-stack applications with Gin's powerful routing, middleware, and JSON handling.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en-US">
      <head>
        <link rel="icon" href="/gin-favicon.svg" />
      </head>
      <body
        className="antialiased"
      >
        {children}
      </body>
    </html>
  );
}
