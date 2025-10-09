package substringmatchingpattern

func hasMatch(s string, p string) bool {
	if p == s {
		return true
	}

	if p == "" {
		return false
	}

	asteriskIdx := index(p, "*")
	if asteriskIdx == -1 {
		return index(s, p) != -1
	}

	pos := 0

	if asteriskIdx != 0 {
		idx := index(s, p[0:asteriskIdx])
		if idx == -1 {
			return false
		}

		pos = idx + len(p[0:asteriskIdx])
	}

	if asteriskIdx != len(p)-1 {
		idx := index(s[pos:], p[asteriskIdx+1:])
		if idx == -1 {
			return false
		}
	}

	return true
}

func index(s, substr string) int {
	for pos := 0; pos < len(s); {
		if len(substr)+pos <= len(s) {
			if s[pos:pos+len(substr)] == substr {
				return pos
			}
			pos++
		} else {
			return -1
		}
	}
	return -1
}
