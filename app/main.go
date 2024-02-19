package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chaky28/codepipeline/app/file"
	"github.com/chaky28/codepipeline/app/repo"
)

var lastPush time.Time

func main() {
	fileData := file.ReadFile("/ghcredentials.txt")
	user, token := file.GetUserAndTokenFromFileData(fileData)

	repo := repo.Repo{
		User:  user,
		Token: token,
		Name:  os.Args[0],
		Ctx:   context.Background(),
	}
	repo.SetClient()

	lastPush = repo.GetLastPush()

	for {
		newPush := repo.GetLastPush()

		if newPush.Compare(lastPush) == 1 {
			fmt.Println("New push detected.")

			lastPush = newPush
		}

		time.Sleep(time.Second * 10)
	}
}
