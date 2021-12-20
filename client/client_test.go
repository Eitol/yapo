package client

import (
	"github.com/Eitol/yapo/client/schema"
	"testing"
)

func TestClientListCategories(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewClient(Options{})
		got, err := c.ListCategories()
		if err != nil {
			t.Errorf("ListCategories() error = %v", err)
			return
		}
		if len(got) == 0 {
			t.Errorf("ListCategories() = %v, want %v", got, "not empty")
		}
	})
}

func TestClientListAdsUntil(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		c := NewClient(Options{})
		nextPage := ""
		for i := 0; i < 5; i++ {
			var err error
			var adsList []*schema.Ad
			adsList, nextPage, err = c.ListAds(ListAdsOptions{
				NextPage:   &nextPage,
				SearchSize: 10,
			})
			if err != nil {
				t.Errorf("ListAds() error = %v", err)
				return
			}
			if len(adsList) == 0 {
				t.Errorf("ListAds() error = empty list")
				return
			}
			if len(nextPage) == 0 {
				t.Errorf("ListAds() error = empty nextPage")
				return
			}
		}

	})

}
