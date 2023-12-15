package gobot

type Updates struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId           int                `json:"update_id"`
	Message            Message            `json:"message"`
	EditMessage        Message            `json:"edited_message"`
	ChannelPost        Message            `json:"channel_post"`
	EditChannelPost    Message            `json:"edited_channel_post"`
	InlineQuery        InlineQuery        `json:"inline_query"`
	ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      CallbackQuery      `json:"callback_query"`
	ShippingQuery      ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
	Poll               Poll               `json:"poll"`
	PollAnswer         PollAnswer         `json:"poll_answer"`
	MyChatMember       ChatMemberUpdated  `json:"my_chat_member"`
	ChatMember         ChatMemberUpdated  `json:"chat_member"`
	ChatJoinRequest    ChatJoinRequest    `json:"chat_join_request"`
}

type WebHookInfo struct {
}

type Message struct {
}

type InlineQuery struct {
}

type ChosenInlineResult struct {
}

type CallbackQuery struct {
}

type ShippingQuery struct {
}

type PreCheckoutQuery struct {
}

type Poll struct {
}

type PollAnswer struct {
}

type ChatMemberUpdated struct {
}

type ChatJoinRequest struct {
}
