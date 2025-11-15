import { authFetch } from "./fetcher";

const BASE = "http://173.255.114.209:8080";

export async function getQuestionsByQuiz(quiz_id: string) {
  const data = await authFetch(`${BASE}/questions/quiz/${quiz_id}`);
  return data.questions;
}

export async function getAnswersByQuestion(question_id: string) {
  const data = await authFetch(`${BASE}/answers/question/${question_id}`);
  return data.answers;
}

export interface Answer {
  id: string;
  content: string;
  is_correct: boolean;
  question_id: string;
}

export interface Question {
  id: string;
  content: string;
  type: boolean;
  quiz_id: string;
  answers: Answer[];
}
