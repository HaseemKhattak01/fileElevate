package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "my_cobra_project",
    Short: "A brief description of your application",
    Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hello from Cobra CLI!")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}