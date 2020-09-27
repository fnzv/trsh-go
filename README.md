## trsh-go

Telegram Remote-Shell is a Golang script that let you comunicate with your Linux server using a Telegram Bot  <br>
This repo is a re-make of my original Python script created a few years ago [here](https://github.com/fnzv/trsh)

-----------------
![](img/bot.png?raw=true)
 <br>

 -------------------------------
## Requirements
- Linux OS (tested on Ubuntu 20.04 and Centos 7)
- Telegram Bot created using @BotFather
- Packages required: jq
- (Optional packages used):  nmap, dig, mtr, net-tools (netstat)

## Quick setup

Before setup: <br>

* Chat with BotFather in order to create a Bot ( https://telegram.me/botfather ), just launch the command /newbot and get your Telegram Token. <br>
 Open the bot chat and send some messages to activate the bot. <br>

* Launch this oneliner+script on your Linux system (where the bot will be running as a service): <br>

```
git clone https://github.com/fnzv/trsh-go /home/trsh/ && cd /home/trsh/ && bash install.sh
```

##### WARNING: this command will install the required/missing packages ( dnsutils, nmap, mtr, net-tools )

##### NOTES:

- You will be asked to insert your Telegram Bot Token acquired on the first step and then follow the quick installation wizard. <br>

- The script will gather your chat-id based on the last unread message you send to the bot on the first step. <br>

After you finished the installation the Golang script binary will run as a system service with systemd (trsh.service).

## Usage

- /dig - Resolve the given domain, supports RR.. example /dig A google.com or /dig MX google.com
- /mtr - Execute a mtr with a following report
- /nmap - Execute a nmap -Pn -A
- /curl - Execute a curl request
- /whois - Whois lookup
- /sysinfo - Display generic system information (disk usage, network & memory)
- /sh - Execute a command with Bash.. example /sh cat namefile , /sh ps auxf | grep ssh

## Tests

The following scripts are tested on Ubuntu 20.04 LTS, Centos 7 and marked as working.


## Run using docker

Inside this repo you will find a prepared Docker image  to build and use in order to run trsh-go:

1) Build the image:

```
docker build -t trsh .
```

2) Run the docker image and specify your ENV vars (below an example):

```
docker run -d -e TGBOT_TOKEN="123456789:qwertyuiopasdfghjklzxcb" -e TGBOT_CHATID="12345678" trsh-go
```

3) Chat with your Bot


## Contributors

Feel free to open issues or contact me directly

## License

Code distributed under MIT licence.

