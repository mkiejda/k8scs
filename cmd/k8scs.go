package main

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"os/exec"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func getContexts() (contexts []string, err error) {
	contextsRaw, err := exec.Command("kubectl", "config", "get-contexts", "-o", "name").Output()
	if err != nil {
		return nil, err
	}
	contexts = strings.Fields(string(contextsRaw))
	return contexts, nil
}

func checkKubectl() (err error) {
	_, err = exec.LookPath("kubectl")
	if err != nil {
		err = errors.New("kubectl missing")
		return err
	}
	return nil
}

func main() {
	err := checkKubectl()
	check(err)

	contexts, err := getContexts()
	check(err)

	prompt := promptui.Select{
		Label: "Select Context",
		Items: contexts,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	cmd := exec.Command("kubectl", "config", "use-context", result)
	err = cmd.Run()
	check(err)

}
