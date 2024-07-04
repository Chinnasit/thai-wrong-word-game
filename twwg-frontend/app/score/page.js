"use client";

import { useSearchParams } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import KRamHappy from "../../public/assets/KRamHappy.png";
import KRamShock from "../../public/assets/KRamShock.png";
export default function Page() {
  const searchParams = useSearchParams();
  const score = searchParams.get("score");
  const total = searchParams.get("total");
  return (
    <>
      <div className="flex flex-col justify-center items-center mt-16 lg:mt-0">
        <div className="flex flex-col text-center my-2 text-[35px] lg:text-[50px]">
          <div>
            {score <= 5 ? (
              <div className="flex flex-col items-center">
                <p className="my-2 lg:text-[38px]">พ่อขุนรามรู้สึกเสียใจ</p>
                <Image width={250} height="auto" src={KRamShock} alt="shock" />
              </div>
            ) : (
              <div className="flex flex-col items-center">
                <p className="my-2 lg:text-[38px]">พ่อขุนรามกดไลก์ให้คุณ</p>
                <Image width={250} height="auto" src={KRamHappy} alt="happy" />
              </div>
            )}
          </div>
          <p className="lg:text-[35px]">
            คุณตอบถูก <span className="text-[60px]">{score}</span> ข้อ
          </p>
          <p className="lg:text-[35px]">
            จากทั้งหมด <span>{total}</span> ข้อ
          </p>
        </div>

        <div className="flex justify-center pt-12 lg:pt-8">
          <Link
            href="/questions"
            className="bg-[#181818] px-16 py-2 text-[50px] text-white rounded-[8px] hover:scale-125 transition duration-500"
          >
            เล่นอีกครั้ง
          </Link>
        </div>
      </div>

      <div className="text-center">
        <Link
          href="/"
          className="fixed bottom-0 left-0 right-0 w-full text-[#181818] text-[30px]"
        >
          กลับหน้าแรก
        </Link>
      </div>
    </>
  );
}
