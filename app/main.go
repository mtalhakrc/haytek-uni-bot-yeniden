package main

import (
	"github.com/haytek-uni-bot-yeniden/app/handlers"
	"github.com/haytek-uni-bot-yeniden/app/service"
	"github.com/haytek-uni-bot-yeniden/pkg/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	ceteleapp := app.New()

	s := service.NewSheetsService(ceteleapp.Sheets, "15sToZcfyEp95WINbv1nuD_sTtTZxn1RmhgkBrlLIw9g")

	commands := ceteleapp.Commands
	scheduled := ceteleapp.Scheduled

	userhandler := handlers.NewUserHandler(ceteleapp.DB, s)
	scheduledHandler := handlers.NewScheduled(ceteleapp.DB, s)
	cetelehandler := handlers.NewCeteleHandler(ceteleapp.DB)

	commands.RegisterCommand("kaydol", userhandler.Kaydol)
	commands.RegisterCommand("kayitsil", userhandler.KayitSil)

	commands.RegisterCommand("start", cetelehandler.Start)
	commands.RegisterCommand("admins", cetelehandler.Admins)

	commands.RegisterCommand("gunlukozet", cetelehandler.GetSpecific)
	commands.RegisterCommand("haftalikozet", cetelehandler.GetHaftalikOzet)

	scheduled.RegisterScheduled("23:00:00", scheduledHandler.CeteleHatirlatmaMesaji)
	scheduled.RegisterScheduled("00:00:00", scheduledHandler.GunlukErkenKontrolMesaji)
	scheduled.RegisterScheduled("01:00:00", scheduledHandler.GunlukRaporMesaji)

	ceteleapp.Start()
}