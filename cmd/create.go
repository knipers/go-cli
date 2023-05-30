/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/knipers/go-cli/internal/database"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
/*
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
	RunE: runCreate(GetAuthorDB(GetDB())),
}*/

func newCreateCmd(authorDb database.Author) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new Author",
		Long:  `Create a new Author`,
		RunE:  runCreate(authorDb),
	}
}

func runCreate(authorDB database.Author) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		_, err := authorDB.Create(name)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetAuthorDB(GetDb()))
	authorCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the author")
	createCmd.MarkFlagRequired("name")
}
