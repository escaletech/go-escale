package stringer

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/escaletech/go-escale/messages"
)

// Extract JSON from a text block
func JsonFromText(s string) (map[string]interface{}, error) {
	re := regexp.MustCompile(`\{([^\[\]]*)\}`)
	matches := re.FindAllStringSubmatch(s, 1)

	if len(matches) == 0 {
		return nil, errors.New(messages.InputParamNotParseable)
	}

	target := matches[0][0]
	replaced := strings.ReplaceAll(target, "\\", "")

	var res map[string]interface{}
	if err := json.Unmarshal([]byte(replaced), &res); err != nil {
		return nil, err
	}

	return res, nil
}
