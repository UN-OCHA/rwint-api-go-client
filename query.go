package rwapi

// A QueryFields is used to include or exclude fields from the response.
type QueryFields struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

// A QueryQuery is used to perform a full text search.
type QueryQuery struct {
	Fields   []string `json:"fields,omitempty"`
	Value    string   `json:"fields,omitempty"`
	Operator string   `json:"operator,omitempty"`
}

// A Query is used to built the request payload.
type Query struct {
	Fields  *QueryFields `json:"fields,omitempty"`
	Limit   int          `json:"limit,omitempty"`
	Offset  int          `json:"offset,omitempty"`
	Sort    []string     `json:"sort,omitempty"`
	Preset  string       `json:"preset,omitempty"`
	Profile string       `json:"profile,omitempty"`
	Query   *QueryQuery  `json:"query,omitempty"`
	Filter  *Filter      `json:"filter,omitempty"`
	Facets  []*Facet     `json:"facets,omitempty"`
}

// NewQuery creates a new ReliefWeb API query payload.
func NewQuery() *Query {
	return &Query{}
}

// SetFields sets the fields to include/exclude from the response.
func (q *Query) SetFields(include []string, exclude []string) {
	if q.Fields == nil {
		q.Fields = &QueryFields{}
	}
	if len(include) != 0 {
		q.Fields.Include = include
	}
	if len(exclude) != 0 {
		q.Fields.Exclude = exclude
	}
}

// SetRange sets the limit and offset of the items to return.
func (q *Query) SetRange(limit, offset int) {
	q.SetLimit(limit)
	q.SetOffset(offset)
}

// SetLimit sets the limit of items to return.
func (q *Query) SetLimit(limit int) {
	q.Limit = limit
}

// SetOffset sets the offset from which to return items.
func (q *Query) SetOffset(offset int) {
	q.Offset = offset
}

// SetPreset sets the query preset.
func (q *Query) SetPreset(preset string) {
	q.Preset = preset
}

// SetProfile sets the query profile.
func (q *Query) SetProfile(profile string) {
	q.Profile = profile
}

// AddSort sets a sort option (several can be added to sort sequentially).
func (q *Query) AddSort(field, direction string) {
	if q.Sort == nil {
		q.Sort = []string{}
	}
	q.Sort = append(q.Sort, field+":"+direction)
}

// SetQuery sets the search query.
func (q *Query) SetQuery(query string, fields []string, operator string) {
	if q.Query == nil {
		q.Query = &QueryQuery{}
	}
	if len(fields) != 0 {
		q.Query.Fields = fields
	}
	if operator != "" {
		q.Query.Operator = operator
	}
	q.Query.Value = query
}

// SetFilter sets the filter (single or conditional type).
func (q *Query) SetFilter(filter *Filter) {
	q.Filter = filter.Flatten()
}

// AddFacet adds a facet query to the query payload.
func (q *Query) AddFacet(facet *Facet) {
	if q.Facets == nil {
		q.Facets = []*Facet{}
	}
	q.Facets = append(q.Facets, facet)
}
