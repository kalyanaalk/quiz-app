import { Quiz } from "@/stores/useQuizStore";

const BASE_URL = "http://localhost:8080/quizzes";

export async function getAllQuizzes(): Promise<Quiz[]> {
  const res = await fetch(`${BASE_URL}/`, { method: "GET" });
  const json = await res.json();
  if (!res.ok) throw new Error(json.error || "Failed to fetch quizzes");
  return json;
}

export async function getQuizById(id: string): Promise<Quiz> {
  const res = await fetch(`${BASE_URL}/${id}`, { method: "GET" });
  const json = await res.json();
  if (!res.ok) throw new Error(json.error || "Failed to fetch quiz detail");
  return json;
}
