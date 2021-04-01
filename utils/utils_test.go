package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateHour(t *testing.T) {
	c := require.New(t)

	c.False(ValidateHour("30:20"))
	c.False(ValidateHour("25:20"))
	c.False(ValidateHour("not-an-hour"))
	c.True(ValidateHour("3:2"))
	c.True(ValidateHour("03:02"))
	c.True(ValidateHour("12:00"))
}
