"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { getAllQuizzes } from "@/lib/api/quiz";
import { useQuizStore } from "@/stores/useQuizStore";

export default function QuizzesPage() {
  const router = useRouter();
  const { quizzes, setQuizzes } = useQuizStore();

  useEffect(() => {
    async function load() {
      try {
        const data = await getAllQuizzes();
        setQuizzes(data);
      } catch (err) {
        console.error(err);
      }
    }
    load();
  }, [setQuizzes]);

  return (
    <main className="min-h-screen flex flex-col items-center justify-center bg-[#38325F] text-white px-4">
      <div className="ng-black rounded-4xl p-16 w-full max-w-3xl">
        <div className="ng-black p-24 border-4 rounded-4xl border-[#38325F] space-y-16">
          <h1 className="text-[#38325F] font-black text-4xl text-center">
            QUIZZES
          </h1>

          <div className="space-y-6">
            {quizzes.map((quiz) => (
              <div
                key={quiz.id}
                className="bg-white/10 p-6 rounded-xl cursor-pointer hover:bg-white/20 transition"
                onClick={() => router.push(`/quizzes/${quiz.id}`)}
              >
                <h2 className="text-2xl font-bold">{quiz.title}</h2>
                <p className="opacity-80">{quiz.description}</p>
                <p className="text-sm opacity-50 mt-2">
                  {quiz.total_questions} Questions • {quiz.duration} sec
                </p>
              </div>
            ))}
          </div>
        </div>
      </div>
    </main>
  );
}
