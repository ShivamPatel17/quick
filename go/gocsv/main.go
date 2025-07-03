package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type NotUsed struct {
	Name string
}

type Client struct { // Our example struct, you can use "-" to ignore a field
	Age    string               `csv:"client_age"`
	Id     string               `csv:"client_id"`
	Name   string               `csv:"client_name"`
	Payday NetEarningsEntryDate `csv:"payday2,payday"`
	// NotUsedStruct NotUsed `csv:"-"`
}
type NetEarningsEntryDate struct {
	int64
	time.Time
}

func (d *NetEarningsEntryDate) MarshalCSV() (string, error) {
	fmt.Println("marhsal", d)
	return d.Format("2006-01-02"), nil
}

func (d *NetEarningsEntryDate) UnmarshalCSV(csv string) (err error) {
	fmt.Println("unmarshal", csv)
	fmt.Println(csv == "")
	if csv == "" {
		d.Time = time.Time{}
		return nil
	}
	parsedTime, err := time.Parse("2006-01-02", csv)
	if err != nil {
		return fmt.Errorf("date parse: %w", err)

	}

	d.Time = parsedTime.UTC().Truncate(24 * time.Hour) // Ensure the time is set to midnight UTC

	return nil
}
func main() {
	clientsFile, err := os.OpenFile("clients.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Client{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("\n\nHello", client.Name, "Payday", client.Payday)
	}

	if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	csvContent, err := gocsv.MarshalString(&clients)                      // Get all clients as CSV string
	//err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
	fmt.Println(csvContent) // Display all clients as CSV string

}
