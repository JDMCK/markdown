package markdown_test

import (
	"encoding/json"
	"fmt"
	"io"
	md "markdown"
	"os"
	"strings"
	"testing"
)

type Test struct {
	Markdown string  `json:"markdown"`
	HTML     string  `json:"html"`
	Example  float64 `json:"example"`
	Section  string  `json:"section"`
}

func TestAll(t *testing.T) {
	jsonFile, err := os.Open("tests.json")
	if err != nil {
		t.Errorf("Failed to open tests file.")
	}
	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)
	var tests []Test

	json.Unmarshal(bytes, &tests)

	succeeded := 0
	failed := 0

	var resultLogs []string

	for _, test := range tests {
		html := md.Parse(test.Markdown)
		if html != test.HTML {
			result := fmt.Sprintf("%v failed: expected %q, got %q", test.Example, test.HTML, html)
			resultLogs = append(resultLogs, result)
			result = fmt.Sprintf("%v \033[31mfailed\033[0m: expected %q, got %q", test.Example, test.HTML, html)
			t.Error(result)
			failed += 1
		} else {
			result := fmt.Sprintf("%v succeeded: expected %q, got %q", test.Example, test.HTML, html)
			resultLogs = append(resultLogs, result)
			result = fmt.Sprintf("%v \033[32msucceeded\033[0m: expected %q, got %q", test.Example, test.HTML, html)
			t.Log(result)
			succeeded += 1
		}
	}

	summary := fmt.Sprintf("succeeded: %d, failed: %d", succeeded, failed)
	resultLogs = append(resultLogs, summary)
	t.Logf("succeeded: \033[32m%d\033[0m, failed: \033[31m%d\033[0m", succeeded, failed)

	resultBytes := []byte(strings.Join(resultLogs, "\n"))
	err = os.WriteFile("test_all.log", resultBytes, 0644)
	if err != nil {
		return
	}
}

func TestHeading(t *testing.T) {
	html := md.Parse("# Title\n")
	expected := "<h1>Title</h1>\n"
	if html != expected {
		t.Errorf("\033[31mfailed\033[0m: expected %q, got %q", expected, html)
	}
}
