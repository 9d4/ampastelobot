package common

import "os"

func Wd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}
