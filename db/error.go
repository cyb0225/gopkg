package db

import "github.com/go-sql-driver/mysql"

const (
	ErrMySQLDupEntry            = 1062
	ErrMySQLDupEntryWithKeyName = 1586
)

// IsDuplicateError 是否是主键冲突错误
func IsDuplicateError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == ErrMySQLDupEntry ||
			mysqlErr.Number == ErrMySQLDupEntryWithKeyName {
			return true
		}
	}
	return false
}
