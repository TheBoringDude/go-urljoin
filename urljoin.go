package urljoin

import (
	"regexp"
	"strings"
)

func UrlJoin(base string, urls ...string) string {
	// unsure about this feat
	if m, _ := regexp.MatchString(`^[^/:]+:\/*`, base); !m {
		base = "http://" + base // prepend http:// if does not start with a protocol
	}

	if len(urls) == 0 {
		return base
	}

	var result []string
	var components []string

	if match, _ := regexp.MatchString(`^[^/:]+:\/*$`, base); match && len(urls) > 1 {
		components = append(append(components, base+urls[0]), urls[:1]...)
	} else {
		components = append(append(components, base), urls...)
	}

	for i, val := range components {
		if val == "" {
			continue
		}

		if i != 0 {
			r := regexp.MustCompile(`^[\/]+`)
			val = r.ReplaceAllString(val, "")
		}

		if i < len(components)-1 {
			r := regexp.MustCompile(`[\/]+$`)
			val = r.ReplaceAllString(val, "")
		} else {
			r := regexp.MustCompile(`[\/]+$`)
			val = r.ReplaceAllString(val, "/")
		}

		result = append(result, val)
	}
	var str = strings.Join(result, "/")
	r := regexp.MustCompile(`\/(\?|&|#[^!])`)
	str = r.ReplaceAllString(str, "$1")

	var parts = strings.Split(str, "?")
	var denom string
	if len(parts) > 1 {
		denom = "?"
	} else {
		denom = ""
	}

	str = parts[0] + denom + strings.Join(parts[1:], "&")

	return str
}
