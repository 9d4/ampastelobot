package session

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/9d4/ampastelobot/action"
	"github.com/9d4/ampastelobot/database"
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

type ErrNoSessionInDB struct {
	s string
}

func (e *ErrNoSessionInDB) Error() string {
	return e.s
}

func NewErrNoSessInDB() error {
	return &ErrNoSessionInDB{s: "no record based on user_id"}
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

	return sum
}

// create session for user and push to db
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

// get user session from db
func GetSession(update tgbotapi.Update) (Session, error) {
	var sess Session
	userID := update.Message.From.ID

	// find userID in DB
	rows, err := database.DB.Query("SELECT * FROM session WHERE user_id = ?", userID)
	if err != nil {
		return sess, errors.New("[session] unable to make query")
	}
	defer rows.Close()

	// take the first record
	columns, _ := rows.Columns()
	length := len(columns)

	// save unused data from scan
	trash := make([]interface{}, length)

	if found := rows.Next(); !found {
		return sess, &ErrNoSessionInDB{s: "no record based on user_id"}
	}

	// values are [id, user_id, session_id, data, created_at, updated_at]
	if err := rows.Scan(&trash[0], &sess.UserID, &sess.SessionID, &trash[3], &sess.CreatedAt, &sess.UpdatedAt); err != nil {
		return sess, err
	}

	return sess, nil
}

// Session Middleware checks wheter user has session in DB or not,
// if it does, then continue, else create one.
func Middleware(botUpdate *action.BotUpdate) bool {
	_, update := botUpdate.Bot, *botUpdate.Update

	// check wheter has session or not
	if _, err := GetSession(update); err != nil {
		switch err {
		case err.(*ErrNoSessionInDB):
			// if error is ErrNoSessionInDB
			// then create session
			CreateSession(update)
			botUpdate.Data["first_time"] = true
		default:
			log.Println(err)
			return false
		}
	}

	return true
}
