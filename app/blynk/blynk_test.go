package blynk

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/9d4/ampastelobot/database"
)

func TestGetBlynkDeviceByUserID(t *testing.T) {
	t.Parallel()

	database.Init()

	// this data based on the database used for test
	testDevices := []BlynkDevice{
		// {Name: "testdvc", Token: "token:wkwkwk"},
		// {Name: "testdvc2", Token: "token:wkwkwkwk"},
	}

	userID := 123

	devices, err := GetBlynkDevicesByUserID(userID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(database.DB)

	for _, dev := range devices {
		found := false

		for _, dt := range testDevices {
			if reflect.DeepEqual(dev, dt) {
				found = true
				break
			}
		}

		if found {
			t.Fatal("Device Not found in DB")
		}
	}
}
