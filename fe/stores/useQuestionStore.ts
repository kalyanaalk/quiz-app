import { create } from "zustand";
import { getQuestionsByQuiz, getAnswersByQuestion, Question } from "@/lib/api/question";

interface UserAnswers {
  [question_id: string]: string[];
}

interface QuestionStore {
  questions: Question[];
  userAnswers: UserAnswers;

  loadQuestions: (quiz_id: string) => Promise<void>;

  selectAnswer: (question_id: string, answer_id: string, isMultiple: boolean) => void;

  reset: () => void;
}

export const useQuestionStore = create<QuestionStore>((set, get) => ({
  questions: [],
  userAnswers: {},

  loadQuestions: async (quiz_id: string) => {
  const qs = await getQuestionsByQuiz(quiz_id);

  const questionsWithAnswers = await Promise.all(
    qs.map(async (q: Question) => {
      const answers = await getAnswersByQuestion(q.id);
      return { ...q, answers };
    })
  );

  set({ questions: questionsWithAnswers });
},

  selectAnswer: (question_id, answer_id, isMultiple) => {
    const { userAnswers } = get();

    if (isMultiple) {
      const current = userAnswers[question_id] || [];
      const exists = current.includes(answer_id);

      const updated = exists
        ? current.filter((id) => id !== answer_id)
        : [...current, answer_id];

      set({
        userAnswers: { ...userAnswers, [question_id]: updated },
      });
    } else {
      set({
        userAnswers: { ...userAnswers, [question_id]: [answer_id] },
      });
    }
  },

  reset: () => set({ questions: [], userAnswers: {} }),
}));
