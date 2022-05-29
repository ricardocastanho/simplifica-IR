package support

import "strings"

func FilterLinesByText(lines []string, text string) []string {
	var fl []string

	for _, line := range lines {
		if strings.Contains(line, text) {
			fl = append(fl, line)
		}
	}

	return fl
}
