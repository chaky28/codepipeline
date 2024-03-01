package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chaky28/codepipeline/app/buildAndDeploy"
	"github.com/chaky28/codepipeline/app/file"
	"github.com/chaky28/codepipeline/app/repo"
)

var lastPush time.Time

func main() {
	fileData := file.ReadFile("/credentials/ghcredentials.txt")
	user, token := file.GetUserAndTokenFromFileData(fileData)
	fmt.Println(os.Args[1])

	repo := repo.Repo{
		User:  user,
		Token: token,
		Name:  os.Args[1],
		Ctx:   context.Background(),
	}

	repo.SetClient()

	lastPush = repo.GetLastPush()

	fmt.Println("Waiting for repo pushes...")

	for {
		newPush := repo.GetLastPush()

		if newPush.Compare(lastPush) == 1 {
			fmt.Println("New push detected.")

			buildAndDeploy.Build("build.json")
			buildAndDeploy.Deploy("deploy.json")

			lastPush = newPush
		}

		time.Sleep(time.Second * 10)
	}
}
