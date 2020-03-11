package goexam

const (
	// DefaultLimit is
	DefaultLimit = 10
)

// BaseFilter is
type BaseFilter struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	PrefixKey string `json:"prefix_key"`
}

// LoadDefault is
func (filter *BaseFilter) LoadDefault() {
	if filter.Limit == 0 {
		filter.Limit = DefaultLimit
	}
}
