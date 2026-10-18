package drivers

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5"
	"github.com/kudryashov216/union-orm/dbms/enums"
	"github.com/kudryashov216/union-orm/types"
	_ "github.com/microsoft/go-mssqldb"
	_ "modernc.org/sqlite"
)

func GetDriverSQL(SQLName types.SQL) (types.SQL, error) {

	switch SQLName {
	case enums.MYSQL:
		return "mysql", nil
	case enums.MARIADB:
		return "mysql", nil
	case enums.POSTGRESQL:
		return "pgx", nil
	case enums.MSSQL:
		return "go-mssqldb", nil
	case enums.SQLITE:
		return "sqlite", nil
	default:
		return "", errors.New("This subsystem is not supported in this version of ORM.")
	}

}
