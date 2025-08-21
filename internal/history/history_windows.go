package history

import (
	"log"
	"os"
	"strings"
)

func ReadFile() string {
	data, err := os.ReadFile(`C:\Projects\TerminalCodex\testing\ConsoleHost_history.txt`)
	if err != nil {
		log.Fatal(err)
	}
	data_Str := string(data)
	return data_Str
}

func GetHistory() []string {
	data_Str := ReadFile()
	lines := strings.Split(data_Str, "\n")
	var cleaned_lines []string
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			cleaned_lines = append(cleaned_lines, trimmedLine)
		}
	}
	return cleaned_lines
}
