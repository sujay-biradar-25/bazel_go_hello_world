package internal_deps

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func InternalFunction() {
	fmt.Println("Hello from InternalFunction!")
	snowflakeNode, err := snowflake.NewNode(1000)
	if err != nil {
		fmt.Println("Error creating snowflake node:", err)
		return
	}
	id := snowflakeNode.Generate()
	fmt.Println(id)
	fmt.Println("Hello from InternalFunction!")

}
