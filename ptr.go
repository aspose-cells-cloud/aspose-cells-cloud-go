package asposecellscloud

// Pointer helpers for constructing model objects whose primitive fields are
// pointers so that omitted values are serialized as null rather than zero.

// Int32Ptr returns a pointer to the given int32 value.
func Int32Ptr(v int32) *int32 {
	return &v
}

// Int64Ptr returns a pointer to the given int64 value.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Float64Ptr returns a pointer to the given float64 value.
func Float64Ptr(v float64) *float64 {
	return &v
}

// BoolPtr returns a pointer to the given bool value.
func BoolPtr(v bool) *bool {
	return &v
}
