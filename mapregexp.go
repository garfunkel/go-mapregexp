// Package implementing regexp methods which return expression groups as a map object.
package mapregexp

import (
	"regexp"
)

// Regexp type which is able to represents expression groups as map objects.
type MapRegexp struct {
	regexp.Regexp
}

// Compile an expression pattern, asserting that it is valid.
func MustCompile(pattern string) *MapRegexp {
	return &MapRegexp{*regexp.MustCompile(pattern)}
}

// Find expression matches in a string, returning groups as a match object.
func (regex *MapRegexp) FindStringSubmatchMap(str string) map[string]string {
	submatchMap := make(map[string]string)
	submatches := regex.FindStringSubmatch(str)

	if submatches == nil {
		return nil
	}

	for index, name := range regex.SubexpNames() {
		if index == 0 || name == "" {
			continue
		}

		submatchMap[name] = submatches[index]
	}

	return submatchMap
}
