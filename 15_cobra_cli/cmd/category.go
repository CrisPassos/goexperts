/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// name, _ := cmd.Flags().GetString("name")
		if category != "" {
			println("Category name:", category)
		}
		exists, _ := cmd.Flags().GetBool("exists")
		if exists {
			println("Verificando se a categoria existe...")
		}
		id, _ := cmd.Flags().GetInt16("id")
		if id != 0 {
			println("Category ID:", id)
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		println("Executando antes de qualquer comando filho")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		println("Executando depois de qualquer comando filho")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		//retorna um erro caso eu queira
		println("Executando a categoria")
		return nil
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Name of the category")
	categoryCmd.Flags().BoolP("exists", "e", false, "Verifica se a categoria existe")
	categoryCmd.Flags().Int16P("id", "i", 0, "Id of the category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
