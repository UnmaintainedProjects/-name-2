package utils

func ToMTProto(chatId int64) int64 {
	return ((chatId * -1) - 1000000000000)
}

func ToBotAPI(channelId int64) int64 {
	return (1000000000000 + channelId) * -1
}
