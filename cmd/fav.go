/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/LeRoid-hub/Mensa-CLI/internal"
	"github.com/fatih/color"
	"github.com/rodaine/table"
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
		favorites := viper.Get("favorites").([]interface{})

		s := make([]string, len(favorites))
		for i, v := range favorites {
			s[i] = v.(string)
		}
		for i := 0; i < len(s); i++ {
			data, err := internal.GetMenu(s[i])
			if err != nil {
				fmt.Println("Error fetching data")
			}
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("Offering", "Dish", "Price")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			color.Green("Mensa: %s", data.Name)

			for _, day := range data.Days {
				for _, menu := range day.Menu {
					for _, meal := range menu.Meal {
						tbl.AddRow(menu.Name, meal.Name, meal.Price)
					}
				}
			}

			tbl.Print()

			fmt.Println()
		}

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
