package db

// Thanks for Sudip Raval https://medium.com/@rsudip90
import (
	"database/sql"
)

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

// NullBool is an alias for sql.NullBool data type
type NullBool sql.NullBool

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 sql.NullFloat64

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

// NullTime is an alias for mysql.NullTime data type
type NullTime sql.NullTime
