package hmacvalidator

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/adyen/adyen-go-api-library/v19/src/webhook"
)

// CalculateHmac calculates the SHA-256 HMAC for the given data and key
func CalculateHmac(data interface{}, secret string) (string, error) {
	switch val := data.(type) {
	case string:
		return encode(val, secret)
	default:
		src := GetDataToSign(data)
		return encode(src, secret)
	}
}

// ValidateHmac validates the HMAC signature of the NotificationRequestItem object. Use for webhooks that provide the
// hmacSignature as part of the payload `AdditionalData` (i.e. Payments)
//
// Params:
// notificationRequestItem: NotificationRequestItem object
// key: HMAC key to generate the signature
func ValidateHmac(notificationRequestItem webhook.NotificationRequestItem, key string) bool {
	expectedSign, err := CalculateHmac(notificationRequestItem, key)
	if err != nil {
		return false
	}
	merchantSign := (*notificationRequestItem.AdditionalData)["hmacSignature"]
	return expectedSign == merchantSign
}

// ValidateHmacPayload validates the HMAC signature of a payload against an expected signature. Use for webhooks that provide the
// hmacSignature in the HTTP header (i.e. Banking, Management API)
//
// Params:
// hmacSignature: HMAC signature to validate
// key: HMAC key to generate the signature
// payload: webhook payload
func ValidateHmacPayload(hmacSignature string, key string, payload string) bool {
	expectedSign, err := CalculateHmac(payload, key)
	if err != nil {
		return false
	}
	return expectedSign == hmacSignature
}

// GetDataToSign converts a notificationRequestItem to string, which later on can be used for calculating a HMAC
func GetDataToSign(notificationRequestItem interface{}) string {
	switch item := notificationRequestItem.(type) {
	case webhook.NotificationRequestItem:
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
		return strings.Join(signedDataList, ":")
	case map[string]string:
		keys := make([]string, 0)
		values := make([]string, 0)

		for k := range item {
			keys = append(keys, replacer(k))
		}
		sort.Strings(keys)
		for _, k := range keys {
			values = append(values, replacer(item[k]))
		}

		return strings.Join(keys, ":") + ":" + strings.Join(values, ":")
	default:
		return ""
	}
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

func replacer(s string) string {
	re1 := regexp.MustCompile(`\\`)
	re2 := regexp.MustCompile(`:`)
	str := re1.ReplaceAllString(s, "\\\\")
	str = re2.ReplaceAllString(str, "\\:")
	return str
}
