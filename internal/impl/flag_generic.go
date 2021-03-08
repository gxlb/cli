package impl

// GenericValue is a generic parseable type identified by a specific flag
type GenericValue interface {
	Set(value string) error
	String() string
}
