package validation

// ValidationDetail is used to describe any issues in validating a struct
type ValidationDetail struct {
	Location []string
	Message  string
	Value    any
}

func (v ValidationDetail) Error() string {
	return v.Message
}
