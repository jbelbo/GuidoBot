package Telegram

// https://core.telegram.org/bots/api#update
type WebhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// https://core.telegram.org/bots/api#sendmessage
type MessageResponse struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
