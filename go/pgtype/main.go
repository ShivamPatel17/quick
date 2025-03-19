package main

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	var p = pgtype.UUID{}

	fmt.Println(p.Valid)
}
