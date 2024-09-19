/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// favCmd represents the fav command
var favCmd = &cobra.Command{
	Use:   "fav",
	Short: "Retrieves the favorite mensa menus",
	Long: `Retrieves the favorite mensa menus for the current day.
	You can add mensas to your favorites using the search command.`,
	Run: func(cmd *cobra.Command, args []string) {
		favorites := viper.GetStringSlice("favorites")
		if len(favorites) == 0 {
			fmt.Println("You have no favorites")
			return
		}

		fmt.Println("Your favorites: " + strings.Join(favorites, ", "))

	},
}

func init() {
	rootCmd.AddCommand(favCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// favCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// favCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
