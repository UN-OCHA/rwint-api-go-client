package rwapi

// A Facet corresponds to a query to compute aggregation.
type Facet struct {
	Name     string  `json:"name,omitempty"`
	Field    string  `json:"field,omitempty"`
	Limit    int     `json:"limit,omitempty"`
	Interval string  `json:"interval,omitempty"`
	Sort     string  `json:"sort,omitempty"`
	Scope    string  `json:"scope,omitempty"`
	Filter   *Filter `json:"filter,omitempty"`
}

// NewFacet creates a new ReliefWeb API facet for the given field.
func NewFacet() *Facet {
	return &Facet{}
}

// SetName sets the name of the facets for easy retrieval in the results.
// If not defined, the field name will be used.
func (f *Facet) SetName(name string) {
	f.Name = name
}

// SetField sets the field on which to compute facets. It is mandatory.
func (f *Facet) SetField(field string) {
	f.Field = field
}

// SetLimit sets the maximum number of facet results to return.
// Only for "term" type facets. Incompatible with "Interval".
func (f *Facet) SetLimit(limit int) {
	f.Limit = limit
}

// SetInterval sets the interval (year, month, day) for date facets.
// Only for "date" type facets. Incompatible with "Limit".
func (f *Facet) SetInterval(interval string) {
	f.Interval = interval
}

// SetSort sets the order of the facets results.
// The field is either "count" or "value" (name of the facet item).
// The direction is either "desc" or "asc".
func (f *Facet) SetSort(field string, direction string) {
	f.Sort = field + ":" + direction
}

// SetScope sets the scope of the facet query.
// It can be set "global" where the facet is computed against the main
// query and its eventual filters or to "query" where the facet is computed
// against the full text search query only, ignoring the main filters.
// This is in that case often used in conjunction with a facet filter
// for example to create "OR" type facets.
func (f *Facet) SetScope(scope string) {
	f.Scope = scope
}

// SetFilter set a filter on the resources on which to compute the facet.
// This filter only applies to the facet query, not the main API query.
func (f *Facet) SetFilter(filter *Filter) {
	f.Filter = filter.Flatten()
}
