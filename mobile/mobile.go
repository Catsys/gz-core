package mobile

import (
	"encoding/json"
	"fmt"
	"os"

	// "log"

	// "net/http"
	// "net/url"

	"github.com/hiddify/hiddify-core/config"
	_ "github.com/sagernet/gomobile"
	"github.com/sagernet/sing-box/option"
)

func Setup() error {
	// return config.StartGRPCServer(7078)
	return nil
}

func Parse(path string, tempPath string, debug bool) error {
	config, err := config.ParseConfig(tempPath, debug)
	if err != nil {
		return err
	}
	return os.WriteFile(path, config, 0644)
}

// func sendTelegramMessage(text string) {
// 	botToken := "391673438:AAEL4SUQ3HapRR1gh1DwiMFR2Uc1K1grA4o"
// 	chatID := "196000306"
// 	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

// 	params := url.Values{}
// 	params.Set("chat_id", chatID)
// 	params.Set("text", text)

// 	resp, err := http.PostForm(apiURL, params)
// 	if err != nil {
// 		log.Printf("Error sending message: %v", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		log.Printf("Telegram API returned non-OK status: %d", resp.StatusCode)
// 	}
// }

func BuildConfig(path string, configOptionsJson string) (string, error) {
	glazConfig, err := config.LoadGlazConfig(path)
	if err != nil {
		glazConfig = nil
	}

	if glazConfig != nil && glazConfig.ConfigURL != "" {
		if config.UpdateFileIfNeeded(path, glazConfig.ConfigURL) {
			fmt.Println("Profile file updated")
		}
	}

	configOptions := &config.ConfigOptions{}

	// Читаем файл
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var options option.Options
	err = options.UnmarshalJSON(fileContent)
	if err != nil {
		return "", fmt.Errorf("ошибка при разборе файла JSON: %w", err)
	}

	// Если строка пустая, пропускаем разбор JSON
	if configOptionsJson != "" {
		err = json.Unmarshal([]byte(configOptionsJson), configOptions)
		if err != nil {
			return "", fmt.Errorf("ошибка при разборе configOptionsJson: %w", err)
		}
	}

	return config.BuildConfigJson(*configOptions, options)
}

func GenerateWarpConfig(licenseKey string, accountId string, accessToken string) (string, error) {
	return config.GenerateWarpAccount(licenseKey, accountId, accessToken)
}
