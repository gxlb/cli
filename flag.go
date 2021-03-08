package cli

import (
	"io/ioutil"
	"regexp"
	"strings"
	"syscall"
)

var (
	commaWhitespace = regexp.MustCompile("[, ]+.*")
)

func flagFromEnvOrFile(envVars []string, filePath string) (val string, ok bool) {
	for _, envVar := range envVars {
		envVar = strings.TrimSpace(envVar)
		if val, ok := syscall.Getenv(envVar); ok {
			return val, true
		}
	}
	for _, fileVar := range strings.Split(filePath, ",") {
		if data, err := ioutil.ReadFile(fileVar); err == nil {
			return string(data), true
		}
	}
	return "", false
}

// Serializer is used to circumvent the limitations of flag.FlagSet.Set
type Serializer interface {
	Serialize() string
}
