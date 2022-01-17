package converter

import (
	"os"

	converter2 "github.com/gotgcalls/converter"
)

var converter = converter2.New(nil)

func Convert(input string) (string, error) {
	output := input + ".raw"
	if _, err := os.Stat(output); err == nil {
		return output, nil
	}
	err := converter.Convert(input, output)
	return output, err
}
