package main

import (
	"log"
)

func main() {
	cassandra := GetCassandraSession()
	query := cassandra.Query("select * from imagerec.matches limit 10")

	for i := range IterRows(*query) {
		log.Print(i)
	}
}

