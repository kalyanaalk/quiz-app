import { authFetch } from "./fetcher";
import { useAuthStore } from "@/stores/useAuthStore";
import { Score, ScoreDetailPayload, ScoreDetailResponse } from "@/stores/useScoreStore";

const BASE_URL = "http://173.255.114.209:8080";

export async function startScore(quiz_id: string): Promise<Score> {
  const user_id = useAuthStore.getState().user?.id;
  if (!user_id) throw new Error("User not logged in");

  return await authFetch(`${BASE_URL}/scores/start`, {
    method: "POST",
    body: JSON.stringify({ user_id, quiz_id }),
  });
}

export async function submitScore(score_id: string, answers: ScoreDetailPayload[]) {
  return await authFetch(`${BASE_URL}/scores/${score_id}/submit`, {
    method: "PUT",
    body: JSON.stringify(answers),
  });
}

export async function getScoreDetail(score_id: string): Promise<ScoreDetailResponse> {
  return await authFetch(`${BASE_URL}/scores/${score_id}/detail`);
}

export async function getQuizResult(quiz_id: string) {
  const user_id = useAuthStore.getState().user?.id;
  if (!user_id) throw new Error("User not logged in");

  return await authFetch(`${BASE_URL}/scores/quiz/${quiz_id}/user/${user_id}`);
}

export async function getScoreDetails(score_id: string) {
  return await authFetch(`${BASE_URL}/score-details/${score_id}`);
}
