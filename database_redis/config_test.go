package database_redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {
	tests:=[]struct{
		name string
		wantErr bool
	}{
		{"base-case",false},
	}

	for _,tt:=range tests{
		t.Run(tt.name, func(t *testing.T) {
			_,err:=Setup()
			assert.EqualValues(t, tt.wantErr,err!=nil)
		})
	}
}
