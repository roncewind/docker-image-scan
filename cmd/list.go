/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		checkDockerExists()
		osCmd := exec.Command("docker", "images")
		var stdout, stderr bytes.Buffer
		osCmd.Stdout = &stdout
		osCmd.Stderr = &stderr
		err := osCmd.Run()
		if err != nil {
			log.Fatalf("osCmd.Run() failed with %s\n", err)
		}
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkDockerExists() {
    path, err := exec.LookPath("docker")
    if err != nil {
        fmt.Printf("didn't find 'docker' executable\n")
    } else {
        fmt.Printf("'docker' executable is in '%s'\n", path)
    }
}