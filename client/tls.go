package client

import (
	"crypto/tls"
	"github.com/Eitol/yapo/client/errordict"
	"github.com/Eitol/yapo/pkg/tlsconfig"
	"net/http"
)

const (
	certPath         = "../certs"
	clientCertPath   = certPath + "/client.cer.pem"
	clientKeyPath    = certPath + "/client.key.pem"
	serverCaFilePath = certPath + "/server.cer.pem"
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
		clientCertPath,
		clientKeyPath,
		serverCaFilePath,
	)
	if err != nil {
		return nil, errordict.ErrUnableToBuildTLSConfig.Cause(err)
	}
	return tlsConfig, nil
}
