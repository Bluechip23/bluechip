package main

import (
    "fmt"
    "os"

    "github.com/BlueChip23/bluechip/app"
    "github.com/cosmos/cosmos-sdk/server"
    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
    rootCmd, _ := NewRootCmd()

    // Add a version command
    rootCmd.AddCommand(&cobra.Command{
        Use:   "version",
        Short: "Print the version of the application",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("Version: %s\n", Version)
            fmt.Printf("Git Commit: %s\n", GitCommit)
            fmt.Printf("Build Date: %s\n", BuildDate)
        },
    })

    if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
        switch e := err.(type) {
        case server.ErrorCode:
            os.Exit(e.Code)
        default:
            os.Exit(1)
        }
    }
}