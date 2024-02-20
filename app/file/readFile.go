package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFile(fileLocation string) string {
	fmt.Println("Reading file from", fileLocation)

	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal("ERROR: Opening file --> " + err.Error())
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal("ERROR: Closing file --> " + err.Error())
		}
	}()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("ERROR: Reading data from file -->" + err.Error())
	}

	return string(data)
}

func GetUserAndTokenFromFileData(data string) (string, string) {
	fmt.Println("Getting user and token from github credentials file data")

	token := strings.Split(strings.Split(data, "\n")[1], "personal_token=")[1]
	user := strings.Split(strings.Split(data, "\n")[0], "user=")[1]

	return strings.Trim(user, "\r\n"), strings.Trim(token, "\r\n")
}
