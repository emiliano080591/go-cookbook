package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup_Flags(t *testing.T) {
	c:=Config{}
	c.Setup()

	assert.EqualValues(t, "",c.subject)
	assert.EqualValues(t, false,c.isAwesome)
	assert.EqualValues(t, 10,c.howAwesome)
}
