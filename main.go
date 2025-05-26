package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello from Bazel Go Hello World!")

	// Generate a snowflake ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}

	id := node.Generate()
	fmt.Printf("Generated Snowflake ID: %d\n", id)

	// Generate a UUID
	newUUID := uuid.New()
	fmt.Printf("Generated UUID: %s\n", newUUID.String())

	fmt.Println("Both dependencies working successfully!")
}
