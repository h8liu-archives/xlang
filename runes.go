package xlang

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func isWhite(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

func isOperator(r rune) bool {
	if r >= '!' && r <= '/' {
		return true
	}
	if r >= ':' && r <= '@' {
		return true
	}
	if r >= '[' && r <= '`' {
		return true
	}
	if r >= '{' && r <= '~' {
		return true
	}
	if r == '\n' {
		return true
	}
	return false
}
