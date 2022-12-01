package config

import "os"

var cfg *Config

type Config struct {
	Database DbConfig
	Bot      BotConfig
	Sheets   SheetsServiceConfig
}

type DbConfig struct {
	Path string
}
type SheetsServiceConfig struct {
	CredentialsPath string
}

type BotConfig struct {
	ID        string
	DebugMode bool
}

func Setup() {
	cfg = &Config{
		Database: DbConfig{
			Path: "/Users/mtalhakrc/go/github.com/haytek-uni-bot-yeniden/pkg/database/deneme.db",
		},
		Bot: BotConfig{
			//ID:        "5325031941:AAHSdWLZKX-2yobRnXIW9rRUH64tDObEEsc", //sıtacer bot
			ID:        "5714183726:AAGIc7aijlKxXzV8fTeimJZWUPYOMeAWTRk", //test bot
			DebugMode: false,
		},
		Sheets: SheetsServiceConfig{
			CredentialsPath: "/Users/mtalhakrc/go/github.com/haytek-uni-bot-yeniden/credentials/fluted-ranger-364116-ea4e986f9ca1.json",
		},
	}

	if os.Getenv("IS_DEVELOPMENT") == "true" {
		cfg.Bot.ID = "5325031941:AAHSdWLZKX-2yobRnXIW9rRUH64tDObEEsc"
		cfg.Bot.DebugMode = true
		cfg.Database.Path = ""
	}
}

func Get() *Config {
	return cfg
}
