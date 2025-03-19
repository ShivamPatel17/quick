package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	type s struct {
		jobUuid *uuid.UUID
	}

	t := s{}

	fmt.Printf("jobUUID:%v", t.jobUuid)
}
