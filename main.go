/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"fmt"

	api "github.com/adyen/adyen-go-api-library/src/api"
	common "github.com/adyen/adyen-go-api-library/src/common"
)

func main() {
	client := api.NewClient(&common.Config{Environment: common.TestEnv})
	fmt.Println("Welcome to Adyen API Client. Env: " + client.GetConfig().Environment)
}
