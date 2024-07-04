"use client";

import { useState, useEffect, useCallback } from "react";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import BackArrow from "../../public/assets/BackArrow.svg";
import ProgressBar from "@/app/components/ProgressBar";
import CryptoJS from "crypto-js";

const secretKey = process.env.NEXT_PUBLIC_SECRET_KEY;

export default function Page() {
  const [currentQuestion, setCurrentQuestion] = useState(0);
  const [questions, setQuestions] = useState([]);
  const [score, setScore] = useState(0);
  const [isCorrect, setIsCorrect] = useState(null);
  const [selectedAnswer, setSelectedAnswer] = useState(null);
  const [isTimerRunning, setIsTimerRunning] = useState(true);
  const router = useRouter();

  useEffect(() => {
    fetchQuestions();
  }, []);

  const fetchQuestions = async () => {
    try {
      const response = await fetch("http://localhost:8000/questions");

      const data = await response.json();
      setQuestions(data);
    } catch (error) {
      console.error("Error fetching questions:", error);
    }
  };

  const handleAnswer = (selectedOption) => {
    setIsTimerRunning(false);
    const question = questions[currentQuestion];
    const isAnswerCorrect = checkAnswer(
      question.id,
      selectedOption,
      question.answer_hash
    );
    setIsCorrect(isAnswerCorrect);
    setSelectedAnswer(selectedOption);

    if (isAnswerCorrect) {
      setScore((prevScore) => prevScore + 1);
    }

    setTimeout(() => {
      handleNextQuestion();
      setSelectedAnswer(null);
      setIsCorrect(null);
    }, 1000);
  };

  const handleNextQuestion = useCallback(() => {
    if (currentQuestion < questions.length - 1) {
      setCurrentQuestion((prevQuestion) => prevQuestion + 1);
      setSelectedAnswer(null);
      setIsCorrect(null);
      setIsTimerRunning(true);
    } else {
      router.push(`/score?score=${score}&total=${questions.length}`);
    }
  }, [currentQuestion, questions.length, router]);

  const handleTimeUp = useCallback(() => {
    handleNextQuestion();
  }, [handleNextQuestion]);

  const question = questions[currentQuestion];

  const checkAnswer = (questionId, selectedOption, answerHash) => {
    const hmac = CryptoJS.HmacSHA256(
      `${questionId}:${selectedOption}`,
      secretKey
    );
    let userAnswerHash = CryptoJS.enc.Base64.stringify(hmac);
    userAnswerHash = userAnswerHash.replace(/\//g, "_").replace(/\+/g, "-");
    return userAnswerHash === answerHash;
  };

  return (
    <div className="mx-3 lg:mx-16">
      <div className="flex my-5">
        <Link href="/">
          <Image
            src={BackArrow}
            alt="Back Arrow Icon"
            width={30}
            height="auto"
            className="lg:w-[75px] h-[75px] hover:scale-125 transition duration-500"
          />
        </Link>
      </div>

      <ProgressBar onTimeUp={handleTimeUp} isRunning={isTimerRunning} />

      {questions.length > 0 && (
        <div className="my-16 lg:my-24">
          <p className="lg:text-[20px]">
            คำถาม {currentQuestion + 1} / {questions.length}
          </p>
          <p
            className={`text-left mt-6 lg:text-center text-[25px] 
              ${question.question_text.length > 80 ? "lg:mx-32" : "lg:mx-0"}`}
          >
            {question.question_text}
          </p>
        </div>
      )}

      {questions.length > 0 && (
        <div className="flex flex-col my-4 gap-y-5 lg:grid lg:grid-cols-2 lg:text-center lg:place-items-center">
          {["a", "b", "c", "d"].map((option) => (
            <div
              key={`${currentQuestion}-${option}`}
              className={`w-full rounded-[8px] border-2 border-[#181818] lg:w-4/5 
              ${
                selectedAnswer === option && isCorrect !== null
                  ? isCorrect
                    ? "bg-green-600 text-white"
                    : "bg-red-600 text-white"
                  : ""
              }`}
            >
              <button
                className=" w-full text-left text-[24px] py-2 ml-4 lg:ml-0 lg:text-center"
                onClick={() => handleAnswer(option)}
                disabled={selectedAnswer !== null}
              >
                <p>{question[`option_${option}`]}</p>
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
