# TelegramBot

TelegramBot is a simple Telegram bot designed to automate interactions with users through the Telegram API. This bot allows you to execute various commands and receive information directly in your Telegram chat.

## Bot Link

You can find our bot at the following link: t.me/boss1223_bot

## Installation

+ ### Clone the Repository
```
git clone https://github.com/yourusername/telegrambot.git
cd telegrambot
```
+ ### Install Dependencies
Ensure you have Go installed on your machine. You can download it from the official Go website. Then, install the required Go libraries:
```
go get github.com/spf13/cobra
go get gopkg.in/telebot.v3
```
+ ### Set Up Environment Variables
Copy your token and past after next command:
```
read -s TELE_TOKEN
```
+ ### Build and Run the Bot
```
go build -ldflags "-X="github.com/RuslanOSheiko/kbot/cmd.appVersion=v1.0.2
```
Then, run the bot:
```
./TelegramBot start
```
