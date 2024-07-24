/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tb "gopkg.in/telebot.v3"
)

var (
	//TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// TelegramBotCmd represents the TelegramBot command
var TelegramBotCmd = &cobra.Command{
	Use:     "TelegramBot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("TelegramBot %s started", appVersion)
		TelegramBot, err := tb.NewBot(tb.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatal(err)
			return
		}

		TelegramBot.Handle(tb.OnText, func(m tb.Context) error {

			log.Print(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hi":
				err = m.Send(fmt.Sprintf("Hi, it's TelegramBot %s", appVersion))

			}

			return err
		})

		TelegramBot.Start()

	},
}

func init() {
	rootCmd.AddCommand(TelegramBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// TelegramBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// TelegramBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
