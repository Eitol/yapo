package schema

type ListCategoryResponse struct {
	Categories []*Category `json:"categories"`
}

type Category struct {
	SubCategories []*Category `json:"categories,omitempty"`
	Color         string      `json:"color"`
	Code          string      `json:"code"`
	FilterValue   FilterValue `json:"filter_value"`
	Icon          *string     `json:"icon"`
	Label         string      `json:"label"`
}

type FilterValue struct {
	Keys   []string `json:"keys"`
	Values []string `json:"values"`
}
