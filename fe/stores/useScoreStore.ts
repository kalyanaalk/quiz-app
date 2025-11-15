import { create } from "zustand";
import { useQuestionStore } from "./useQuestionStore";

export interface Score {
  id: string;
  user_id: string;
  quiz_id: string;
  start_at?: string;
  finished_at?: string;
}

export interface ScoreDetailPayload {
  question_id: string;
  answer_ids: string[];
}

export interface ScoreDetailResponse {
  score: Score;
  details: Array<{
    question_id: string;
    answer_ids: string[];
    is_correct: boolean;
  }>;
}


interface ScoreStore {
  currentScore: Score | null;

  setScore: (score: Score | null) => void;

  buildSubmitPayload: () => ScoreDetailPayload[];
}

export const useScoreStore = create<ScoreStore>((set) => ({
  currentScore: null,

  setScore: (score) => set({ currentScore: score }),

  buildSubmitPayload: () => {
    const qStore = useQuestionStore.getState();
    const answers = qStore.userAnswers;

    return Object.entries(answers).map(([question_id, ids]) => ({
      question_id,
      answer_ids: ids ?? [],
    }));
  },
}));
