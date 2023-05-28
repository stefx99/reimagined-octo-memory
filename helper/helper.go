package helper

import (
	"log"
	"math/rand"
	"time"
)

func ParseDate(s string) time.Time {
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		log.Fatalf("Can't parse reservation date\n%v", err)
	}

	return date
}

func StringDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	ranStr := make([]byte, length)

	for i := 0; i < length; i++ {
		ranStr[i] = charset[rand.Intn(len(charset))]
	}
	return string(ranStr)
}
