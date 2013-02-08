package ltsv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strings"
)

// These are the errors that can be returned in ParseError.Error
var (
	ErrFieldFormat = errors.New("wrong LTSV field format")
	ErrLabelName   = errors.New("unexpected label name")
)

// A Reader reads from a Labeled TSV (LTSV) file.
//
// As returned by NewReader, a Reader expects input conforming LTSV (http://ltsv.org/)
type LTSVReader struct {
	Reader      *csv.Reader
	FieldLabels map[string]struct{}
	initialized bool
}

// NewReader returns a new LTSVReader that reads from r.
func NewReader(r io.Reader) *LTSVReader {
	reader := csv.NewReader(r)
	reader.Comma = '\t'
	return &LTSVReader{
		Reader:      reader,
		FieldLabels: make(map[string]struct{}),
	}
}

// error creates a new Error based on err.
// TODO(ymotongpoo): enhance parse error
func (r *LTSVReader) error(err error) error {
	return err
}

// Read reads one record from r. The record is a map of string with
// each key and value representing one field.
func (r *LTSVReader) Read() (record map[string]string, err error) {
	rawRecord, err := r.Reader.Read()
	if err != nil {
		return nil, err
	}

	record = make(map[string]string)
	for _, field := range rawRecord {
		data := strings.SplitN(field, ":", 2)
		if len(data) != 2 {
			return record, r.error(ErrFieldFormat)
		}
		if r.initialized {
			if _, ok := r.FieldLabels[data[0]]; ok {
				record[data[0]] = data[1]
			} else {
				return nil, r.error(ErrLabelName)
			}
		} else {
			record[data[0]] = data[1]
		}
	}

	if !r.initialized {
		for label, _ := range record {
			r.FieldLabels[label] = struct{}{}
		}
		r.initialized = true
	}

	return record, nil
}

// ReadAll reads all the remainig records from r.
// Each records is a slice of map of fields.
// TODO(ymotongpoo): compare with the case of using csv.ReadAll()
func (r *LTSVReader) ReadAll() (records []map[string]string, err error) {
	for {
		record, err := r.Read()
		fmt.Printf("record: %v\n", record)
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	panic("unreachable")
}
