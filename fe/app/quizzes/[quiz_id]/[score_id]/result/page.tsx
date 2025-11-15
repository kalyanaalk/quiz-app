"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { getScoreDetail, getScoreDetails } from "@/lib/api/score";

export default function QuizResultPage() {
  const { quiz_id, score_id } = useParams();
  const router = useRouter();

  const [loading, setLoading] = useState(true);
  const [detail, setDetail] = useState<any>(null);
  const [detailsList, setDetailsList] = useState<any>([]);

  useEffect(() => {
    async function load() {
      try {
        const scoreDetail = await getScoreDetail(score_id as string);
        const scoreDetailsList = await getScoreDetails(score_id as string);

        setDetail(scoreDetail);
        setDetailsList(scoreDetailsList);
      } finally {
        setLoading(false);
      }
    }
    load();
  }, [score_id]);

  if (loading)
    return (
      <div className="text-center text-white mt-20">
        Loading result...
      </div>
    );

  if (!detail)
    return (
      <div className="text-center text-red-300 mt-20">
        Result not found
      </div>
    );

  return (
    <main className="min-h-screen flex flex-col items-center bg-[#38325F] text-white px-6 py-10">
      <div className="bg-white/10 p-8 rounded-3xl w-full max-w-2xl text-center space-y-6">

        <h1 className="text-3xl font-black">Quiz Result</h1>

        {/* <div className="text-xl">
          Score:{" "}
          <span className="font-bold text-green-300">
            {correct}
          </span>{" "}
        </div> */}

        <div className="opacity-70 text-sm">
          Score ID: {score_id}
        </div>

        {/* Result Details List */}
        <div className="text-left mt-6 space-y-4">
          <h2 className="text-lg font-bold">Your Answers</h2>

          {detail?.details?.length ? (
  detail.details.map((item: any, index: number) => (
    <div
      key={index}
      className={`p-4 rounded-xl ${
        item.is_correct ? "bg-green-700/40" : "bg-red-700/40"
      }`}
    >
      <p className="font-semibold">
        Q{index + 1}: {item.question_id}
      </p>
      <p className="text-sm opacity-80">
        Answers: {item.answer_ids.join(", ")}
      </p>
      <p className="mt-1 text-sm">
        {item.is_correct ? "✅ Correct" : "❌ Incorrect"}
      </p>
    </div>
  ))
) : (
  <p className="text-sm opacity-50">No score detail available.</p>
)}

        </div>

        <div className="mt-10 text-left">
          <h2 className="text-lg font-bold">Raw Details (Debug)</h2>

          <pre className="bg-black/20 p-4 rounded-xl text-xs overflow-x-auto">
            {JSON.stringify(detailsList, null, 2)}
          </pre>
        </div>

        <button
          onClick={() => router.push(`/quizzes/${quiz_id}`)}
          className="mt-6 px-6 py-3 rounded-xl bg-white/20 hover:bg-white/30 transition w-full"
        >
          Reattempt Quiz
        </button>
      </div>
    </main>
  );
}
