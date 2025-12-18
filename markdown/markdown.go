package markdown

import "strings"

// Render translates markdown to HTML
func Render(input string) string {
	lines := strings.Split(input, "\n")
	var out strings.Builder
	inList := false

	for _, line := range lines {
		line = strings.TrimRight(line, " ")

		// Headers
		if level := headerLevel(line); level > 0 {
			closeList(&out, &inList)
			out.WriteString("<h")
			out.WriteByte(byte('0' + level))
			out.WriteString(">")
			out.WriteString(renderInline(strings.TrimSpace(line[level:])))
			out.WriteString("</h")
			out.WriteByte(byte('0' + level))
			out.WriteString(">")
			continue
		}

		// Unordered list
		if strings.HasPrefix(line, "* ") {
			if !inList {
				out.WriteString("<ul>")
				inList = true
			}
			out.WriteString("<li>")
			out.WriteString(renderInline(line[2:]))
			out.WriteString("</li>")
			continue
		}

		// Paragraph
		closeList(&out, &inList)
		if line != "" {
			out.WriteString("<p>")
			out.WriteString(renderInline(line))
			out.WriteString("</p>")
		}
	}

	closeList(&out, &inList)
	return out.String()
}

// --- helpers ---

func headerLevel(s string) int {
	level := 0
	for level < len(s) && s[level] == '#' {
		level++
	}
	if level > 0 && level <= 6 && len(s) > level && s[level] == ' ' {
		return level
	}
	return 0
}

func closeList(out *strings.Builder, inList *bool) {
	if *inList {
		out.WriteString("</ul>")
		*inList = false
	}
}

func renderInline(s string) string {
	var out strings.Builder
	i := 0

	for i < len(s) {
		switch {
		case strings.HasPrefix(s[i:], "__"):
			out.WriteString("<strong>")
			i += 2
			j := strings.Index(s[i:], "__")
			out.WriteString(s[i : i+j])
			out.WriteString("</strong>")
			i += j + 2

		case s[i] == '_':
			out.WriteString("<em>")
			i++
			j := strings.IndexByte(s[i:], '_')
			out.WriteString(s[i : i+j])
			out.WriteString("</em>")
			i += j + 1

		default:
			out.WriteByte(s[i])
			i++
		}
	}

	return out.String()
}
