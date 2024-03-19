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
curl --location --request GET 'localhost:3000/{API_KEY}' \
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

```bash
curl --location --request GET 'localhost:3000/{API_KEY}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": {
        "text": "/tokens",
        "chat": {
            "id": 42
        }
    }
}'
```
