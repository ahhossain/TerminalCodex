package history

import (
	"reflect"
	"testing"
)

func TestProcessFile(t *testing.T) {
	lines := GetHistory()
	if reflect.TypeOf(lines) != reflect.TypeOf([]string{}) {
		t.Errorf("ProcessLines() did not return a list of strings")
	}
}
