package helpers

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lib "github.com/sallescosta/gPush/lib/textinput"
	"log"
)

func DeleteRepositorio(originOrSecondary string) (out string, err error) {
	out, err = execCommand("git", "remote", "remove", originOrSecondary)
	return string(out), err
}

func SetarRepositorios(originOrSecondary, repoName string) (out string, err error) {
	out, err = execCommand("git", "remote", "add", originOrSecondary, repoName)
	return string(out), err
}

func SetRepos(originOrSecondary string) {
	p := tea.NewProgram(lib.InputModel(originOrSecondary))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	if lib.InputValue == "" {
		return
	}

	out, err := DeleteRepositorio(originOrSecondary)
	if err != nil {
		fmt.Println(out)
		return
	}

	out, err = SetarRepositorios(originOrSecondary, lib.InputValue)
	if err != nil {
		fmt.Println(out)
		return
	}
	fmt.Printf("%s is the new %s repository\n", lib.InputValue, originOrSecondary)
}
