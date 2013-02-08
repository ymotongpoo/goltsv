package ltsv

import (
	"bufio"
	"errors"
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
	reader *bufio.Reader
}

// NewReader returns a new LTSVReader that reads from r.
func NewReader(r io.Reader) *LTSVReader {
	return &LTSVReader{bufio.NewReader(r)}
}

// error creates a new Error based on err.
// TODO(ymotongpoo): enhance parse error
func (r *LTSVReader) error(err error) error {
	return err
}

// Read reads one record from r. The record is a map of string with
// each key and value representing one field.
func (r *LTSVReader) Read() (record map[string]string, err error) {
	var line []byte
	record = make(map[string]string)

	for {
		line, _, err = r.reader.ReadLine()
		if err != nil {
			return nil, err
		}

		sline := strings.TrimSpace(string(line))
		if sline == "" {
			// Skip empty line
			continue
		}
		tokens := strings.Split(sline, "\t")
		if len(tokens) == 0 {
			return nil, r.error(ErrFieldFormat)
		}
		for _, field := range tokens {
			if field == "" {
				continue
			}
			data := strings.SplitN(field, ":", 2)
			if len(data) != 2 {
				return record, r.error(ErrLabelName)
			}
			record[data[0]] = data[1]
		}
		return record, nil
	}
	return
}

// ReadAll reads all the remainig records from r.
// Each records is a slice of map of fields.
// TODO(ymotongpoo): compare with the case of using csv.ReadAll()
func (r *LTSVReader) ReadAll() (records []map[string]string, err error) {
	for {
		record, err := r.Read()
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
