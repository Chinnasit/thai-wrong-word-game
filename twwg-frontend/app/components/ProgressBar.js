"use client";

import Image from "next/image";
import Clock from "../../public/assets/Clock.svg";
import { useState, useEffect, useRef } from "react";

const TOTAL_TIME = 45000; // 45 seconds in milliseconds

export default function ProgressBar({onTimeUp, isRunning}) {
    const [startTime, setStartTime] = useState(null);
    const [now, setNow] = useState(null);
    const intervalRef = useRef(null);
    const timeoutRef = useRef(null);

    useEffect(() => {
      if (isRunning) {
        setStartTime(Date.now());
        setNow(Date.now());
    
        intervalRef.current = setInterval(() => {
          setNow(Date.now());
        }, 100);
    
        timeoutRef.current = setTimeout(() => {
          onTimeUp();
        }, TOTAL_TIME);
      } else {
        clearInterval(intervalRef.current);
        clearTimeout(timeoutRef.current);
      }
  
      return () => {
        clearInterval(intervalRef.current);
        clearTimeout(timeoutRef.current);
      };
    }, [onTimeUp, isRunning]);

    const elapsedTime = now && startTime ? now - startTime : 0;
    const progress = Math.min((elapsedTime / TOTAL_TIME) * 100, 100);
    const timeLeft = Math.max(TOTAL_TIME - elapsedTime, 0);

  return (
    <div className="w-full">
      <div className="h-8 relative w-full rounded-full overflow-hidden bg-white">
        <div
          className="h-full bg-[#181818] transition-all duration-100 ease-linear absolute left-0 top-0"
          style={{ width: `${progress}%` }}
        ></div>
        <div className="absolute left-2 top-1/2 transform -translate-y-1/2 flex items-center space-x-2">
          <Image src={Clock} alt="Clock Icon" width={25} height={25}/>
          <span className="text-white">
          {Math.ceil(timeLeft / 1000)} วินาที
          
          </span>
        </div>
      </div>
    </div>
  );
}
