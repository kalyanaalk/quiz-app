import { create } from "zustand";

export interface Quiz {
  id: string;
  title: string;
  description: string;
  total_questions: number;
  duration: number;
}

export interface QuizQuestion {
  id: string;
  content: string;
  type: boolean;       
  quiz_id: string;
}

export interface Answer {
  id: string;
  content: string;
  is_correct: boolean;
  question_id: string;
}

interface QuizStore {
  quizzes: Quiz[];
  selectedQuiz: Quiz | null;

  questions: QuizQuestion[];
  answers: Record<string, Answer[]>;

  setQuizzes: (data: Quiz[]) => void;
  setSelectedQuiz: (quiz: Quiz | null) => void;

  setQuestions: (data: QuizQuestion[]) => void;
  setAnswersForQuestion: (question_id: string, answers: Answer[]) => void;
}

export const useQuizStore = create<QuizStore>((set) => ({
  quizzes: [],
  selectedQuiz: null,

  questions: [],
  answers: {},

  setQuizzes: (data) => set({ quizzes: data }),
  setSelectedQuiz: (quiz) => set({ selectedQuiz: quiz }),

  setQuestions: (data) => set({ questions: data }),

  setAnswersForQuestion: (question_id, answers) =>
    set((state) => ({
      answers: {
        ...state.answers,
        [question_id]: answers,
      },
    })),
}));
