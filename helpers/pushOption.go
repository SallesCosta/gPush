package helpers

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func execCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func PushTo(repo string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("⬆️ Pushing %s\n", repo)
	out, err := execCommand("git", "push", repo)
	if err != nil {
		fmt.Printf("❌  %v\n", out)
		return
	}
	log.Printf("Push to %s done: %s\n", repo, out)
}

func PushOption() {
	var wg sync.WaitGroup
	wg.Add(2)

	go PushTo("origin", &wg)
	go PushTo("secondary", &wg)

	wg.Wait()
	return
}
