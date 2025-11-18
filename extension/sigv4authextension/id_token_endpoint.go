// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"

import (
	"io"
	"net/http"
)

type IdentityTokenEndpoint string

func (j IdentityTokenEndpoint) GetIdentityToken() ([]byte, error) {
	resp, err := http.Get(string(j))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	token, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return token, nil
}
