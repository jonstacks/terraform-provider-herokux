package connect

import (
	"encoding/json"
	"strings"
	"time"
)

// PossiblyEmptyTime works exactly like time, however in the case of an empty string, it also returns a nil pointer and no error.
type PossiblyEmptyTime struct {
	time.Time
}

// UnmarshalJSON first checks if the string is empty, if so it sets the pointer to nil and returns no error. Otherwise, it functions
// exactly like time.Time's UnMarshall
func (t *PossiblyEmptyTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}

	err = json.Unmarshal([]byte(b), &t.Time)
	return
}
