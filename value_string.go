package cli

type StringValue struct {
	Value       string
	Destination *string
	Default     string
	DefaultText string
	Enums       []string
	hasBeenSet  bool
}
