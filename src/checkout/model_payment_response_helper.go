/*
 * PaymentResponse helper
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. 
 *
 */

package checkout

import (
	"encoding/json"
)

// UnmarshalJSON unmarshals a quoted json string to PaymentResponse struct
func (req *PaymentResponse) UnmarshalJSON(b []byte) error {

	type PaymentResponseAlias PaymentResponse

	temp := &struct {
		*PaymentResponseAlias
	}{
		PaymentResponseAlias: (*PaymentResponseAlias)(req),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	if temp.Action != nil {

		actiontype := temp.Action.(map[string]interface{})["type"].(string)

		switch actiontype {
		case "donation":
			intermediate := &struct {
				Action *CheckoutDonationAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "qrCode":
			intermediate := &struct {
				Action *CheckoutQrCodeAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "redirect":
			intermediate := &struct {
				Action *CheckoutRedirectAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "sdk":
			intermediate := &struct {
				Action *CheckoutSDKAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "threeDS2Challenge":
			intermediate := &struct {
				Action *CheckoutThreeDS2ChallengeAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "threeDS2Fingerprint":
			intermediate := &struct {
				Action *CheckoutThreeDS2FingerPrintAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "threeDS2Action", "threeDS2":
			intermediate := &struct {
				Action *CheckoutThreeDS2Action `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "await":
			intermediate := &struct {
				Action *CheckoutAwaitAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "voucher":
			intermediate := &struct {
				Action *CheckoutVoucherAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		case "oneTimePasscode":
			intermediate := &struct {
				Action *CheckoutOneTimePasscodeAction `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action

		default:
			intermediate := &struct {
				Action map[string]interface{} `json:"action"`
				*PaymentResponseAlias
			}{
				PaymentResponseAlias: (*PaymentResponseAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.Action = intermediate.Action
		}
	}

	return nil
}
