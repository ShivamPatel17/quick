package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	val *int64
}
type Msg struct {
	idPtr *int64
	id    int64
}

func main() {

	d := &data{
		val: nil,
	}

	msg, err := json.Marshal(&Msg{
		idPtr: nil,
		id:    *d.val,
	})

	if err != nil {
		fmt.Printf("Error marshalling message: %v\n", err)
	}

	fmt.Printf("Marshalled message: %s\n", string(msg))
}
