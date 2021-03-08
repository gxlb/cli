package cli

import (
	"cli/internal/impl"
)

// Author represents someone who has contributed to a cli project.
type Author = impl.Author

// App is the main structure of a cli application. It is recommended that
// an app be created with the cli.NewApp() function
type App = impl.App
