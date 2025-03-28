import { ThemeProvider } from "@/providers/theme-provider";

interface RootLayoutProps {
  children: React.ReactNode;
}

export function RootLayout({ children }: Readonly<RootLayoutProps>) {
  return (
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <div className="w-full px-3 my-3">{children}</div>
    </ThemeProvider>
  );
}
