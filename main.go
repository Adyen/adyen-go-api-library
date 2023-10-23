/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"fmt"

	"github.com/adyen/adyen-go-api-library/v8/src/adyen"
	"github.com/adyen/adyen-go-api-library/v8/src/common"
)

func main() {
	client := adyen.NewClient(&common.Config{Environment: common.TestEnv})
	fmt.Println("Welcome to Adyen API Client. Env: " + client.GetConfig().Environment)
}
