import { ProtectedLayout } from "@/components/protected_layout";

export default function QuizzesLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <ProtectedLayout>{children}</ProtectedLayout>;
}