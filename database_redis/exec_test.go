package database_redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExec(t *testing.T) {
	tests:=[]struct{
		name string
		wantErr bool
	}{
		{"base-case",false},
	}

	for _,tt:=range tests{
		t.Run(tt.name, func(t *testing.T) {
			err:=Exec()
			assert.EqualValues(t, tt.wantErr,err!=nil)
		})
	}
}
