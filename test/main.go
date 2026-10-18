package main

import (
	"fmt"

	"github.com/kudryashov216/union-orm/dbms"
	"github.com/kudryashov216/union-orm/dbms/enums"
)

func main() {

	template, _ := dbms.GetConnectionTemplate(enums.SQLITE)
	template.PathToDB = "./test.db"

	fmt.Println(template)

	// union.GetConnectionToSQL(enums.SQLITE)

}
