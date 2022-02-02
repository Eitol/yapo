package client

type APIProxy interface {
	DoRequest(resource string) ([]byte, error)
}

type httpAPIEndpointProxy struct {
	url string
}

func (p *httpAPIEndpointProxy) DoRequest(resource string) ([]byte, error) {
	return doRequest(resource)
}

func newHttpApiProxy() *httpAPIEndpointProxy {
	return &httpAPIEndpointProxy{
		url: APIHost,
	}
}
