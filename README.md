# GuidoBot
A simple Telegram bot

## Installation

```bash
cp .env.example .env
```

## Usage
```bash
go run .
```

## Webhook request example
```bash
curl --location --request GET 'localhost:4000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": {
        "text": "pepe",
        "chat": {
            "id": 42
        },
        "entities": [
            {
                "type": "mention"
            },
            {
                "type": "bold"
            }
        ]
    }
}'
```
