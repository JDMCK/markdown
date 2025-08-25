package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

type BlockType int

const (
	Paragraph  = iota // nothing
	Heading           // #
	Thematic          // ---, ===, ___
	Blockquote        // > (continuous)
	UList             // *, -, +
	OList             // %d.
	CodeBlock         //  ```, ~~~
)

type Block struct {
	Type     BlockType
	Content  []string
	Children []*Block
}

func Parse(md string) string {
	lines := strings.Split(md, "\n")

	var captures []string
	var sb strings.Builder

	for _, line := range lines {
		reg := regexp.MustCompile("(#{1,6}) (.*)")
		captures = reg.FindStringSubmatch(line)
		if len(captures) >= 3 {
			n := len(captures[1])
			sb.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n", n, captures[2], n))
		}
	}
	return sb.String()
}

func parseBlocks(lines []string) *Block {
	var root Block

	lines := strings.Split(md, "\n")

	var currentType BlockType
	for i, line := range lines {

	}
}
