import { create } from "zustand";

export interface Quiz {
  id: string;
  title: string;
  description: string;
  total_questions: number;
  duration: number;
}

interface QuizStore {
  quizzes: Quiz[];
  selectedQuiz: Quiz | null;

  setQuizzes: (data: Quiz[]) => void;
  setSelectedQuiz: (quiz: Quiz | null) => void;
}

export const useQuizStore = create<QuizStore>((set) => ({
  quizzes: [],
  selectedQuiz: null,

  setQuizzes: (data) => set({ quizzes: data }),
  setSelectedQuiz: (quiz) => set({ selectedQuiz: quiz }),
}));
