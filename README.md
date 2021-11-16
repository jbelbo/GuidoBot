# GuidoBot
A simple Telegram bot


# Installation

### .env
```
PORT=
DATABASE_URL=
HEROKU=
```

# Usage

go run .

```bash
curl --location --request GET 'localhost:4000' \
--header 'Content-Type: application/json' \
--data-raw '{
    "message": {"text": "hola", "chat" : { "id" : 42 }}
}'
```

