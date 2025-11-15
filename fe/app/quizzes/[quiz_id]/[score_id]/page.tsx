"use client";

import { useEffect } from "react";
import { useParams, useRouter } from "next/navigation";

import { useQuestionStore } from "@/stores/useQuestionStore";
import { useScoreStore } from "@/stores/useScoreStore";
import { submitScore } from "@/lib/api/score";

import { useAuthStore } from "@/stores/useAuthStore";

export default function QuizWorkPage() {
  const { quiz_id, score_id } = useParams(); 
  const router = useRouter();

  const { user } = useAuthStore();

  const {
    questions,
    userAnswers,
    loadQuestions,
    selectAnswer,
  } = useQuestionStore();

  const { setScore, currentScore, buildSubmitPayload } = useScoreStore();

  useEffect(() => {
    async function init() {
      if (!user) {
        alert("Anda harus login.");
        return router.push("/login");
      }

      await loadQuestions(quiz_id as string);

      setScore({
        id: score_id as string,
        quiz_id: quiz_id as string,
        user_id: user.id,
      });
    }

    init();
  }, [quiz_id, score_id, user, loadQuestions, setScore, router]);

  async function handleSubmit() {
    if (!currentScore)
      return alert("Score belum dimulai!");

    const payload = buildSubmitPayload();

    await submitScore(currentScore.id, payload);

    router.push(`/quizzes/${quiz_id}/${currentScore.id}/result`);
  }

  if (questions.length === 0)
    return (
      <main className="min-h-screen flex items-center justify-center bg-[#38325F] text-white">
        <p>Loading questions...</p>
      </main>
    );

  return (
    <main className="min-h-screen bg-[#38325F] text-white p-8">
      <div className="max-w-3xl mx-auto space-y-10">

        <h1 className="text-3xl font-bold text-center mb-10">Quiz</h1>

        {questions.map((q, index) => {
          const selected = userAnswers[q.id] || [];
          const isMultiple = q.type === true;

          return (
            <div key={q.id} className="bg-white text-black p-6 rounded-xl space-y-4">
              <h2 className="text-xl font-semibold">
                {index + 1}. {q.content}
              </h2>

              <div className="space-y-2">
                {q.answers.map((a) => {
                  const isSelected = selected.includes(a.id);

                  return (
                    <label
                      key={a.id}
                      className="flex items-center gap-3 p-2 border rounded-lg cursor-pointer hover:bg-gray-100"
                    >
                      <input
                        type={isMultiple ? "checkbox" : "radio"}
                        name={q.id}
                        checked={isSelected}
                        onChange={() => selectAnswer(q.id, a.id, isMultiple)}
                      />
                      <span>{a.content}</span>
                    </label>
                  );
                })}
              </div>
            </div>
          );
        })}

        <button
          onClick={handleSubmit}
          className="w-full py-4 bg-green-500 hover:bg-green-600 text-white font-semibold text-lg rounded-xl"
        >
          Submit Quiz
        </button>

      </div>
    </main>
  );
}
