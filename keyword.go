package xlang

var keywords = func() map[string]bool {
	ret := make(map[string]bool)

	for _, k := range []string{
		"func",
		"var",
		"import",
		"for",
		"if",
		"else",
		"return",
	} {
		ret[k] = true
	}

	return ret
}()

func isKeyword(s string) bool {
	return keywords[s]
}
