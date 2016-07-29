package rwapi

import (
	"encoding/json"
	"reflect"
)

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

// GetItems returns the resource items from the response payload.
// This function takes a pointer to a slice. The resource items will
// be unserialized into the slice elements' type.
func (r *Result) GetItems(a interface{}) error {
	av := reflect.ValueOf(a).Elem()
	at := av.Type().Elem()

	for _, item := range r.Data {
		v := reflect.New(at)
		i := v.Interface()
		if err := json.Unmarshal(item.Fields, &i); err != nil {
			return err
		}
		av.Set(reflect.Append(av, reflect.Indirect(v)))
	}
	return nil
}
