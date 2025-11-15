"use client";

import { useEffect } from "react";
import { useParams } from "next/navigation";
import { getQuizById } from "@/lib/api/quiz";
import { useQuizStore } from "@/stores/useQuizStore";

export default function QuizDetailPage() {
  const { quiz_id } = useParams();
  const { selectedQuiz, setSelectedQuiz } = useQuizStore();

  useEffect(() => {
    async function load() {
      try {
        const data = await getQuizById(quiz_id as string);
        setSelectedQuiz(data);
      } catch (err) {
        console.error(err);
      }
    }
    load();
  }, [quiz_id, setSelectedQuiz]);

  if (!selectedQuiz) {
    return (
      <main className="min-h-screen flex items-center justify-center text-white bg-[#38325F]">
        Loading quiz...
      </main>
    );
  }

  return (
    <main className="min-h-screen flex flex-col items-center justify-center bg-[#38325F] text-white px-4">
      <div className="ng-black rounded-4xl p-16 w-full max-w-2xl">
        <div className="ng-black p-24 border-4 rounded-4xl border-[#38325F] space-y-10">
          <h1 className="text-[#38325F] font-black text-4xl text-center">
            {selectedQuiz.title}
          </h1>

          <p className="text-center opacity-80">{selectedQuiz.description}</p>

          <div className="text-center opacity-60">
            {selectedQuiz.total_questions} Questions • {selectedQuiz.duration} sec
          </div>

          <button className="mt-8 bg-white/20 px-6 py-3 rounded-xl hover:bg-white/30 transition w-full">
            Start Quiz
          </button>
        </div>
      </div>
    </main>
  );
}
