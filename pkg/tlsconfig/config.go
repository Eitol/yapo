package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/Eitol/yapo/pkg/cerrors"
	"io/ioutil"
)

var (
	ErrLoadingX509KeyPair = cerrors.Error{
		Code: "ErrLoadingX509KeyPair",
	}

	ErrReadingTheLocalCaCertificate = cerrors.Error{
		Code: "ErrReadingTheLocalCaCertificate",
	}
)

func NewClientValidationTLSConfig(clientCertFile, clientKeyFile, caCertFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, ErrLoadingX509KeyPair.Cause(err)
	}

	caCertPool, err := buildCaCertPool(caCertFile)
	if err != nil {
		return nil, err
	}

	tlsConfig := buildTlsConfig(cert, caCertPool)
	return tlsConfig, nil
}

func buildCaCertPool(caCertFile string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, ErrReadingTheLocalCaCertificate.Cause(err).WithMeta(map[string]string{
			"caCertFile": caCertFile,
		})
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool, nil
}

func buildTlsConfig(cert tls.Certificate, caCertPool *x509.CertPool) *tls.Config {
	return &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
}
