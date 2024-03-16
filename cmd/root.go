package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var (
	copyToClipboard bool

	rootCmd = &cobra.Command{
		Use:   "doggo",
		Short: "doggo - go get a dog",
		Run: func(cmd *cobra.Command, args []string) {
			dog := `
    / \__
   (    @\___
   /         O
 /   (_____/
/_____/   U
`
			blue := color.New(color.FgHiBlue).SprintFunc()
			fmt.Println(blue(dog))
			if copyToClipboard {
				copyToClipboardFunc(dog)
			}
		},
	}
)

func setupCobraUsageTemplate() {
	cobra.AddTemplateFunc("StyleHeading", color.New(color.FgBlue).SprintFunc())
	usageTemplate := rootCmd.UsageTemplate()
	usageTemplate = strings.NewReplacer(
		`Usage:`, `{{StyleHeading "Usage:"}}`,
		`Flags:`, `{{StyleHeading "Flags:"}}`,
	).Replace(usageTemplate)
	rootCmd.SetUsageTemplate(usageTemplate)
}

func init() {
	setupCobraUsageTemplate()
	rootCmd.Flags().BoolVarP(&copyToClipboard, "copy", "c", false, "copy ASCII dog to clipboard")
	rootCmd.DisableFlagsInUseLine = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func copyToClipboardFunc(data string) {
	err := clipboard.Init()
	if err != nil {
		fmt.Println(color.HiRedString(fmt.Sprintf("Error: %s", err)))
		return
	}
	clipboard.Write(clipboard.FmtText, []byte(data))
}
