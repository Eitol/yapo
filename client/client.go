package client

import (
	"encoding/json"
	"fmt"
	"github.com/Eitol/yapo/client/errordict"
	"github.com/Eitol/yapo/client/schema"
	"github.com/Eitol/yapo/pkg/iocloser"
	"io/ioutil"
)

const (
	APIVersion             = "v1.4"
	APIHost                = "https://apps.yapo.cl/api/" + APIVersion + "/public"
	ListAdsEndpoint        = "/ads"
	ListCategoriesEndpoint = "/categories/filter"

	///////////////////////////

	SearchPageSize          = 200
	CategoryIDQueryParamKey = "category"
	NextPageQueryParamKey   = "o"
	LimitQueryParamKey      = "lim"
)

type ListAdsOptions struct {
	NextPage   *string
	CategoryID *string
	SearchSize uint
}

type Client interface {
	ListCategories() ([]*schema.Category, error)
	ListAds(opts ListAdsOptions) ([]*schema.Ad, string, error)
	GetAdDetail(id string) (*schema.Ad, error)
}

type Options struct {
	apiProxy ApiProxy
}

func NewClient(opts Options) Client {
	if opts.apiProxy == nil {
		opts.apiProxy = newHttpApiProxy()
	}
	return &client{opts: opts}
}

type client struct {
	nextPageID string
	opts       Options
	apiProxy   ApiProxy
}

func (c client) GetAdDetail(id string) (*schema.Ad, error) {
	//TODO implement me
	panic("implement me")
}

func (c client) ListCategories() ([]*schema.Category, error) {
	body, err := c.opts.apiProxy.DoRequest(ListCategoriesEndpoint)
	if err != nil {
		return nil, err
	}
	out, err := parseCategoryResponse(body)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c client) ListAds(opts ListAdsOptions) ([]*schema.Ad, string, error) {
	endpoint := c.buildAdListEndpoint(opts)
	body, err := c.opts.apiProxy.DoRequest(endpoint)
	if err != nil {
		return nil, "", err
	}
	out, nextPage, err := parseAdsListResponse(body)
	if err != nil {
		return nil, "", err
	}
	return out, nextPage, nil
}

func (c client) buildAdListEndpoint(opts ListAdsOptions) string {
	pageSize := c.getPageSize(opts)
	endpoint := fmt.Sprintf("%s?%s=%d",
		ListAdsEndpoint, LimitQueryParamKey, pageSize,
	)
	if opts.NextPage != nil && len(*opts.NextPage) > 0 {
		endpoint += fmt.Sprintf("&%s=%s",
			NextPageQueryParamKey, *opts.NextPage,
		)
	}
	if opts.CategoryID != nil && len(*opts.CategoryID) > 0 {
		endpoint += fmt.Sprintf("&%s=%s",
			CategoryIDQueryParamKey, *opts.CategoryID,
		)
	}
	return endpoint
}

func (c client) getPageSize(opts ListAdsOptions) uint {
	if opts.SearchSize == 0 {
		return SearchPageSize
	} else {
		return opts.SearchSize
	}
}

func doRequest(endpoint string) ([]byte, error) {
	httpClient, err := buildHttpClient()
	if err != nil {
		return nil, err
	}
	result, err := httpClient.Get(APIHost + endpoint)
	if err != nil {
		return nil, errordict.ErrExecutingTheRequest.Cause(err)
	}
	defer iocloser.Close(result.Body)
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, errordict.ErrReadingTheResponseBody.Cause(err)
	}
	return body, nil
}

func parseCategoryResponse(resp []byte) ([]*schema.Category, error) {
	var result *schema.ListCategoryResponse
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return nil, errordict.ErrUnableToParseAdSearchResponse.Cause(err)
	}
	return result.Categories, nil
}

func parseAdsListResponse(resp []byte) ([]*schema.Ad, string, error) {
	var result *schema.GetAdsResponse
	err := json.Unmarshal(resp, &result)
	if err != nil {
		return nil, "", errordict.ErrUnableToParseAdSearchResponse.Cause(err)
	}
	out, nextPage := extractAdsListFromResponse(result)
	return out, nextPage, nil
}

func extractAdsListFromResponse(result *schema.GetAdsResponse) ([]*schema.Ad, string) {
	out := make([]*schema.Ad, 0, len(result.ListAds))
	for _, ad := range result.ListAds {
		out = append(out, ad.Ad)
	}
	return out, result.NextPage
}
