package client

import (
	"crypto/tls"
	"github.com/Eitol/yapo/client/certs"
	"github.com/Eitol/yapo/client/errordict"
	"github.com/Eitol/yapo/pkg/tlsconfig"
	"net/http"
)

func buildHttpClient() (*http.Client, error) {
	tlsConfig, err := buildCfg()
	if err != nil {
		return nil, err
	}
	return buildClientFromTlSConfig(tlsConfig), nil
}

func buildClientFromTlSConfig(tlsConfig *tls.Config) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}
}

func buildCfg() (*tls.Config, error) {
	tlsConfig, err := tlsconfig.NewClientValidationTLSConfig(
		certs.ClientCert,
		certs.ClientKey,
		certs.ServerCert,
	)
	if err != nil {
		return nil, errordict.ErrUnableToBuildTLSConfig.Cause(err)
	}
	return tlsConfig, nil
}
