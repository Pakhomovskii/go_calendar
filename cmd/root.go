package cmd

import (
	"fmt"
	_ "github.com/fatih/color"
	"github.com/spf13/cobra"
	_ "math"
	"os"
	"os/exec"
	_ "os/exec"
)

var clearScreen = func() {
	cmd := exec.Command(`printf '\33c\e[3J'`)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Exit the application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exiting...")
		os.Exit(0)
	},
}

var rootCmd = &cobra.Command{
	Use:   "go_cal",
	Short: "Prints current month and infinite watch. Use no-color command if you CLI does not support colors",
	Run: func(cmd *cobra.Command, args []string) {

		generateOutput(false)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.AddCommand(noColorCmd)
	rootCmd.AddCommand(exitCmd)
}
