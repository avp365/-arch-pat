package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	result, _ := ReverseInt(123)
	require.Equal(t, 321, result, "fail work func 123")

	result, _ = ReverseInt(-649)
	require.Equal(t, -946, result, "fail work func -946")

	result, _ = ReverseInt(0)
	require.Equal(t, 0, result, "fail work func 0")

	result, _ = ReverseInt("a")
	require.Equal(t, 0, result, "fail work func a")

	result, _ = ReverseInt(0.1)
	require.Equal(t, 0, result, "fail work func 0.1")

}
