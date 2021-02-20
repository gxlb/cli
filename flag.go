package cli

type Flag struct {
	LogicName string   // logic name of the flag
	Name      string   // name of the flag
	Aliases   []string // aliases of the flag
	Usage     string   // usage string
	Required  bool     // if required
	Hidden    bool     // hidden this flag
	EnvVars   []string // environment values
	FilePath  string   // file path
	Value     Value    // value of this flag
	names     []string
	logicName string
}

type FlagName struct {
	LogicName string   // logic name of the flag
	Names     []string // name+aliases of the flag
	Usage     string   // usage string
	Required  bool     // if required
	Hidden    bool     // hidden this flag
	EnvVars   []string // environment values
	FilePath  string   // file path
	Value     Value    // value of this flag
}
