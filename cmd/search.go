/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/LeRoid-hub/Mensa-CLI/internal"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Gets a Mensa to display or add it to your favorites",
	Long: `This command will prompt you to select a state, city and mensa.
	You can then choose to display the menu or save the mensa to your favorites.`,

	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Select Day",
			Items: []string{"baden-wuerttemberg", "bayern", "berlin", "brandenburg", "bremen",
				"hamburg", "hessen", "mecklenburg-vorpommern", "niedersachsen", "nordrhein-westfalen", "rheinland-pfalz", "saarland", "sachsen", "sachsen-anhalt", "schleswig-holstein", "thueringen"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		cities, err := internal.GetState(result)
		if err != nil {
			fmt.Println("Error fetching data")
		}

		city := strings.Split(cities, ",")

		prompt = promptui.Select{
			Label: "Select City",
			Items: city,
		}

		_, result, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		selectedCity := result

		mensen, err := internal.GetMensen(result)
		if err != nil {
			fmt.Println("Error fetching data")
		}

		mensa := strings.Split(mensen, ",")

		prompt = promptui.Select{
			Label: "Select Mensa",
			Items: mensa,
		}

		_, result, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		selectedMensa := result

		prompt = promptui.Select{
			Label: "Do you want to see the menu or save the mensa to your favorites?",
			Items: []string{"menu", "favorites"},
		}

		_, result, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == "menu" {
			data, err := internal.GetMenu(selectedCity, selectedMensa)
			if err != nil {
				fmt.Println("Error fetching data")
			}
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("Offering", "Dish", "Price")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, day := range data.Days {
				for _, menu := range day.Menu {
					for _, meal := range menu.Meal {
						tbl.AddRow(menu.Name, meal.Name, meal.Price)
					}
				}
			}

			tbl.Print()
		}

		if result == "favorites" {
			favorites := viper.Get("favorites").([]string)
			favorites = append(favorites, selectedCity+"/"+selectedMensa)
			fmt.Println("Added " + selectedCity + "/" + selectedMensa + " to your favorites")
			viper.Set("favorites", favorites)
			fmt.Println(viper.Get("favorites"))
			viper.SafeWriteConfig()
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
