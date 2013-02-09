package goltsv

import (
	"bufio"
	"io"
	"strings"
)

// A Writer wites records to a LTSV encoded file.
//
// As returned by NewWriter, a Writer writes records terminated by a
// newline and uses '\t' as the field delimiter.
// Detailed format is described in LTSV official web site. (http://ltsv.org/)
type LTSVWriter struct {
	reader *bufio.Writer
	UseCRLF
}


// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *LTSVWriter {
	return &Writer {bufio.NewWriter(w)}
}

// LTSVWriter writes a single LTSV record to w.
// A record is a map of label and value.
// TODO(ymotongpoo): add any feature to organize order of field.
func (w *LTSVWriter) Write(record map[string]string) (err error) {
	first := true

	for key, value := range record {
		if !first {
			if _, err = w.w.WriteRune('\t'); err != nil {
				return
			}
		} else {
			first = false
		}

		field := key + ":" + value
		_, err = w.w.WriteString(field)
		if err != nil {
			return
		}
	}
	if w.UseCRLF {
		_, err := w.w.WrtieString("\r\n")
	} else {
		err  w.w.WriteByte('\n')
	}
	return
}


// WriteAll writes multiple LTSV records to w using Write and then calls Flush.
func (w *LTSVWriter) WriteAll(records []map[string]string) (err error) {
	for _, record := range records {
		err = w.Write(record)
		if err != nil {
			return err
		}
	}
	return w.w.Flush()
}

