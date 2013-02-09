========
 goltsv
========

LTSV (Labeled Tab Separeted Value) reader/writer for Go.

.. image:: https://drone.io/github.com/ymotongpoo/goltsv/status.png

Example
=======

Reader
------

::

   package main
   
   import (
   	"bytes"
   	"fmt"
   	"github.com/ymotongpoo/goltsv"
   )
      
   func main() {
   	data := `
   egg:たまご	ham:ハム	bread:パン
   gyoza:ぎょうざ	ramen:ラーメン	sushi:すし	yakiniku:焼肉
   yamanashi:山梨	tokyo:東京	okinawa:沖縄	hokkaido:北海道
   `
   	b := bytes.NewBufferString(data)
      
   	// Read LTSV file into map[string]string
   	reader := goltsv.NewReader(b)
   	records, err := reader.ReadAll()
   	if err != nil {
   		panic(err)
   	}
      
   	// dump
   	for i, record := range records {
   		fmt.Printf("===== Data %d\n", i)
   		for k, v := range record {
   			fmt.Printf("\t%s --> %s\n", k, v)
   		}
   	}
   }
   

Writer
------

::

   package main
   
   import (
   	"bytes"
   	"github.com/ymotongpoo/goltsv"
   )
   
   func main() {
   	data := []map[string]string {
   		{"Python": "3.3.0", "Ruby": "2.0 rc2", "Perl": "5.16.2"},
   		{"spam": "foo", "egg": "bar", "ham": "buz"},
   		{"sauce": "ソース", "salt": "しお", "sugar": "さとう", "vinegar": "す"},
   	}
   
   	b := &bytes.Buffer{}
   	writer := goltsv.NewWriter(b)
   	err := writer.WriteAll(data)
   	if err != nil {
   		panic(err)
   	}
   	fmt.Printf("%v", b.String())
   }


License
=======

This packages is distributed under conditions of New BSD License

