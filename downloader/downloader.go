package downloader

import (
	"io"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func Download(b *gotgbot.Bot, audio *gotgbot.Audio) (string, error) {
	output := audio.FileUniqueId + ".dl"
	if _, err := os.Stat(output); err == nil {
		return output, nil
	}
	file, err := b.GetFile(audio.FileId)
	if err != nil {
		return "", err
	}
	downloadUrl := "https://api.telegram.org/file/bot" + b.Token + "/" + file.FilePath
	res, err := http.Get(downloadUrl)
	if err != nil {
		return "", err
	}
	file2, err := os.Create(output)
	if err != nil {
		return "", err
	}
	defer file2.Close()
	_, err = io.Copy(file2, res.Body)
	return output, err
}
