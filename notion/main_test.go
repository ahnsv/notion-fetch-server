package notion

import (
	"reflect"
	"testing"

	"github.com/kjk/notionapi"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
		want *notionapi.Page
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
