package main

import (
        "log"
        "os"
        "strconv"
        "strings"
        "os/exec"
        "gopkg.in/telegram-bot-api.v4"

)
// System dependencies if used: net-tools (netstat), mtr, nmap, dig

// ENV variables required: TGBOT_TOKEN and TGBOT_CHATID

// Global variables 
var clean_domain string
var tgbot_token string
var tgbot_chatid string
var chat_id string

func checkErr(err error) {
    if err != nil {
        log.Println("Error while running cmd",err )
    }
}

func removeCharacters(input string, characters string) string {
    filter := func(r rune) rune {
            if strings.IndexRune(characters, r) < 0 {
                    return r
            }
            return -1
    }

    return strings.Map(filter, input)

}


func exec_shell(command string) string {


log.Println("/bin/bash -c",command)
out, err := exec.Command("/bin/bash","-c",command).Output()
checkErr(err)
        return string(out)
}

func main() {

        f, err := os.OpenFile("bot.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                log.Fatal(err)
        }
        defer f.Close()
        log.SetOutput(f)
        if os.Getenv("TGBOT_TOKEN") != "" && os.Getenv("TGBOT_CHATID") != ""  {
        	tgbot_token = os.Getenv("TGBOT_TOKEN")
            tgbot_chatid =  os.Getenv("TGBOT_CHATID")
        } else {
        	log.Println("Telegram env vars are not set... skipping notification")
        }
        bot, err := tgbotapi.NewBotAPI(tgbot_token)
        if err != nil {
                log.Panic(err)
        }
        //bot.Debug = true
        log.Printf("Authorized on account %s", bot.Self.UserName)

        u := tgbotapi.NewUpdate(0)
        u.Timeout = 60
        updates, err := bot.GetUpdatesChan(u)

        for update := range updates {
            if update.Message == nil {
                        continue
             }
           // chat_id := update.Message.Chat.ID
            chat_id:= strconv.Itoa(int(update.Message.Chat.ID))

            if strings.Contains(chat_id , tgbot_chatid) {

                log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
                if strings.Contains(update.Message.Text,"/sh") {
                    
                  cmd_raw:=update.Message.Text
                  words := strings.Fields(cmd_raw)
                 if len(words) == 1 {
                      log.Println("Bot missing parameter")
                    }
                if len(words) >= 2 {

                      cmd_clean:=  strings.Replace(cmd_raw, "/sh", "", 1)
                      log.Println("Rimmed cmd ",cmd_clean)
                      output:=exec_shell(cmd_clean)
                      msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
                      bot.Send(msg)
                                   }

                                                             }

                if strings.Contains(update.Message.Text,"/sysinfo") {
 
                      output:=exec_shell("df -h && free -m && netstat -tunlp")
                      msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
                      bot.Send(msg)
                                   

                                                             }
                if strings.Contains(update.Message.Text,"/dig") {

                    cmd_raw:=update.Message.Text

                    words:= strings.Fields(cmd_raw)

                    if len(words) == 2 {


                        cmd_clean:= words[1]
                                   
                      output:=exec_shell("dig +short "+cmd_clean)
                      msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
                      bot.Send(msg)
                                   
                                         }
                                                             }
 
                if strings.Contains(update.Message.Text,"/nmap") {
                    cmd_raw:=update.Message.Text

                    words := strings.Fields(cmd_raw)

                    if len(words) == 2 {


                        cmd_clean:= words[1]
                                   
                      output:=exec_shell("nmap -Pn -A "+cmd_clean)
                      msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
                      bot.Send(msg)
                                   
                                         }
                                                             }

                if strings.Contains(update.Message.Text,"/whois") {
                    cmd_raw:=update.Message.Text

                    words := strings.Fields(cmd_raw)

                    if len(words) == 2 {


                        cmd_clean:= words[1]
                                   
                      output:=exec_shell("whois "+cmd_clean)
                      msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
                      bot.Send(msg)
                                   
                                         }
		}
 
                if strings.Contains(update.Message.Text,"/curl") {
                    cmd_raw:=update.Message.Text

                    words := strings.Fields(cmd_raw)

                    if len(words) == 2 {


                        cmd_clean:=words[1]
                                   
                      output:=exec_shell("curl "+cmd_clean)
                      msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
                      bot.Send(msg)
                                   
                                         }
                                                             }


              if update.Message.Text == "/help" {
              msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Usage: \n /sh <CMD> - Executes a command on the shell of the remote system \n/sysinfo - Display generic troubleshoot system information (disk usage, network, memory)\n/dig - Resolve the given domain, supports RR.. example /dig A google.com or /dig MX google.com\n /mtr <hostname | ip> - prints reachability report to a specific hosts using the tool MTR\n /nmap - runs the command NMAP \n /curl executes a curl request and shows status code")

              bot.Send(msg)


                                            }

                                        }

   } // for range updates



} //main
