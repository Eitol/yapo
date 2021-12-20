package client

type ApiProxy interface {
	DoRequest(resource string) ([]byte, error)
}

type httpApiEndpointProxy struct {
	url string
}

func (p *httpApiEndpointProxy) DoRequest(resource string) ([]byte, error) {
	return doRequest(resource)
}

func newHttpApiProxy() *httpApiEndpointProxy {
	return &httpApiEndpointProxy{
		url: APIHost,
	}
}
