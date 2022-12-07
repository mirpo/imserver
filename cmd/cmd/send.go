/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"math/big"
	"time"
)

const source1Jwt = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzb3VyY2VJZCI6MX0.5VtXO9J1YF2sv8SwTfvsVseqHMjEwhFBHJLpSuj-i34"

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send sample log to imserver",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		val, _ := rand.Int(rand.Reader, big.NewInt(10000000))
		client := resty.New().
			SetBaseURL("http://0.0.0.0:1323/v1").
			SetHeader("Content-Type", "application/json").
			SetHeader("Accept", "application/json")
		resp, _ := client.R().
			SetBody(fmt.Sprintf(`{"metrics":"{\"ts\": \"%s\", \"value\": %d}"}`, time.Now(), val)).
			SetAuthToken(source1Jwt).
			Post("/logs")
		fmt.Println("Logs stored at server:", string(resp.Body()))
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
