package orm

import (
	"database/sql"
	"fmt"

	"github.com/kudryashov216/union-orm/internal/driverunit/drivers"
	"github.com/kudryashov216/union-orm/types"
)

func GetConnectionToSQL(SQLName types.SQL, rulesConnection any) (*sql.DB, error) {

	driver, err := drivers.GetDriverSQL(SQLName)

	if err != nil {
		return nil, err
	}

	fmt.Println(driver)

	return nil, nil

	// return sql.Open(string(driver), dataSourceName)

}
