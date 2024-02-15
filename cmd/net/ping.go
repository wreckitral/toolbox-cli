/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This command ping a remote url and return the response",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}

		// fmt.Println("ping called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}

	NetCmd.AddCommand(pingCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
