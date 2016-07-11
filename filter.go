package rwapi

// A FilterRangeValue is the value for a range type filter.
// From and To should be numbers (float, int).
type FilterRangeValue struct {
	From interface{} `json:"from,omitempty"`
	To   interface{} `json:"to,omitempty"`
}

// ReliefWeb API query filter.
type Filter struct {
	Operator   string      `json:"operator,omitempty"`
	Negate     bool        `json:"negate,omitempty"`
	Conditions []*Filter   `json:"conditions,omitempty"`
	Field      string      `json:"field,omitempty"`
	Value      interface{} `json:"value,omitempty"`
}

// NewFilter creates a new ReliefWeb API query filter.
func NewFilter() *Filter {
	return &Filter{}
}

// SetOperator sets the operator filter.
// Only usefull for Condition type filter or fitler with Array type value.
func (f *Filter) SetOperator(operator string) {
	if operator != "" {
		f.Operator = operator
	}
}

// SetNegate negates the filter (find items that don't match the condition).
func (f *Filter) SetNegate(negate bool) {
	f.Negate = negate
}

// SetField sets the field on which to apply the filter.
func (f *Filter) SetField(field string) {
	f.Field = field
}

// SetValue sets the filter value.
// It can be a FilterRangeValue with From and To fields
// or an array of values (string, bool, int or float)
// or a single value (string, bool, int or float).
func (f *Filter) SetValue(value interface{}) {
	f.Value = value
}

// AddCondition adds a condition to the list of conditions of the filter.
func (f *Filter) AddCondition(field string, value interface{}, operator string, negate bool) {
	filter := NewFilter()
	filter.SetField(field)
	filter.SetValue(value)
	filter.SetOperator(operator)
	filter.SetNegate(negate)
	f.AddFilter(filter)
}

// AddFilter adds a filter to the list of conditions of the filter.
func (f *Filter) AddFilter(filter *Filter) {
	if f.Conditions == nil {
		f.Conditions = []*Filter{}
	}
	// We flatten the filter before adding it to the list
	// of conditions.
	f.Conditions = append(f.Conditions, filter.Flatten())
}

// Flatten returns the first condition filter if there is only one condition
// otherwise it returns the actual filter.
func (f *Filter) Flatten() *Filter {
	if len(f.Conditions) == 1 {
		return f.Conditions[0]
	}
	return f
}
