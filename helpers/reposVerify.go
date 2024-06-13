package helpers

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
)

func ReposVerify() (originalRepo, secondaryRepo string, isClean bool) {

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("\n❌  Could not open the repository. Please ensure this program is run from the root directory of your project: %v", err)
	}

	wt, err := repo.Worktree()
	if err != nil {
		log.Fatalf("\n❌  Could not get the worktree: %v", err)
	}

	status, err := wt.Status()
	if err != nil {
		log.Fatalf("\n❌  Could not get the status: %v", err)
	}

	isClean = status.IsClean()

	original, err := repo.Remote("origin")
	if err != nil {
		fmt.Printf("\n❌  Could not fetch the 'origin' remote repository. Please ensure it is properly configured. Error: %v",
			err)
		return "", "undefined", isClean
	}

	originalRepo = original.Config().URLs[0]

	secondary, err := repo.Remote("secondary")
	if err != nil {
		fmt.Printf("❌  Could not fetch the 'secondary' remote repository. Please ensure it is properly configured. "+
			"Error: %v\n",
			err)
		return originalRepo, "", isClean
	}

	secondaryRepo = secondary.Config().URLs[0]
	return originalRepo, secondaryRepo, isClean
}
