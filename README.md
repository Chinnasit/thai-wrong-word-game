# ภาพรวม
  **Repo นี้คืออิหยัง ??**    
  web application นี้เป็นแอปพลิเคชั่นสำหรับการตอบคำถาม (Quiz App) เกี่ยวกับคำที่มักเขียนผิดในภาษาไทย
  โดยในแต่ละครั้งที่เล่นระบบจะสุ่ม 10 คำถามให้เล่น ผู้ใช้ตอบคำถามจนกระทั่งครบทุกข้อแล้วแสดงคะแนนที่ได้  
  
  **พัฒนาด้วย**   
  Frontend : Next.js 14, Tailwind CSS  
  Backend : Go-Fiber  
  Database : PostgreSQL  
# วิธีการรัน Web Application
**Running Database Server:**  
  * config ไฟล์ .env.example ที่ twwg-backend\pkg\common\envs  
  * รัน docker-compose
     ```bash
     cd twwg-backend
     docker-compose up -d
     ```  
   * เปิด pgadmin4 ผ่านเบราว์เซอร์ที่ localhost:5050 หรือ port ที่ config ตาม docker-compose จากนั้น login และเชื่อมต่อกับ postgresql
  
**Running Backend:**  
   การรันและติดตั้ง Backend Dependencies  
   ```bash
    cd twwg-backend
    go get .
    go run main.go
   ```
**Running Frontend:**  
   การรันและติดตั้ง Backend Dependencies  
   ```bash
    cd twwg-frontend
    npm install
    npm run dev
   ```
