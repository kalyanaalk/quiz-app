import { Suspense } from 'react';
import dynamic from 'next/dynamic'; 

const LoginFormContent = dynamic(() => import('@/components/login'), {
    loading: () => (
    <div className="bg-[url('/background/navbar.jpg')] p-24 border-4 rounded-4xl border-[#38325F] space-y-16">
      <h1 className="text-[#38325F] font-black text-4xl text-center">
        WELCOME BACK!
      </h1>
      <div className="flex flex-col gap-4 w-full max-w-xs mx-auto">
        <p className="text-[#38325F] text-center text-lg">Loading login form...</p>
      </div>
    </div>
  ),
});

export default function LoginPage() {
  return (
    <main className="h-screen flex flex-col items-center justify-center bg-[#38325F] text-white px-4">
      <div className="bg-[url('/background/navbar.jpg')] rounded-4xl p-16">
        <Suspense fallback={<div>Preparing login form...</div>}>
          <LoginFormContent />
        </Suspense>
      </div>
    </main>
  );
}