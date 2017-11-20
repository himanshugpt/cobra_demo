package main

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var (
	Env string
	Name string
	Test bool
	Create bool
	IsDry bool
)

var RootCmd = &cobra.Command{
	Use:   "Demo",
	Short: "CLI tool",
	Long:  `CLI Tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("in Run Root ")
		if len(Env) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

  fmt.Println("Env ", Env)
  fmt.Println("Name ", Name)
	fmt.Println("Test ", Test)
  fmt.Println("Create ", Create)
  fmt.Println("IsDry ", IsDry)
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Env, "env", "e","dev", "Environment to run")
	RootCmd.PersistentFlags().StringVarP(&Name, "name", "n", "none", "Name")
	RootCmd.PersistentFlags().BoolVarP(&IsDry, "dry", "d", false, "is dry run")
	RootCmd.PersistentFlags().BoolVarP(&Test, "test", "t", false, "run tests")
	RootCmd.PersistentFlags().BoolVarP(&Create, "create", "c", false, "create suite")
	fmt.Println("in init ")
}
