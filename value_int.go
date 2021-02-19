package cli

type IntValue struct {
	Value       int
	Destination *int
	Default     int
	DefaultText string
	Enums       []int
	Min         int
	Max         int
	hasBeenSet  bool
}
