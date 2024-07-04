"use client";

import Link from "next/link";
import { Canvas } from "@react-three/fiber";
import { Environment, OrbitControls } from "@react-three/drei";
import { Model } from "./components/Model";

export default function Home() {
  return (
    <main className="flex flex-col justify-center items-center">
      <div className="mx-4 flex flex-col items-center">
        <div className="h-[350px] mt-8">
          <Canvas>
            <Environment preset="studio" />
            <OrbitControls />
            <Model />
          </Canvas>
        </div>

        <p className="text-[#181818] text-[19px] text-center mt-4 lg:text-[35px]">
          พ่อขุนรามถูกใจสิ่งนี้! 9,999+ คำไทยเขียนผิด ท้าให้ทดสอบความเซียน!
          กดเล่นด่วน!
        </p>
      </div>
      <div className="flex justify-center pt-28 lg:pt-10">
        <Link
          href="questions"
          className="bg-[#181818] px-16 py-2 text-[50px] text-white rounded-[8px] hover:scale-125 transition duration-500"
        >
          เริ่มเกม
        </Link>
      </div>
    </main>
  );
}
