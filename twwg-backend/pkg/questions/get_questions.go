package questions

import (
	"Chinnasit/pkg/common/entities"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type ClientQuestion struct {
	ID           uint   `json:"id"`
	QuestionText string `json:"question_text"`
	OptionA      string `json:"option_a"`
	OptionB      string `json:"option_b"`
	OptionC      string `json:"option_c"`
	OptionD      string `json:"option_d"`
	AnswerHash   string `json:"answer_hash"`
}

func (h handler) GetQuestions(c *fiber.Ctx) error {
	// get random questions
	var questions []entities.Question
	result := h.DB.Order("RANDOM()").Limit(10).Find(&questions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch questions",
		})
	}

	// mapping request before response
	clientQuestions := make([]ClientQuestion, len(questions))
	for i, q := range questions {
		answerHash := generateAnswerHash(q.ID, q.CorrectOption)
		clientQuestions[i] = ClientQuestion{
			ID:           q.ID,
			QuestionText: q.QuestionText,
			OptionA:      q.OptionA,
			OptionB:      q.OptionB,
			OptionC:      q.OptionC,
			OptionD:      q.OptionD,
			AnswerHash:   answerHash,
		}
	}

	return c.Status(fiber.StatusOK).JSON(clientQuestions)
}

func generateAnswerHash(questionID uint, correctOption string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	projectRoot := filepath.Dir(wd)
	err = godotenv.Load(filepath.Join(projectRoot, "/pkg/common/envs/.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY_ANSWER_HASH")
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(fmt.Sprintf("%d:%s", questionID, correctOption)))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
