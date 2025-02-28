package utils

import (
	"context"
	"strings"

	"github.com/spf13/viper"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Encode converts a int64 number to a Base62 string.
func EncodeToBase62(num int64) string {
	if num == 0 {
		return "0"
	}

	var encoded strings.Builder
	base := int64(len(base62Chars))

	for num > 0 {
		remainder := num % base
		encoded.WriteByte(base62Chars[remainder])
		num /= base
	}

	// Reverse the encoded string
	runes := []rune(encoded.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// Decode converts a Base62 string back to a int64 number.
func DecodeFromBase62(encoded string) int64 {
	var num int64
	base := int64(len(base62Chars))

	for _, char := range encoded {
		num = num*base + int64(strings.IndexRune(base62Chars, char))
	}

	return num
}

func LoadConfig(serviceName string, config any) error {
	viper.SetConfigName(serviceName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(config); err != nil {
		return err
	}

	return nil
}

func ProvideBackgroundContext() context.Context {
	return context.Background()
}
