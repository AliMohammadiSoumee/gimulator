package simulator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatch(t *testing.T) {
	testCases := []struct {
		left  interface{}
		right interface{}
		match bool
	}{
		{
			left:  "foo",
			right: "foo",
			match: true,
		},
		{
			left:  "foo",
			right: "bar",
			match: false,
		},
		{
			left:  1,
			right: 2,
			match: false,
		},
		{
			left:  interface{}("foo"),
			right: "foo",
			match: true,
		},
		{
			left:  []string{"a", "b"},
			right: []string{"a", "b"},
			match: true,
		},
		{
			left:  5,
			right: 5,
			match: true,
		},
		{
			left:  map[string]string{"a": "b"},
			right: map[string]string{"a": "b", "b": "c"},
			match: true,
		},
		{
			left:  map[string]string{"a": "b"},
			right: map[string]string{},
			match: false,
		},
		{
			left:  interface{}(map[string]string{"a": "b"}),
			right: map[string]string{},
			match: false,
		},
		{
			left:  map[string]string{"a": "b"},
			right: interface{}(map[string]string{}),
			match: false,
		},
		{
			left:  interface{}(map[string]string{"a": "b"}),
			right: interface{}(map[string]string{}),
			match: false,
		},
		{
			left:  nil,
			right: 5,
			match: true,
		},
		{
			left:  map[string]string{"a": "b"},
			right: map[string]interface{}{"a": "b"},
			match: true,
		},
		{
			left:  interface{}(map[string]string{"a": "b"}),
			right: map[string]interface{}{"a": "b"},
			match: true,
		},
		{
			left:  []interface{}{"best", "size", 85},
			right: []interface{}{"best", "size", 85},
			match: true,
		},
		{
			left:  []interface{}{"best", "size", 85},
			right: interface{}([]interface{}{"best", "size", 85}),
			match: true,
		},
		{
			left:  interface{}([]interface{}{"best", "size", 85}),
			right: []interface{}{"best", "size", 85},
			match: true,
		},
		{
			left:  interface{}([]interface{}{"best", "size", 85}),
			right: []interface{}{"best", "size", 75},
			match: false,
		},
		{
			left:  interface{}(5),
			right: 5,
			match: true,
		},
		{
			left:  interface{}(5),
			right: 6,
			match: false,
		},
		{
			left:  map[interface{}]interface{}{"foo": 5},
			right: map[interface{}]interface{}{"foo": 5, "bar": 6},
			match: true,
		},
		{
			left:  map[interface{}]interface{}{"foo": 5},
			right: map[interface{}]interface{}{"foo": 6, "bar": 6},
			match: false,
		},
		{
			left:  map[interface{}]int{"foo": 5},
			right: map[interface{}]interface{}{"foo": 5, "bar": 6},
			match: true,
		},
	}
	for _, tc := range testCases {
		require.Equal(t, match(tc.left, tc.right), tc.match)
	}
}
