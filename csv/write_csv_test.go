package csv

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBooks_ToCSV(t *testing.T) {
	tests := []struct {
		name    string
		books   *Books
		wantW   string
		wantErr bool
	}{
		{
			name: "base-case",
			books: &Books{
				Book{
					Author: "F Scott Fitzgerald",
					Title:  "The Great Gatsby",
				},
			},
			wantW:   "Author,Title\nF Scott Fitzgerald,The Great Gatsby\n",
			wantErr: false,
		},
	}

	for _,tt:=range tests{
		t.Run(tt.name, func(t *testing.T) {
			w:=&bytes.Buffer{}
			err:= tt.books.ToCSV(w)

			assert.Nil(t,err)

			gotW:=w.String()
			assert.EqualValues(t, tt.wantW,gotW)
		})
	}
}

func TestWriteCSVOutput(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteCSVOutput()
			assert.EqualValues(t, (err != nil),tt.wantErr)
		})
	}
}

func TestWriteCSVBuffer(t *testing.T) {
	tests := []struct {
		name    string
		want    *bytes.Buffer
		wantErr bool
	}{
		{"base-case", bytes.NewBufferString("Author,Title\nF Scott Fitzgerald,The Great Gatsby\nJ D Salinger,The Catcher in the Rye\n"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteCSVBuffer()

			assert.EqualValues(t, (err != nil),tt.wantErr)
			assert.EqualValues(t, got,tt.want)
		})
	}
}
