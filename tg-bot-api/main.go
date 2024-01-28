package main

import (
	"log"
	"time"

	//"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


var buttonInlineClass = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("танк", "code_tank"),
		tgbotapi.NewInlineKeyboardButtonData("боец", "code_boes"),
		tgbotapi.NewInlineKeyboardButtonData("убийца", "code_ydisa"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("стрелок", "code_strelok"),
		tgbotapi.NewInlineKeyboardButtonData("маг", "code_mash"),
		tgbotapi.NewInlineKeyboardButtonData("поддержка", "code_pod"),
	),
)

var buttonFighterClass = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("чонг", "code_shong"),
		tgbotapi.NewInlineKeyboardButtonData("теризла", "code_teriz"),
		tgbotapi.NewInlineKeyboardButtonData("зилонг", "code_zilong"),
		tgbotapi.NewInlineKeyboardButtonData("альфа", "code_alpha"),
	),
)

var buttonstelokClass = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("москва", "code_moskva"),
		tgbotapi.NewInlineKeyboardButtonData("клинт", "code_klint"),
	),
)
var buttonstankClass = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Фредрин", "code_fred"),
	),
)
var buttonsmurdererClass = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("алукард", "code_alu"),
		tgbotapi.NewInlineKeyboardButtonData("ланселот", "code_lans"),
		tgbotapi.NewInlineKeyboardButtonData("наташа", "code_natasha"),
	),
)
var buttonmagicianClass = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("валентина", "code_valentina"),
		tgbotapi.NewInlineKeyboardButtonData("гвиневра", "code_gvin"),
		
	),
)


func main() {
	bot, err := tgbotapi.NewBotAPI("bot-api")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {

			
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "привет!")
			bot.Send(msg)
			time.Sleep(3 * time.Second)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Это телеграм бот в котором вы можете посмотреть гайд на любово персонажа из инры Mobail Legend")
			bot.Send(msg)
			time.Sleep(3 * time.Second)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Удачного использывания")
			bot.Send(msg)
			time.Sleep(2 * time.Second)
			msgs := tgbotapi.NewMessage(update.Message.Chat.ID, "выберите класс")
			msgs.ReplyMarkup = buttonInlineClass

			// Send the message
			bot.Send(msgs)
			
						
			
		} else if update.CallbackQuery != nil && update.CallbackQuery.Data != "" {

			
			switch update.CallbackQuery.Data {

			case "code_shong":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
					// Удаляем сообщение о начале загрузки
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("output(compress-video-online.com)shong.mp4"))

				video.Thumb = tgbotapi.FileID("output(compress-video-online.com)shong.mp4")

				bot.Send(video)
			case "code_teriz":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
					
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2023-12-16_22-59-28.mp4"))

				video.Thumb = tgbotapi.FileID("video_2023-12-16_22-59-28.mp4")

				bot.Send(video)
			case "code_zilong":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
				
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2023-12-24_21-33-44silong.mp4"))

				video.Thumb = tgbotapi.FileID("video_2023-12-24_21-33-44silong.mp4")

				bot.Send(video)
			case "code_alpha":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
					
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2024-01-23_11-49-04.mp4"))

				video.Thumb = tgbotapi.FileID("video_2024-01-23_11-49-04.mp4")

				bot.Send(video)




			case "code_boes":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "выберите бойца")
				msg.ReplyMarkup = buttonFighterClass

				bot.Send(msg)

			}
				

			
			



			switch update.CallbackQuery.Data {

			case "code_moskva":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
					
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2023-12-24_21-35-07moskv.mp4"))

				video.Thumb = tgbotapi.FileID("video_2023-12-24_21-35-07moskv.mp4")

				bot.Send(video)
			case "code_klint":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
				
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2024-01-23_11-49-04.mp4"))

				video.Thumb = tgbotapi.FileID("video_2024-01-23_11-49-04.mp4")

				bot.Send(video)

			case "code_strelok":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "выберите стрелка")
				msg.ReplyMarkup = buttonstelokClass

			
				bot.Send(msg)

			}

			switch update.CallbackQuery.Data {

			case "code_fred":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
			
				
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("output(compress-video-online.com) (15).mp4"))

				video.Thumb = tgbotapi.FileID("output(compress-video-online.com) (15).mp4")

				bot.Send(video)

			case "code_tank":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "выберите танка")
				msg.ReplyMarkup = buttonstankClass

			
				bot.Send(msg)

			}
			switch update.CallbackQuery.Data {

			case "code_alu":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
				
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
			
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2023-12-24_21-34-00alukard.mp4"))

				video.Thumb = tgbotapi.FileID("video_2023-12-24_21-34-00alukard.mp4")

				bot.Send(video)
			case "code_lans":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
				
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
				
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2023-12-24_21-34-18lans.mp4"))

				video.Thumb = tgbotapi.FileID("video_2023-12-24_21-34-18lans.mp4")

				bot.Send(video)
			case "code_natasha":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
					
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
			
	
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2024-01-05_19-03-10 natasha.mp4"))

				video.Thumb = tgbotapi.FileID("video_2024-01-05_19-03-10 natasha.mp4")

				bot.Send(video)

			case "code_ydisa":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "выберите убийцу")
				msg.ReplyMarkup = buttonsmurdererClass

				bot.Send(msg)

			}
			switch update.CallbackQuery.Data {

			case "code_code_gvin":
			
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
					
				
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
	
			
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2024-01-23_11-48-55.mp4"))

				video.Thumb = tgbotapi.FileID("video_2024-01-23_11-48-55.mp4")

				bot.Send(video)
			case "code_valentina":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Пожайлуста подождите идет загрузка видео")
				sentMsg, err := bot.Send(msg)
					if err != nil {
						log.Println("Error sending message:", err)
						return
					}
					
					time.Sleep(5 * time.Second)
			
				deleteMsg := tgbotapi.NewDeleteMessage(sentMsg.Chat.ID, sentMsg.MessageID)
				bot.Send(deleteMsg)
		
		
				video := tgbotapi.NewVideo(update.CallbackQuery.Message.Chat.ID, tgbotapi.FilePath("video_2023-12-24_21-34-47valentina.mp4"))

				video.Thumb = tgbotapi.FileID("video_2023-12-24_21-34-47valentina.mp4")

				bot.Send(video)

			case "code_mash":
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "выберите мага")
				msg.ReplyMarkup = buttonmagicianClass

		
				bot.Send(msg)

			}

			

		}
	}

}