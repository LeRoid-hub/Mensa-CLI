/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the server to get the data from",
	Long: `This command will set the server to get the data from.
	You can set the server to your own server or the default server.
	To set the server to the default server, use the argument "default"
	To set the server to your own server pass your server as the argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("No server provided")
			return
		}

		if len(args) > 1 {
			fmt.Println("Too many arguments")
			return
		}

		server := args[0]
		if server == "default" {
			fmt.Println("Setting server to default")
			viper.Set("Server", "https://mensa.barfuss.email")
			viper.WriteConfig()
			return
		}

		_, err := url.ParseRequestURI(server)
		if err != nil {
			fmt.Println("Invalid server")
			return
		}

		h, err := http.Get(server)
		if err != nil {
			fmt.Println("Server not reachable")
			return
		}
		if h.StatusCode != 200 {
			fmt.Println("Server not reachable")
			return
		}

		body, err := io.ReadAll(h.Body)
		if err != nil {
			fmt.Println("Error reading body")
			return
		}

		if string(body) != "{\"message\":\"Mensen API\"}" {
			fmt.Println("Invalid server")
			return
		}

		defer h.Body.Close()

		viper.Set("Server", server)
		fmt.Println("Server set to", server)
		viper.WriteConfig()
	},
}

func init() {
	serverCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
