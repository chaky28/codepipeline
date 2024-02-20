package buildAndDeploy

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/chaky28/codepipeline/app/file"
)

type Config struct {
	Commands []struct {
		Name string   `json:"name"`
		Args []string `json:"args"`
		Go   bool     `json:"go"`
	} `json:"commands"`
}

func Build(buildConf string) {
	configStr := file.ReadFile(buildConf)

	var jsonConfig Config
	err := json.Unmarshal([]byte(configStr), &jsonConfig)
	if err != nil {
		log.Fatal("ERROR: Parsing build config file -->" + err.Error())
	}

	for _, command := range jsonConfig.Commands {
		trimmedCommand := strings.TrimSpace(command.Name)

		if trimmedCommand == "" {
			log.Fatal("ERROR: Empty command detected")
		}

		bash := exec.Command(trimmedCommand, command.Args...)

		if command.Go {
			err = bash.Run()
		}

		if err != nil {
			log.Fatal("ERROR: Running command " + trimmedCommand + " with args " + strings.Join(command.Args, " ") + "--> " + err.Error())
		}

		output, _ := bash.Output()

		fmt.Println(string(output))
	}
}

func Deploy(deployConf string) {
	configStr := file.ReadFile(deployConf)

	var jsonConfig Config
	err := json.Unmarshal([]byte(configStr), &jsonConfig)
	if err != nil {
		log.Fatal("ERROR: Parsing deploy config file -->" + err.Error())
	}

	for _, command := range jsonConfig.Commands {
		trimmedCommand := strings.TrimSpace(command.Name)

		if trimmedCommand == "" {
			log.Fatal("ERROR: Empty command detected.")
		}

		bash := exec.Command(trimmedCommand, command.Args...)

		if command.Go {
			err = bash.Run()
		}

		if err != nil {
			log.Fatal("ERROR: Running command " + trimmedCommand + " with args " + strings.Join(command.Args, " ") + " --> " + err.Error())
		}

		output, _ := bash.Output()

		fmt.Println(string(output))
	}
}
