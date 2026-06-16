package lib

func withTrailingSlash(s string) string {
	if len(s) > 0 && s[len(s)-1] != '/' {
		return s + "/"
	}

	return s
}

func withLeadingSlash(s string) string {
	if len(s) > 0 && s[0] != '/' {
		return "/" + s
	}
	return s
}
