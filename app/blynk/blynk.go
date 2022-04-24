package blynk

import (
	"time"

	"github.com/9d4/ampastelobot/database"
)

const (
	blynkServerAddress = "https://blynk-cloud.com/"
)

type BlynkDevice struct {
	Name  string
	Token string
}

// used to save value from database
type BlynkDeviceModel struct {
	Id         int
	UserID     int
	DeviceName string
	Token      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// if no record found nil will be returned
func GetBlynkDevicesByUserID(userID int) ([]*BlynkDevice, error) {
	sql := `SELECT * FROM blynk_tokens WHERE user_id=?`

	rows, err := database.DB.Query(sql, userID)
	if err != nil {
		return nil, nil
	}
	defer rows.Close()

	devices := []*BlynkDevice{}

	for rows.Next() {
		deviceModel := &BlynkDeviceModel{}

		_ = rows.Scan(deviceModel.Id,
			deviceModel.UserID,
			deviceModel.DeviceName,
			deviceModel.Token,
			deviceModel.CreatedAt,
			deviceModel.UpdatedAt,
		)

		if err != nil {
			continue
		}

		devices = append(devices, &BlynkDevice{
			Name:  deviceModel.DeviceName,
			Token: deviceModel.Token,
		})
	}

	return devices, nil
}

// if no record found nil will be returned
func GetBlynkDeviceByID(deviceID int) (*BlynkDevice, error) {
	sql := `SELECT * FROM blynk_tokens WHERE id=?`

	rows, err := database.DB.Query(sql, deviceID)
	if err != nil {
		return nil, nil
	}

	deviceModel := &BlynkDeviceModel{}

	if ok := rows.Next(); !ok {
		return nil, nil
	}

	err = rows.Scan(deviceModel.Id,
		deviceModel.UserID,
		deviceModel.DeviceName,
		deviceModel.Token,
		deviceModel.CreatedAt,
		deviceModel.UpdatedAt,
	)
	if err != nil {
		return nil, nil
	}

	device := &BlynkDevice{
		Name:  deviceModel.DeviceName,
		Token: deviceModel.Token,
	}

	return device, nil
}
