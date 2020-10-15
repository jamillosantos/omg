package internal

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jamillosantos/omg/config"
)

func sToIntf(strs ...string) []interface{} {
	output := make([]interface{}, len(strs))
	for i, v := range strs {
		output[i] = v
	}
	return output
}

func List(c *config.OmgConfig) ([]string, error) {
	output := make([]string, 0)
	for _, src := range c.Src {
		if strings.HasPrefix(src, "!") {
			pattern := src[1:]
			// Check if the exclusion pattern matches any files.
			for i := len(output) - 1; i >= 0; i-- {
				if match, err := filepath.Match(pattern, output[i]); match {
					// If matches, remove the file off the list.
					output = append(output[:i], output[i+1:]...)
				} else if err != nil {
					return nil, err
				}
			}
			continue
		}

		matches, err := filepath.Glob(src)
		if err != nil {
			return nil, err
		}

		for _, m := range matches {
			stat, err := os.Stat(m)
			if err != nil {
				return nil, err
			}
			if stat.IsDir() {
				err := filepath.Walk(m, func(path string, info os.FileInfo, err error) error {
					if filepath.Ext(path) == ".proto" {
						output = append(output, path)
					} else if err != nil {
						return err
					}

					return nil
				})
				if err != nil {
					return nil, err
				}
			} else {
				output = append(output, m)
			}
		}
	}
	return output, nil
}
