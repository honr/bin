package wsdir

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	dirsFilename     = os.ExpandEnv("$HOME/.tmux/DIRS")
	fallbackMatchers = map[*regexp.Regexp]string{
		regexp.MustCompile("^[A-Z]*$"): ".upper",
		regexp.MustCompile("^[a-z]*$"): ".lower",
		regexp.MustCompile("^[0-9]*$"): ".num",
	}
)

// Get looks up keys in string table in file `dirsFilename`.
func Get(sty string) ([]string, error) {
	oldDollarSTY := os.Getenv("STY")
	os.Setenv("STY", sty)
	defer func() { os.Setenv("STY", oldDollarSTY) }()
	keys := []string{sty}
	for re, tag := range fallbackMatchers {
		if re.MatchString(sty) {
			keys = append(keys, tag)
		}
	}
	matches := []string{}
	f, err := os.Open(dirsFilename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") { // Ignore comment lines.
			continue
		}
		lineChunks := strings.SplitN(line, " ", 2)
		if len(lineChunks) < 2 {
			continue
		}
		lineKey := lineChunks[0]
		// There should be only a few keys, so this is not too inefficient.
		for _, key := range keys {
			if lineKey != key {
				continue
			}
			value := lineChunks[1]
			matches = append(matches, os.ExpandEnv(value))
		}
	}
	n := len(matches)
	if n == 0 {
		return matches, fmt.Errorf("Didn't find any matches")
	}
	for i := (n - 1) / 2; i >= 0; i-- {
		matches[n - i - 1], matches[i] = matches[i], matches[n - i -1]
	}
	return matches, nil
}
