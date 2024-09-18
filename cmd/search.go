/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/LeRoid-hub/Mensa-CLI/internal"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

		prompt = promptui.Select{
			Label: "Do you want to see the menu or save the mensa to your favorites?",
			Items: []string{"menu", "favorites"},
		}

		_, result, err = prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		if result == "menu" {
			menu, err := internal.GetMenu(result)
			if err != nil {
				fmt.Println("Error fetching data")
			}

			fmt.Println(menu)
		}

		if result == "favorites" {
			// viper
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
