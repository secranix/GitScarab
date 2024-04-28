package main

import (
	"github.com/secranix/GitScarab/internal/env"
)

func main() {
	env.GetEnv(env.GitlabPAT, env.WithOptionalFlag())
	env.GetEnv(env.GithubUser)
}
