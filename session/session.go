package session

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/database"
)

type SessionData struct {
	MenuState string
}

type Session struct {
	UserID    int64
	SessionID string
	Data      SessionData
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSession(UserID int64) Session {
	sessID := fmt.Sprintf("%x", GenerateSessionID(UserID))

	return Session{
		UserID:    UserID,
		SessionID: sessID,
		Data:      NewSessionData(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewSessionData() SessionData {
	return SessionData{
		MenuState: "",
	}
}

// return []byte containing 16-base numbers
func GenerateSessionID(userID int64) []byte {
	s := fmt.Sprintf("%d %d", userID, time.Now().Unix())
	h := sha1.New()

	h.Write([]byte(s))
	sum := h.Sum(nil)

	fmt.Println(sum)

	return sum
}

func CreateSession(update tgbotapi.Update) (Session, error) {
	sess := NewSession(update.Message.From.ID)
	sess.Data.MenuState = ""

	stmt, err := database.DB.Prepare("INSERT INTO session (user_id, session_id, data) VALUES (?,?,?)")
	if err != nil {
		return sess, errors.New("[session] unable to make query")
	}

	if _, err := stmt.Exec(sess.UserID, sess.SessionID, ""); // tmp the data empty
	err != nil {
		return sess, errors.New("[session] unable to exec query")
	}

	return sess, nil
}
