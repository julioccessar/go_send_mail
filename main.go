package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	ViperConf()
	Send("hello there! with viper Split to!")
}

func Send(body string) {
	from := viper.GetString("mail.from")
	pass := viper.GetString("mail.pwd")
	toArr := strings.Split(viper.GetString("mail.to"), ",")

	msg := "From: " + from + "\n" +
		"To: " + strings.Join(toArr, ",") + "\n" +
		"Subject: Hello there!\n\n" +
		body

	err := smtp.SendMail(viper.GetString("smtp.server")+":"+viper.GetString("smtp.port"),
		smtp.PlainAuth("", from, pass, viper.GetString("smtp.server")),
		from, toArr, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Message sent!")
}

func ViperConf() {
	// Leer configuración
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	log.Printf("Init conf ok")
}
