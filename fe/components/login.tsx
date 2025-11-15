"use client"; 

import { useState } from "react";
import { loginUser } from "@/lib/api/auth";
import { useAuthStore } from "@/stores/useAuthStore";
import { useRouter, useSearchParams } from "next/navigation";
import Link from "next/link";

export default function LoginFormContent() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const { login } = useAuthStore();
  const router = useRouter();
  const searchParams = useSearchParams(); 

  const redirectPath = searchParams.get("redirect") || "/quizzes";

  const handleLogin = async () => {
    setError("");
    try {
      const { token, user } = await loginUser({ email, password });
      login(token, user);
      router.push(redirectPath);
    } catch (err: unknown) {
      let errorMessage = "Login failed. Please try again.";

      if (err instanceof Error) {
        if (err.message.includes("Invalid credentials")) {
          errorMessage = "Incorrect email or password.";
        } else {
          errorMessage = err.message;
        }
      }
      setError(errorMessage);
    }
  };

  return (
    <div className="bg-[url('/background/navbar.jpg')] p-24 border-4 rounded-4xl border-[#38325F] space-y-16">
      <h1 className="text-[#38325F] font-black text-4xl text-center">
        WELCOME BACK!
      </h1>
      <div className="flex flex-col gap-6 w-full max-w-xs mx-auto">
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="bg-white text-[#38325F] focus:border border-[#38325F] px-4 py-3 rounded-2xl text-lg focus:outline-none placeholder-[#A0AEC0]"
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="bg-white text-[#38325F] focus:border border-[#38325F] px-4 py-3 rounded-2xl text-lg focus:outline-none placeholder-[#A0AEC0]"
        />
        {error && (
          <p className="text-red-500 text-sm text-center">{error}</p>
        )}
        <button
          onClick={handleLogin}
          className="bg-black hover:bg-[#48426E] text-xl text-white font-semibold py-4 px-6 rounded-2xl shadow transition"
        >
          Login
        </button>
        <Link href="/register" className="text-center text-[#38325F]">
          Don&apos;t have an account?{" "}
          <span className="font-black hover:text-black">Sign Up</span>
        </Link>
      </div>
    </div>
  );
}