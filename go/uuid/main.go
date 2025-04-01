package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	unmarshalToJson()
}

func unmarshalToJson() {
	type s struct {
		JobUuid uuid.UUID `json:"jobUuid"`
	}

	o := s{}
	jsonString := `{"jobUuid-bad": "123e4567-e89b-12d3-a456-426614174000"}` // or `{"jobUuid": null}` for null
	err := json.Unmarshal([]byte(jsonString), &o)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(o.JobUuid)
}
func initializationInStructs() {
	type s struct {
		jobUuid *uuid.UUID
	}

	t := s{}

	fmt.Printf("jobUUID:%v", t.jobUuid)

}
