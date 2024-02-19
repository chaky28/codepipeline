package repo

import (
	"context"
	"log"
	"time"

	"github.com/google/go-github/v59/github"
)

type Repo struct {
	User   string
	Token  string
	Ctx    context.Context
	Name   string
	Client *github.Client
}

func (repo Repo) GetLastPush() time.Time {
	loc, _ := time.LoadLocation("Local")
	return repo.GetRepo().PushedAt.In(loc)
}

func (repo *Repo) SetClient() {
	repo.Client = github.NewClient(nil).WithAuthToken(repo.Token)
}

func (repo Repo) GetRepo() *github.Repository {
	result, _, err := repo.Client.Repositories.Get(repo.Ctx, repo.User, repo.Name)
	if err != nil {
		log.Fatal("ERROR: Getting repo -->" + err.Error())
	}
	return result
}
