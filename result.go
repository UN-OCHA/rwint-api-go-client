package rwapi

import "encoding/json"

// A ResultItem stores a result item (resource).
type ResultItem struct {
	Id     string          `json:"id"`
	Score  float64         `json:"score"`
	Href   string          `json:"href"`
	Fields json.RawMessage `json:"fields"`
}

// A ResultFacetData contains the value of the facet item
// and the number of occurences it appears in the set of resources
// matching the API query.
type ResultFacetData struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// A ResultFacet contains the computed aggreation data for
// the field defined in the corresponding facet query.
type ResultFacet struct {
	Type    string             `json:"type"`
	Data    []*ResultFacetData `json:"data"`
	Missing int                `json:"missing"`
	More    bool               `json:"more"`
}

// A ResultEmbedded contains the list of facetted results
// resulting of the API query.
type ResultEmbedded struct {
	Facets map[string]*ResultFacet `json:"facets"`
}

// A Result represents the response payload.
type Result struct {
	TotalCount int             `json:"totalCount"`
	Count      int             `json:"count"`
	Data       []*ResultItem   `json:"data"`
	Embedded   *ResultEmbedded `json:"embedded"`
}
