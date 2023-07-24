package command_line_argument

// CommandLineArgumentInfo is a structure that has command line argument information.
type CommandLineArgumentInfo struct {
	FlagName      string
	Usage         string
	DefaultValue  interface{}
	valueOriginal interface{}
	value         interface{}
}