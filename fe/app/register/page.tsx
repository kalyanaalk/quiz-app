"use client";
import { useState } from "react";
import { registerUser } from "@/lib/api/auth";
import { useAuthStore } from "@/stores/useAuthStore";
import { useRouter } from "next/navigation";
import Link from "next/link";

export default function RegisterPage() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const { login } = useAuthStore();
  const router = useRouter();

  const validateEmail = (email: string) =>
    /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);

  const handleRegister = async () => {
    setError("");

    if (!username || !email || !password || !confirmPassword) {
      return setError("All fields are required.");
    }

    if (!validateEmail(email)) {
      return setError("Invalid email format.");
    }

    if (password !== confirmPassword) {
      return setError("Passwords do not match.");
    }

    try {
      const { token, user } = await registerUser({ username, email, password });
      login(token, user);
      router.push("/quizzes");
    } catch (err: unknown) {
      let rawError: string = "An unknown error occurred during registration.";
    
      if (err instanceof Error) {
        rawError = err.message;
      }
    
      if (rawError.includes(`uni_users_email`)) {
        setError("This email is already registered.");
      } else if (rawError.includes(`uni_users_username`)) {
        setError("This username is already taken.");
      } else {
        setError(rawError); 
      }
    }
  };

  return (
    <main className="h-screen flex flex-col items-center justify-center bg-[#38325F] text-white px-4">
      <div className="ng-black rounded-4xl p-16">
        <div className="ng-black p-24 border-4 rounded-4xl border-[#38325F] space-y-16">
          <h1 className="text-[#38325F] font-black text-4xl text-center">
            CREATE ACCOUNT
          </h1>
          <div className="flex flex-col gap-4 w-full max-w-xs mx-auto">
            <input
              className="bg-white text-[#38325F] focus:border border-[#38325F] px-4 py-3 rounded-2xl text-lg focus:outline-none placeholder-[#A0AEC0]"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
            <input
              className="bg-white text-[#38325F] focus:border border-[#38325F] px-4 py-3 rounded-2xl text-lg focus:outline-none placeholder-[#A0AEC0]"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <input
              className="bg-white text-[#38325F] focus:border border-[#38325F] px-4 py-3 rounded-2xl text-lg focus:outline-none placeholder-[#A0AEC0]"
              type="password"
              placeholder="Password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <input
              className="bg-white text-[#38325F] focus:border border-[#38325F] px-4 py-3 rounded-2xl text-lg focus:outline-none placeholder-[#A0AEC0]"
              type="password"
              placeholder="Confirm Password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
            {error && (
              <p className="text-red-500 text-sm font-semibold text-center">
                {error}
              </p>
            )}
            <button
              className="bg-black hover:bg-[#48426E] text-xl text-white font-semibold py-4 px-6 rounded-2xl shadow transition"
              onClick={handleRegister}
            >
              Register
            </button>
            <Link href="/login" className="text-center text-[#38325F]">
              Already have an account?{" "}
              <span className="font-black hover:text-black">Sign In</span>
            </Link>
          </div>
        </div>
      </div>
    </main>
  );
}