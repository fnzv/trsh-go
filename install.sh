echo "############## Starting trsh setup ################"

echo "Installing the last trsh-go binary release from github...."
cp release/trsh /usr/bin/trsh

echo "Installing requirements...jq"
if [ -f /etc/redhat-release ]; then
  yum install jq -y
fi

if [ -f /etc/lsb-release ]; then
  apt-get install jq -y
fi

echo "Type your Telegram BOT Token, followed by [ENTER]:"

read TGBOT_TOKEN

echo "To find out your Telegram Chat ID send a message to the configured BOT and then confirm here with [ENTER]"
read tmpvar

chat_id=$(curl -ss https://api.telegram.org/bot$TGBOT_TOKEN/getUpdates | jq .result[0].message.from.id)


if [ $chat_id != "null" ]; then
  echo "Your chat id is "$chat_id
else
  echo "Chat id is null, enter manually your chat id or CTRL+Z this script and re-run after sending a message to the bot, [ENTER]:"
  read chat_id
fi


echo 'TGBOT_TOKEN="'$TGBOT_TOKEN'"' >> /etc/trsh.env
echo 'TGBOT_CHATID="'$chat_id'"' >> /etc/trsh.env

echo "Configuring trsh service.."
cp setup/trsh.service /etc/systemd/system/trsh.service
systemctl daemon-reload
systemctl start trsh
echo "Your Telegram bot is now ready to receive commands! Send /help for more information"


