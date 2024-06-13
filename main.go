package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sallescosta/bkrepo/helpers"
	"github.com/sallescosta/bkrepo/lib"
	"github.com/spf13/cobra"
)

var settings bool
var originalRepo, secondaryRepo string

func main() {
	var cmdSettings = &cobra.Command{
		Use:   "settings",
		Short: "Push, or set the original and secondary repositories",
		Long:  `Push or set the original and secondary repositories.`,
		Run: func(cmd *cobra.Command, args []string) {
			_, _, isClean := helpers.ReposVerify()
			p := tea.NewProgram(lib.Settings(originalRepo, secondaryRepo, isClean))
			if _, err := p.Run(); err != nil {
				fmt.Printf("Something went wrong: %v", err)
			}
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "gPush",
		Short: "gPush is a CLI for pushing to two repositories at once.",
		Long: "gPush is a CLI for pushing to two repositories at once. " +
			"It is useful when you have a project that is mirrored in two different repositories.",

		Run: func(cmd *cobra.Command, args []string) {
			if !settings {
				originalRepo, secondaryRepo, isClean := helpers.ReposVerify()
				if !isClean {
					fmt.Println("❌  There are uncommitted changes. Please commit them before running this program.\n")
					fmt.Println("git commit -m \"feat: my new feature\"\n\n")
					return
				}

				if originalRepo == "" {
					fmt.Println("Origin repository not found.")
					return
				}

				if secondaryRepo == "" || secondaryRepo == "undefined" {
					fmt.Println("❌  Secondary repository not found. Check the settings menu")
					p := tea.NewProgram(lib.Settings(originalRepo, secondaryRepo, isClean))
					if _, err := p.Run(); err != nil {
						fmt.Printf("something went wrong: %v", err)
					}
					return
				}

				fmt.Println("✅  Original repository: ", originalRepo)
				fmt.Println("✅  Secondary repository", secondaryRepo)
			}

			if settings {
				cmdSettings.Run(cmd, args)
				return
			}

			helpers.PushOption()
		},
	}

	rootCmd.Flags().BoolVarP(&settings, "settings", "s", false, "open menu settings")
	rootCmd.AddCommand(cmdSettings)
	rootCmd.Execute()
}
