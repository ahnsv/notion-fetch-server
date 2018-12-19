package notion

import (
	"reflect"
	"testing"

	"github.com/kjk/notionapi"
)

func TestGetNotionPage(t *testing.T) {
	type args struct {
		c      *notionapi.Client
		pageID string
	}
	tests := []struct {
		name    string
		args    args
		want    *notionapi.Page
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNotionPage(tt.args.c, tt.args.pageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNotionPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNotionPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerializePage(t *testing.T) {
	type args struct {
		c      *notionapi.Client
		pageID string
	}
	tests := []struct {
		name    string
		args    args
		want    *notionapi.Page
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SerializePage(tt.args.c, tt.args.pageID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SerializePage() = %v, want %v", got, tt.want)
			}
		})
	}
}
