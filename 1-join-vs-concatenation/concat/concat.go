package concat

import "strings"

var someStrings = []string{"francisco", "vinicius", "gouveia", "de", "freitas"}
var concatResult string

func WithOperador() *string {
	spacePrefix := " "

	for _, value := range someStrings {
		concatResult += value + spacePrefix
	}
	return &concatResult
}

func WithJoin() *string {
	joinResult := strings.Join(someStrings, " ")
	return &joinResult
}