package utils

// TrimQuotes removes enclosing quoates in a printable string
func TrimQuotes(toTrim []byte) string {
	if string(toTrim[0]) == "\"" && string(toTrim[len(toTrim)-1]) == "\"" {
		return string(toTrim[1 : len(toTrim)-1])
	}
	return string(toTrim)
}

// FormatDate handles the case of '/' being substituted with '\/'
func FormatDate(date string) string {
	var formatted []byte

	for _, char := range date {
		if char != '\\' {
			formatted = append(formatted, byte(char))
		}
	}
	return string(formatted)
}
