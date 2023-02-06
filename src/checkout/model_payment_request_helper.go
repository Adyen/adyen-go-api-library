/*
 * PaymentRequest helper
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments.
 *
 */

package checkout

import (
	"encoding/json"
)

// UnmarshalJSON unmarshals a quoted json string to PaymentRequest struct
func (req *PaymentRequest) UnmarshalJSON(b []byte) error {

	type PaymentRequestAlias PaymentRequest

	temp := &struct {
		*PaymentRequestAlias
	}{
		PaymentRequestAlias: (*PaymentRequestAlias)(req),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	if temp.PaymentMethod != nil {

		pmtype := temp.PaymentMethod.(map[string]interface{})["type"].(string)

		switch pmtype {
		case "scheme":
			intermediate := &struct {
				PaymentMethod *CardDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod
		case "ach":
			intermediate := &struct {
				PaymentMethod *AchDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "amazonpay":
			intermediate := &struct {
				PaymentMethod *AmazonPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "androidpay":
			intermediate := &struct {
				PaymentMethod *AndroidPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "applepay":
			intermediate := &struct {
				PaymentMethod *ApplePayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "directdebit_GB":
			intermediate := &struct {
				PaymentMethod *BacsDirectDebitDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "billdesk_online":
			intermediate := &struct {
				PaymentMethod *BillDeskDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "billdesk_wallet":
			intermediate := &struct {
				PaymentMethod *BillDeskDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "blik":
			intermediate := &struct {
				PaymentMethod *BlikDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "cellulant":
			intermediate := &struct {
				PaymentMethod *CellulantDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "doku_mandiri_va", "doku_cimb_va", "doku_danamon_va", "doku_bni_va", "doku_permata_lite_atm", "doku_permata_atm", "doku_bri_va", "doku_bca_va", "doku_alfamart", "doku_indomaret", "doku_sinarmas_va":
			intermediate := &struct {
				PaymentMethod *DokuDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "dotpay":
			intermediate := &struct {
				PaymentMethod *DotpayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "dragonpay_ebanking", "dragonpay_otc_banking", "dragonpay_otc_non_banking", "dragonpay_otc_philippines":
			intermediate := &struct {
				PaymentMethod *DragonpayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "econtext_seveneleven", "econtext_stores":
			intermediate := &struct {
				PaymentMethod *EcontextVoucherDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "giropay":
			intermediate := &struct {
				PaymentMethod *GiropayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "paywithgoogle":
			intermediate := &struct {
				PaymentMethod *GooglePayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "ideal":
			intermediate := &struct {
				PaymentMethod *IdealDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "klarna", "klarnapayments", "klarnapayments_account", "klarnapayments_b2b", "klarna_paynow", "klarna_account", "klarna_b2b":
			intermediate := &struct {
				PaymentMethod *KlarnaDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "masterpass":
			intermediate := &struct {
				PaymentMethod *MasterpassDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "mbway":
			intermediate := &struct {
				PaymentMethod *MbwayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "molpay_ebanking_fpx_MY", "molpay_ebanking_TH", "molpay_ebanking_VN", "molpay_ebanking_MY", "molpay_ebanking_direct_MY", "molpay_fpx":
			intermediate := &struct {
				PaymentMethod *MolPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "mobilepay":
			intermediate := &struct {
				PaymentMethod *MobilePayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "paypal":
			intermediate := &struct {
				PaymentMethod *PayPalDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "payu_IN_upi":
			intermediate := &struct {
				PaymentMethod *PayUUpiDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "samsungpay":
			intermediate := &struct {
				PaymentMethod *SamsungPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "sepadirectdebit":
			intermediate := &struct {
				PaymentMethod *SepaDirectDebitDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "vipps":
			intermediate := &struct {
				PaymentMethod *VippsDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "visacheckout":
			intermediate := &struct {
				PaymentMethod *VisaCheckoutDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "wechatpay":
			intermediate := &struct {
				PaymentMethod *WeChatPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "wechatpayMiniProgram":
			intermediate := &struct {
				PaymentMethod *WeChatPayMiniProgramDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		default:
			intermediate := &struct {
				PaymentMethod map[string]interface{} `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod
		}
	}

	return nil
}
