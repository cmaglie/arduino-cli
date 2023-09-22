// This file is part of arduino-cli.
//
// Copyright 2023 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package f_test

import (
	"strings"
	"testing"

	f "github.com/arduino/arduino-cli/internal/algorithms"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	a := []string{"aaa", "bbb", "ccc"}
	require.Equal(t, []string{"bbb", "ccc"}, f.Filter(a, func(x string) bool { return x > "b" }))
	b := []int{5, 9, 15, 2, 4, -2}
	require.Equal(t, []int{5, 9, 15}, f.Filter(b, func(x int) bool { return x > 4 }))
}

func TestIsEmpty(t *testing.T) {
	require.True(t, f.Equals(int(0))(0))
	require.False(t, f.Equals(int(1))(0))
	require.True(t, f.Equals("")(""))
	require.False(t, f.Equals("abc")(""))
	require.False(t, f.NotEquals(int(0))(0))
	require.True(t, f.NotEquals(int(1))(0))
	require.False(t, f.NotEquals("")(""))
	require.True(t, f.NotEquals("abc")(""))
}

func TestMap(t *testing.T) {
	value := "hello, world , how are,you? "
	parts := f.Map(strings.Split(value, ","), strings.TrimSpace)

	require.Equal(t, 4, len(parts))
	require.Equal(t, "hello", parts[0])
	require.Equal(t, "world", parts[1])
	require.Equal(t, "how are", parts[2])
	require.Equal(t, "you?", parts[3])

	var n []string
	require.Nil(t, f.Map(n, strings.TrimSpace)) // nil input -> nil output
	n = []string{}
	require.Nil(t, f.Map(n, strings.TrimSpace)) // empty input -> nil output
}

func TestReduce(t *testing.T) {
	sum := func(x, y int) int { return x + y }
	one := func(x, y int) int { return 1 }
	require.Equal(t, 15, f.Reduce([]int{1, 2, 3, 4, 5}, sum))
	require.Equal(t, 1, f.Reduce([]int{1, 2, 3, 4, 5}, one))
	// empty set -> just return the zero value
	require.Equal(t, 0, f.Reduce([]int{}, sum))
	require.Equal(t, 0, f.Reduce([]int{}, one))
	// single element -> just return the single element (do not call reduce function)
	require.Equal(t, 10, f.Reduce([]int{10}, sum))
	require.Equal(t, 10, f.Reduce([]int{10}, one))
}
