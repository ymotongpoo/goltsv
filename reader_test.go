package goltsv

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

var readTests = []struct {
	str     string
	records interface{}
}{
	{
		`
hoge:foo	bar:baz
perl:5.17.8	ruby:2.0	python:2.6
sushi:寿司	tennpura:天ぷら	ramen:ラーメン	gyoza:餃子
		`,
		[]map[string]string{
			{"hoge": "foo", "bar": "baz"},
			{"perl": "5.17.8", "ruby": "2.0", "python": "2.6"},
			{"sushi": "寿司", "tennpura": "天ぷら", "ramen": "ラーメン", "gyoza": "餃子"},
		},
	},
}

func TestRead(t *testing.T) {
	n := 0
	reader := NewReader(bytes.NewBufferString(readTests[0].str))
	for {
		_, e := reader.Read()
		if e != nil {
			if e != io.EOF {
				t.Errorf("expected Read got %v", e)
			}
			break
		}
		n++
	}
	if n != 3 {
		t.Errorf("expected 3 records got %v", n)
	}
}

func TestReadAll(t *testing.T) {
	for _, w := range readTests {
		reader := NewReader(bytes.NewBufferString(w.str))
		records, e := reader.ReadAll()
		if e != nil {
			t.Errorf("expected Read got %v", e)
			continue
		}
		if reflect.DeepEqual(records, w.records) == false {
			t.Errorf("expected %v but %v", w.records, records)
			continue
		}
	}
}
