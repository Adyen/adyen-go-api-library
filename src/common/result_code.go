/*
 * Adyen API
 *
 * Contact: support@adyen.com
 */

package common

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ResultCode int

const (
	AuthenticationFinished ResultCode = iota
	AuthenticationNotRequired
	Authorised
	Cancelled
	ChallengeShopper
	Error
	IdentifyShopper
	Pending
	PresentToShopper
	Received
	RedirectShopper
	Refused
	Success
)

// ResultCodeValues should be in same order as the const defined above
var ResultCodeValues = [...]string{
	"AuthenticationFinished",
	"AuthenticationNotRequired",
	"Authorised",
	"Cancelled",
	"ChallengeShopper",
	"Error",
	"IdentifyShopper",
	"Pending",
	"PresentToShopper",
	"Received",
	"RedirectShopper",
	"Refused",
	"Success",
}

func (s ResultCode) String() string {
	return ResultCodeValues[s]
}

func ResultCodeFromString(value string) (ResultCode, error) {
	for p, v := range ResultCodeValues {
		if v == value {
			return ResultCode(p), nil
		}
	}
	return ResultCode(0), fmt.Errorf("ResultCode enum not found for %s", value)
}

// MarshalJSON marshals the enum as a quoted json string
func (s *ResultCode) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(s.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmarshals a quoted json string to the enum value
func (s *ResultCode) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	r, err := ResultCodeFromString(j)
	if err != nil {
		return err
	}
	*s = r
	return nil
}
