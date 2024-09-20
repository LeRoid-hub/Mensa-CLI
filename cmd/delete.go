/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a favorite mensa",
	Long:  `Returns a list of your favorite mensas and prompts you to select one to delete.`,
	Run: func(cmd *cobra.Command, args []string) {
		favorites := viper.Get("favorites").([]interface{})

		s := make([]string, len(favorites))
		for i, v := range favorites {
			s[i] = v.(string)
		}

		prompt := promptui.Select{
			Label: "Select Mensa to delete",
			Items: s,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		for i := 0; i < len(s); i++ {
			if s[i] == result {
				s = append(s[:i], s[i+1:]...)
				break
			}
		}

		viper.Set("favorites", s)
		viper.WriteConfig()

	},
}

func init() {
	favCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
