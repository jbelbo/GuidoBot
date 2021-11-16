package Telegram

// https://core.telegram.org/bots/api#update
type WebhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
		From struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		}
		Entities []MessageEntity `json:"entities"`
	} `json:"message"`
}

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type string `json:"type"`
}

// https://core.telegram.org/bots/api#sendmessage
type MessageResponse struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
