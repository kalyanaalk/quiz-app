import { authFetch } from "./fetcher";
import { Quiz } from "@/stores/useQuizStore";

const BASE_URL = "http://localhost:8080/quizzes";

export async function getAllQuizzes(): Promise<Quiz[]> {
  return await authFetch(`${BASE_URL}/`);
}

export async function getQuizById(id: string): Promise<Quiz> {
  return await authFetch(`${BASE_URL}/${id}`);
}
