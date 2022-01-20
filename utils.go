package dbuilder

import "strconv"

// QuoteStrings return an array of quoted string
func QuoteStrings(list []string) []string {
	result := make([]string, 0, len(list))

	for _, str := range list {
		result = append(result, strconv.Quote(str))
	}

	return result
}
