"use client";
import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();

  const handleRegister = () => {
    router.push("/register");
  };

  const handleLogin = () => {
    router.push("/login");
  };

  return (
    <main className="h-screen flex flex-col items-center justify-center bg-black text-white px-4">
      <div className="bg-black rounded-4xl p-16">
        <div className="p-24 border-4 rounded-4xl border-[#38325F] space-y-16">

          <div className="flex flex-col gap-4 w-full max-w-xs">
            <button
              onClick={handleRegister}
              className="bg-[#38325F] hover:bg-[#48426E] text-xl text-white font-semibold py-5 px-6 rounded-2xl shadow transition"
            >
              REGISTER
            </button>
            <button
              onClick={handleLogin}
              className="bg-[#38325F] hover:bg-[#48426E] text-xl text-white font-semibold py-5 px-6 rounded-2xl shadow transition"
            >
              LOGIN
            </button>
          </div>
        </div>
      </div>
    </main>
  );
}