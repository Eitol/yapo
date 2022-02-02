package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/Eitol/yapo/pkg/cerrors"
)

var (
	ErrLoadingX509KeyPair = cerrors.Error{
		Code: "ErrLoadingX509KeyPair",
	}
)

func NewClientValidationTLSConfig(clientCert, clientKey, caCert string) (*tls.Config, error) {
	cert, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
	if err != nil {
		return nil, ErrLoadingX509KeyPair.Cause(err)
	}

	caCertPool, err := buildCaCertPool(caCert)
	if err != nil {
		return nil, err
	}

	tlsConfig := buildTlsConfig(cert, caCertPool)
	return tlsConfig, nil
}

func buildCaCertPool(caCert string) (*x509.CertPool, error) {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(caCert))
	return caCertPool, nil
}

func buildTlsConfig(cert tls.Certificate, caCertPool *x509.CertPool) *tls.Config {
	return &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
}
