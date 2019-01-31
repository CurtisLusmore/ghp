package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_convertMarkdown(t *testing.T) {
	type args struct {
		mdText []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "Test 1",
			args:    args{mdText: []byte("*test*")},
			want:    []byte("<p><em>test</em></p>"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertMarkdown(tt.args.mdText)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertMarkdown() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.want = bytes.TrimSpace(tt.want)
			got = bytes.TrimSpace(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
