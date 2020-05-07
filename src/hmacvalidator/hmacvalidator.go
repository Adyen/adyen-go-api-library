package hmacvalidator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/adyen/adyen-go-api-library/src/notification"
	"regexp"
	"strconv"
	"strings"
)

func CalculateHmac(data interface{}, secret string) (string, error) {
	switch data.(type) {
	case string:
		return encode(data.(string), secret)
	default:
		src := GetDataToSign(data)
		return encode(src, secret)
	}
}

func ValidateHmac(notificationRequestItem notification.NotificationRequestItem, key string) bool {
	expectedSign, err := CalculateHmac(notificationRequestItem, key)
	if err != nil {
		return false
	}
	merchantSign := (*notificationRequestItem.AdditionalData)["HmacSignature"]
	return expectedSign == merchantSign
}

func GetDataToSign(notificationRequestItem interface{}) (res string) {
	switch isNotificationRequestItem(notificationRequestItem) {
	case true:
		item := notificationRequestItem.(notification.NotificationRequestItem)
		signedDataList := []string{
			item.PspReference,
			item.OriginalReference,
			item.MerchantAccountCode,
			item.MerchantReference,
			strconv.Itoa(int(item.Amount.Value)),
			item.Amount.Currency,
			item.EventCode,
			item.Success,
		}
		res = strings.Join(signedDataList, ":")
	case false:
		item := notificationRequestItem.(map[string]string)
		keys := make([]string, 0)
		values := make([]string, 0)

		for k, v := range item {
			keys = append(keys, replacer(k))
			values = append(values, replacer(v))
		}

		res = strings.Join(keys, ":") + ":" + strings.Join(values, ":")
	}
	return
}

func encode(data string, secret string) (string, error) {
	key, err := hex.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to generate HMAC: %s", err.Error())
	}
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func isNotificationRequestItem(item interface{}) bool {
	switch item.(type) {
	case notification.NotificationRequestItem:
		return true
	default:
		return false
	}
}

func replacer(s string) (str string) {
	var re1 = regexp.MustCompile(`\\`)
	var re2 = regexp.MustCompile(`:`)
	str = re1.ReplaceAllString(s, "\\\\")
	str = re2.ReplaceAllString(str, "\\:")
	return
}
