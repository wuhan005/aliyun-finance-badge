// Copyright 2022 E99p1ant. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handler

import (
	"bytes"
	"io"
	"net/http"

	"github.com/narqo/go-badge"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currency, balance, err := getBalance()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	status := currency + " " + balance
	subject := "阿里云账户余额"
	lang := r.URL.Query().Get("lang")
	if lang == "en" {
		subject = "Aliyun Balance"
	}

	badgeBytes, err := badge.RenderBytes(subject, status, "#ff6a00")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = io.Copy(w, bytes.NewReader(badgeBytes))
}
