package goltsv

import (
	"bytes"
	"reflect"
	"testing"
)

// TODO(ymotongpoo): Output remains for testing ordered LTSV
var writeTests = []struct {
	Input   map[string]string
	Output  string
	UseCRLF bool
}{
	{
		Input:  map[string]string{"hoge": "foo"},
		Output: "hoge:foo\n",
	},
	{
		Input:   map[string]string{"hoge": "foo"},
		Output:  "hoge:foo\r\n",
		UseCRLF: true,
	},
	{
		Input: map[string]string{"perl": "5.17.8", "ruby": "2.0", "python": "3.3.0"},
		Output: "perl:5.17.8	ruby:2.0	python:3.3.0\n",
	},
	{
		Input: map[string]string{"sushi": "寿司", "tennpura": "天ぷら", "ramen": "ラーメン", "gyoza": "餃子"},
		Output: "sushi:寿司	tennpura:天ぷら	ramen:ラーメン	gyoza:餃子ushi:寿司	tennpura:天ぷら	ramen:ラーメン	gyoza:餃子\n",
	},
}

// TODO(ymotongpoo): Output remains for testing ordered LTSV
var writeAllTests = []struct {
	Input   []map[string]string
	Output  string
	UseCRLF bool
}{
	{
		Input: []map[string]string{
			{"hoge": "foo", "bar": "baz"},
			{"perl": "5.17.8", "ruby": "2.0", "python": "2.6"},
			{"sushi": "寿司", "tennpura": "天ぷら", "ramen": "ラーメン", "gyoza": "餃子"}},
		Output: `hoge:foo	bar:baz
perl:5.17.8	ruby:2.0	python:2.6
sushi:寿司	tennpura:天ぷら	ramen:ラーメン	gyoza:餃子\n`,
	},
}

func TestWrite(t *testing.T) {
	for n, tt := range writeTests {
		b := &bytes.Buffer{}
		f := NewWriter(b)
		f.UseCRLF = tt.UseCRLF
		err := f.Write(tt.Input)
		err = f.Flush()
		if err != nil {
			t.Errorf("Unexpected error: %s\n", err)
		}

		// Note: In Go, map doesn't guarantee order of elements, so comparing
		// original map and map generated from output string by LTSVReader.
		r := NewReader(b)
		out, err := r.Read()
		if err != nil {
			t.Errorf("Unexpected error: %s\n", err)
		}

		if !reflect.DeepEqual(out, tt.Input) {
			t.Errorf("#%d: out=%q want %q", n, b.String(), tt.Output)
		}
	}
}

func TestWriteAll(t *testing.T) {
	for n, tt := range writeAllTests {
		b := &bytes.Buffer{}
		f := NewWriter(b)
		f.UseCRLF = tt.UseCRLF
		err := f.WriteAll(tt.Input)
		if err != nil {
			t.Errorf("Unexpected error: %s\n", err)
		}

		// Note: In Go, map doesn't guarantee order of elements, so comparing
		// original map and map generated from output string by LTSVReader.
		r := NewReader(b)
		out, err := r.ReadAll()
		if !reflect.DeepEqual(out, tt.Input) {
			t.Errorf("#%d: out=%q want %q", n, b.String(), tt.Output)
		}
	}
}

func TestFlush(t *testing.T) {
	b := &bytes.Buffer{}
	f := NewWriter(b)
	inputs := make([]map[string]string, 0)
	for _, tt := range writeTests {
		f.UseCRLF = tt.UseCRLF
		err := f.Write(tt.Input)
		if err != nil {
			t.Errorf("Unexpected error: %s\n", err)
		}
		inputs = append(inputs, tt.Input)
	}

	err := f.Flush()
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}

	// Note: In Go, map doesn't guarantee order of elements, so comparing
	// original map and map generated from output string by LTSVReader.
	r := NewReader(b)
	out, err := r.ReadAll()
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}

	if !reflect.DeepEqual(out, inputs) {
		t.Errorf("out=%q want %q", b.String(), inputs)
	}
}
