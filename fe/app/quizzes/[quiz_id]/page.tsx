"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { getQuizById } from "@/lib/api/quiz";
import { startScore, getQuizResult } from "@/lib/api/score";
import { useAuthStore } from "@/stores/useAuthStore";
import { useScoreStore, Score } from "@/stores/useScoreStore";

interface QuizDetail {
  id: string;
  title: string;
  description: string;
  total_questions: number;
  duration: number;
}

export default function QuizDetailPage() {
  const { quiz_id } = useParams();
  const router = useRouter();

  const user = useAuthStore((s) => s.user);
  const setScore = useScoreStore((s) => s.setScore);

  const [quiz, setQuiz] = useState<QuizDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [starting, setStarting] = useState(false);

  const [scoreHistory, setScoreHistory] = useState<any[]>([]);
  const [loadingScores, setLoadingScores] = useState(true);

  useEffect(() => {
    async function load() {
      try {
        const q = await getQuizById(quiz_id as string);
        setQuiz(q);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, [quiz_id]);

  useEffect(() => {
    async function loadScores() {
      if (!user) return;
      try {
        const res = await getQuizResult(quiz_id as string);
        setScoreHistory(res);
      } finally {
        setLoadingScores(false);
      }
    }
    loadScores();
  }, [quiz_id, user]);
  
  const handleStart = async () => {
    if (!user) return alert("Login required");

    try {
      setStarting(true);

      const newScore: Score = await startScore(quiz_id as string);
      setScore(newScore);

      router.push(`/quizzes/${quiz_id}/${newScore.id}`);
    } finally {
      setStarting(false);
    }
  };

  if (loading)
    return <div className="text-center text-white mt-20">Loading...</div>;

  if (!quiz)
    return <div className="text-center text-red-300 mt-20">Quiz not found</div>;

  return (
    <main className="min-h-screen flex flex-col items-center bg-[#38325F] text-white px-4 py-10">
      <div className="bg-white/10 rounded-4xl p-12 w-full max-w-3xl space-y-12">

        <div className="space-y-5">
          <h1 className="text-[#ffffff] font-black text-4xl text-center">
            {quiz.title}
          </h1>

          <p className="text-center opacity-80">{quiz.description}</p>

          <div className="text-center opacity-60">
            {quiz.total_questions} Questions • {quiz.duration} sec
          </div>

          <button
            onClick={handleStart}
            disabled={starting}
            className={`mt-6 px-6 py-3 rounded-xl w-full transition 
            ${starting ? "bg-white/10 cursor-not-allowed" : "bg-white/20 hover:bg-white/30"}`}
          >
            {starting ? "Starting..." : "Start Quiz"}
          </button>
        </div>

        <div className="mt-10">
          <h2 className="text-2xl font-bold mb-4 text-center">Your Scores</h2>

          {loadingScores ? (
            <p className="text-center opacity-50">Loading history...</p>
          ) : scoreHistory.length === 0 ? (
            <p className="text-center opacity-50">No attempts yet.</p>
          ) : (
            <table className="w-full border-separate border-spacing-y-2">
              <thead>
                <tr className="text-left bg-white/10 rounded-xl">
                  <th className="p-3">Attempt</th>
                  <th className="p-3">Score</th>
                  <th className="p-3">Finished At</th>
                </tr>
              </thead>

              <tbody>
                {scoreHistory.map((row: any, index: number) => (
                  <tr
                    key={row.id}
                    onClick={() =>
                      router.push(`/quizzes/${quiz_id}/${row.id}/result`)
                    }
                    className="cursor-pointer bg-white/5 hover:bg-white/20 transition rounded-xl"
                  >
                    <td className="p-3">Attempt #{index + 1}</td>
                    <td className="p-3">{row.score ?? "-"}</td>
                    <td className="p-3">{row.finished_at || "Not finished"}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>
      </div>
    </main>
  );
}
