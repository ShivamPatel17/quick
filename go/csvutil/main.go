package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/jszwec/csvutil"
)

type MyStruct struct {
	Name    string `csv:"name"`
	Age     int    `csv:"age"`
	City    string `csv:"city"`
	Enabled bool   `csv:"enabled"`
}

func WriteCSV(csvdata interface{}, w io.Writer) error {
	csvBytes, err := csvutil.Marshal(csvdata)
	if err != nil {
		return fmt.Errorf("error Marshaling csvdata: %w", err)
	}

	_, err = w.Write(csvBytes)
	if err != nil {
		return fmt.Errorf("error writing csvBytes: %w", err)
	}
	return nil
}

func main() {
	data := []*MyStruct{
		{Name: "Alice", Age: 30, City: "New York", Enabled: true},
		{Name: "Bob", Age: 25, City: "London", Enabled: false},
		{Name: "Charlie", Age: 35, City: "Tokyo", Enabled: true},
	}
	data = append(data, &MyStruct{
		Name: "Shivam",
	})

	data = append(data, &MyStruct{})

	var buf bytes.Buffer
	err := WriteCSV(data, &buf)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(buf.String())
}
