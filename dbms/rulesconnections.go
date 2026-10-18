package dbms

import (
	"errors"

	"github.com/kudryashov216/union-orm/dbms/enums"
	"github.com/kudryashov216/union-orm/types"
)

func GetConnectionTemplate(SQLName types.SQL) (map[string]any, error) {

	switch SQLName {

	// case enums.MYSQL:

	// case enums.MARIADB:

	// case enums.POSTGRESQL:

	// case enums.MSSQL:

	case enums.SQLITE:
		rule := make(map[string]any)
		rule["ID"] = SQLName
		rule["PathToDB"] = ""
		return rule, nil
	default:
		return nil, errors.New("This subsystem is not supported in this version of ORM.")

	}

}
