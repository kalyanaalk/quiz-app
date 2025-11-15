"use client";
import { useEffect, useState } from "react";
import { useRouter, usePathname } from "next/navigation";
import { useAuthStore } from "@/stores/useAuthStore";

export const ProtectedLayout = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const { token } = useAuthStore();
  const router = useRouter();
  const pathname = usePathname();
  const [hasHydrated, setHasHydrated] = useState(false);

  useEffect(() => {
    setHasHydrated(true);
  }, []);

  useEffect(() => {
    if (hasHydrated && !token) {
      router.replace(`/login?redirect=${pathname}`);
    }
  }, [token, hasHydrated, pathname, router]);

  if (!hasHydrated) {
    return null;
  }

  return <>{children}</>;
};