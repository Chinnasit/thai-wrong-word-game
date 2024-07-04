package db

import (
	"Chinnasit/pkg/common/entities"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Trace SQL command
type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n----------------\n", sql)
}

// Initial the database
func Init() *gorm.DB {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	projectRoot := filepath.Dir(wd)
	err = godotenv.Load(filepath.Join(projectRoot, "/pkg/common/envs/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: &SqlLogger{}})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(entities.Question{})

	var count int64
	db.Model(&entities.Question{}).Count(&count)
	if count == 0 {
		CreateQuestions(db)
	}

	return db
}

func CreateQuestions(db *gorm.DB) {
	questions := []entities.Question{
		{QuestionText: "การประชุมรูปแบบหนึ่งซึ่งมีวัตถุประสงค์เพื่อแลกเปลี่ยนความรู้ความคิดเห็น",
			OptionA:       "สัมมนา",
			OptionB:       "สัมนา",
			OptionC:       "สัมณา",
			OptionD:       "สมัมณา",
			CorrectOption: "a"},
		{QuestionText: "ชื่อเดือนที่ 7",
			OptionA:       "กรกดาคม",
			OptionB:       "กรกฎคม",
			OptionC:       "กรกฎาคม",
			OptionD:       "กรกฏาคม",
			CorrectOption: "c"},
		{QuestionText: "การศึกษาระหว่างประถมศึกษากับอุดมศึกษา",
			OptionA:       "มัธยมศึกษา",
			OptionB:       "มัทธยมศึกษา",
			OptionC:       "มัธยมมศึกษา",
			OptionD:       "มัทธยมมศึกษา",
			CorrectOption: "a"},
		{QuestionText: "อำนาจซึ่งแฝงอยู่ในบุคคลหรือรัฐ ซึ่งสามารถบันดาลให้เป็นไปตามความประสงค์",
			OptionA:       "อิทธิพล",
			OptionB:       "อิทธพล",
			OptionC:       "อิทธิพน",
			OptionD:       "อิทธพน",
			CorrectOption: "a"},
		{QuestionText: "ผู้จัดเลือกเฟ้น รวบรวม ปรับปรุง และรับผิดชอบเรื่องลงพิมพ์",
			OptionA:       "บรรณาธิการ",
			OptionB:       "บรรณธิการ",
			OptionC:       "บรรณาที่การ",
			OptionD:       "บรรณทิการ",
			CorrectOption: "a"},
		{QuestionText: "ชื่อเดือนที่ 11",
			OptionA:       "พฤศจิกายน",
			OptionB:       "พฤศิกายน",
			OptionC:       "พฤษจิกายน",
			OptionD:       "พฤศภาคม",
			CorrectOption: "a"},
		{QuestionText: "กฎเกณฑ์ที่ผู้มีอำนาจตราขึ้นเพื่อใช้บังคับบุคคลให้ปฏิบัติตามเป็นการทั่วไป",
			OptionA:       "กฏหมาย",
			OptionB:       "กฎหมาย",
			OptionC:       "กฏหมายย",
			OptionD:       "กฎหมายย",
			CorrectOption: "b"},
		{QuestionText: "ชื่อไม้ล้มลุกชนิดหนึ่ง กลิ่นฉุน ใช้ปรุงอาหาร",
			OptionA:       "กะเพรา",
			OptionB:       "กะเพร่า",
			OptionC:       "กระเพรา",
			OptionD:       "กระเพร่า",
			CorrectOption: "a"},
		{QuestionText: "บุคคลซึ่งได้รับเลือกหรือได้รับแต่งตั้งเพื่อกระทำกิจการบางอย่างเป็นคณะ",
			OptionA:       "กรรมการ",
			OptionB:       "กรรมการณ",
			OptionC:       "กรรมกาน",
			OptionD:       "กรรมการร",
			CorrectOption: "a"},
		{QuestionText: "งานที่จำต้องทำ",
			OptionA:       "ภาระกิจ",
			OptionB:       "ภารกิจ",
			OptionC:       "ภาระกิจก",
			OptionD:       "ภารกิจก",
			CorrectOption: "b"},
		{QuestionText: "แนวความคิดเห็น",
			OptionA:       "ทัศนคติ",
			OptionB:       "ทัศนคิติ",
			OptionC:       "ทัศนคติิ",
			OptionD:       "ทัศนคติิ",
			CorrectOption: "a"},
		{QuestionText: "บประกาศที่ให้ไว้เพื่อเป็นเกียรติแก่บุคคลซึ่งมีความสามารถ หรือบำเพ็ญประโยชน์แก่สังคม",
			OptionA:       "เกียรติบัตร",
			OptionB:       "เกียรติบัต",
			OptionC:       "เกียรติบัตร์",
			OptionD:       "เกียรบัตร",
			CorrectOption: "a"},
		{QuestionText: "สิ่งที่ทิ้งไปเพราะไม่ใช้แล้ว",
			OptionA:       "ขยะ",
			OptionB:       "ขยะะ",
			OptionC:       "ขยั",
			OptionD:       "ขยัะ",
			CorrectOption: "a"},
		{QuestionText: "คำที่ถูกต้องคือคำใด?",
			OptionA:       "ดอกไม้",
			OptionB:       "ดอกไมย",
			OptionC:       "ดอกไม้์",
			OptionD:       "ดอกไม์",
			CorrectOption: "a"},
		{QuestionText: "ประมวลความประพฤติที่ผู้ประกอบอาชีพการงานแต่ละอย่างกำหนดขึ้น เพื่อรักษาและส่งเสริมเกียรติคุณชื่อเสียงและฐานะของสมาชิก",
			OptionA:       "จรรยาบรรณ",
			OptionB:       "จรรยาบรรณณ",
			OptionC:       "จรรยาบรร",
			OptionD:       "จรรยาบรรณร",
			CorrectOption: "a"},
		{QuestionText: "ตัวหนังสือ",
			OptionA:       "อักษร",
			OptionB:       "อักษน",
			OptionC:       "อักษร",
			OptionD:       "อักษร์",
			CorrectOption: "a"},
		{QuestionText: "สถานที่เก็บรวบรวมและแสดงสิ่งต่าง ๆ ที่มีความสำคัญด้านวัฒนธรรมหรือด้านวิทยาศาสตร์ โดยมีความมุ่งหมายเพื่อให้เป็นประโยชน์ต่อการศึกษา และก่อให้เกิดความเพลิดเพลินใจ",
			OptionA:       "พิพิธภัณฑ์",
			OptionB:       "พิพิธภัณ",
			OptionC:       "พิพิธพันธ์",
			OptionD:       "พิพิธพันธ",
			CorrectOption: "a"},
		{QuestionText: "ห้องหรืออาคารที่มีระบบจัดเก็บรวบรวมรักษาหนังสือประเภทต่าง ๆ ซึ่งอาจรวมทั้งต้นฉบับ ลายมือเขียน ไมโครฟิล์ม เป็นต้น เพื่อใช้เป็นที่ค้นคว้าหาความรู้",
			OptionA:       "ห้องสมุด",
			OptionB:       "ห้องสมัย",
			OptionC:       "ห้องสมัฏ",
			OptionD:       "ห้องสมัศ",
			CorrectOption: "a"},
		{QuestionText: "ชื่อไม้เถาชนิดหนึ่ง ทอดเลื้อยตามพื้นดินหรือบนผิวนํ้า ดอกสีขาวหรือม่วงอ่อน ลำต้นกลวง ยอดกินได้",
			OptionA:       "ผักบุ้ง",
			OptionB:       "ผักบุ่ง",
			OptionC:       "ผักบุ๊ง",
			OptionD:       "ผักบุง",
			CorrectOption: "a"},
		{QuestionText: "ชื่อดาวเคราะห์ดวงที่ 5 และเป็นดวงที่ใหญ่ที่สุดในระบบสุริยะ",
			OptionA:       "พฤหัสบดี",
			OptionB:       "พฤหัส",
			OptionC:       "พฤหัศบดี",
			OptionD:       "พฤหัสษบดี",
			CorrectOption: "a"},
		{QuestionText: "ส่วนย่อย",
			OptionA:       "แผนก",
			OptionB:       "แผนค",
			OptionC:       "แผนค์",
			OptionD:       "แผน",
			CorrectOption: "a"},
		{QuestionText: "เป็นพืชล้มลุกชนิดหนึ่ง จัดอยู่ในตระกูลหญ้า",
			OptionA:       "หญ้าคา",
			OptionB:       "หญ้าฆา",
			OptionC:       "หญ้าข่า",
			OptionD:       "หญ้าขา",
			CorrectOption: "a"},
		{QuestionText: "อาหาร (ใช้แก่พระภิกษุสามเณร) ",
			OptionA:       "บิณฑบาต",
			OptionB:       "บิณฑบาท",
			OptionC:       "บิณบาต",
			OptionD:       "บิณบาท",
			CorrectOption: "a"},
		{QuestionText: "ที่ดิน ทุ่ง นา ไร่",
			OptionA:       "เกษตร",
			OptionB:       "กะเสด",
			OptionC:       "กะเศด",
			OptionD:       "เกษด",
			CorrectOption: "a"},
		{QuestionText: "องค์กรปกครองส่วนท้องถิ่นรูปแบบหนึ่ง ประกอบด้วยนายกเทศมนตรีและสภาเทศบาล",
			OptionA:       "เทศบาล",
			OptionB:       "เทศบาลณ",
			OptionC:       "เทศบาลล",
			OptionD:       "เทศบาน",
			CorrectOption: "a"},
		{QuestionText: "ติดต่อสื่อสารเพื่อส่งเสริมความเข้าใจอันถูกต้องต่อกัน",
			OptionA:       "ประชาสัมพันธ์",
			OptionB:       "ประชาสัมพัน",
			OptionC:       "ประชาสัมพันท",
			OptionD:       "ประชาสัมภัณ",
			CorrectOption: "a"},
		{QuestionText: "การแสดงผลงาน สินค้า ผลิตภัณฑ์ หรือกิจกรรม ให้คนทั่วไปชม",
			OptionA:       "นิทรรศการ",
			OptionB:       "นิทรรศกร",
			OptionC:       "นิทรรศการณ",
			OptionD:       "นิทรรษการ",
			CorrectOption: "a"},
		{QuestionText: "ข้อตกลงกัน คำมั่น",
			OptionA:       "สัญญา",
			OptionB:       "สัญยา",
			OptionC:       "สันญา",
			OptionD:       "สันยา",
			CorrectOption: "a"},
		{QuestionText: "คำที่ถูกต้องคือคำใด?",
			OptionA:       "คำศัพท์",
			OptionB:       "คำศัพ",
			OptionC:       "คำศัพท",
			OptionD:       "คำศัฟ",
			CorrectOption: "a"},
		{QuestionText: "เอกสารแสดงคุณวุฒิ ตามปรกติตํ่ากว่าระดับอุดมศึกษา?",
			OptionA:       "ประกาศนียบัตร",
			OptionB:       "ประกาศนียบัต",
			OptionC:       "ประกาศณียบัตร",
			OptionD:       "ประกาศณียบัต",
			CorrectOption: "a"},
	}

	result := db.Create(&questions)
	if result.Error != nil {
		panic(result.Error)
	}
}
