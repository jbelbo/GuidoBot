# GuidoBot
A simple Telegram bot

# Installation

export DATABASE_URL= ......
export PORT=4000


# Usage

go run .

curl --location --request GET 'localhost:4000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": {"text": "hola", "chat" : { "id" : 42 }}
}'

