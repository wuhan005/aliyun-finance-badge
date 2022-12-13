// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handler

import (
	"os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/pkg/errors"
)

func getBalance() (string, string, error) {
	accessKeyID := os.Getenv("ALIYUN_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALIYUN_ACCESS_KEY_SECRET")
	regionID := os.Getenv("ALIYUN_REGION_ID")

	config := sdk.NewConfig()
	credential := credentials.NewAccessKeyCredential(accessKeyID, accessKeySecret)
	client, err := bssopenapi.NewClientWithOptions(regionID, config, credential)
	if err != nil {
		return "", "", errors.Wrap(err, "create client")
	}

	request := bssopenapi.CreateQueryAccountBalanceRequest()
	request.Scheme = "https"
	response, err := client.QueryAccountBalance(request)
	if err != nil {
		return "", "", errors.Wrap(err, "query account balance")
	}
	return response.Data.Currency, response.Data.AvailableAmount, nil
}
