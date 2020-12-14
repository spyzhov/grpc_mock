package manager

import (
	"encoding/json"
	"testing"
)

func Test_compare(t *testing.T) {
	get := func(value string) interface{} {
		if value == "null" {
			return nil
		}
		result := new(interface{})
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			panic(err)
		}
		return *result
	}
	tests := []struct {
		name  string
		left  string
		right string
		want  bool
	}{
		{
			name:  "nil",
			left:  "null",
			right: "null",
			want:  true,
		},
		{
			name:  "nil-value",
			left:  "null",
			right: `""`,
			want:  false,
		},
		{
			name:  "nil-n/a",
			left:  "null",
			right: `"` + notImportant + `"`,
			want:  true,
		},
		{
			name:  "int-n/a",
			left:  "1",
			right: `"` + notImportant + `"`,
			want:  true,
		},
		{
			name:  "slice-n/a",
			left:  "[]",
			right: `"` + notImportant + `"`,
			want:  true,
		},
		{
			name:  "object-n/a",
			left:  "{}",
			right: `"` + notImportant + `"`,
			want:  true,
		},
		{
			name:  "slices",
			left:  `[]`,
			right: `[]`,
			want:  true,
		},
		{
			name:  "slices",
			left:  `[1,2,3]`,
			right: `[1,2,3]`,
			want:  true,
		},
		{
			name:  "slices",
			left:  `[1,"` + notImportant + `",3]`,
			right: `[1,2,"` + notImportant + `"]`,
			want:  true,
		},
		{
			name:  "slices",
			left:  `[1,"` + notImportant + `"]`,
			right: `[1,2,"` + notImportant + `"]`,
			want:  false,
		},
		{
			name:  "slices",
			left:  `[1,2,3]`,
			right: `[1,2,"3"]`,
			want:  false,
		},
		{
			name:  "slices",
			left:  `[]`,
			right: `1`,
			want:  false,
		},
		{
			name:  "objects",
			left:  `{}`,
			right: `{}`,
			want:  true,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar"}`,
			right: `{"foo":"bar"}`,
			want:  true,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar","bar":"` + notImportant + `"}`,
			right: `{"foo":"` + notImportant + `"}`,
			want:  true,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar","bar":"` + notImportant + `"}`,
			right: `{"foo":"bar","bar":[]}`,
			want:  true,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar"}`,
			right: `{"foo":"bar","bar":"` + notImportant + `"}`,
			want:  true,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar"}`,
			right: `{"foo":"foo","bar":"` + notImportant + `"}`,
			want:  false,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar"}`,
			right: `{"foo":"foo","bar":"xxx"}`,
			want:  false,
		},
		{
			name:  "objects",
			left:  `{"foo":"bar"}`,
			right: `1`,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compare(get(tt.left), get(tt.right)); got != tt.want {
				t.Errorf("compare() = %v, want %v\n left: %s\nright: %s", got, tt.want, tt.left, tt.right)
			}
		})
	}
}
