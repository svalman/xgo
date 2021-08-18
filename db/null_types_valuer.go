package db

import (
	"database/sql/driver"
)

// Value implements the Scanner interface for NullInt64
func (ni NullInt64) Value() (driver.Value, error) {

	if ni.Valid {
		return ni.Int64, nil
	} else {
		return 0, nil
	}

}

func (nb NullBool) Value() (driver.Value, error) {

	if nb.Valid {
		return nb.Bool, nil
	} else {
		return false, nil
	}

}

func (nf NullFloat64) Value() (driver.Value, error) {

	if nf.Valid {
		return nf.Float64, nil
	} else {
		return 0, nil
	}

}

func (ns NullString) Value() (driver.Value, error) {

	if ns.Valid {
		return ns.String, nil
	} else {
		return "", nil
	}

}

func (nt NullTime) Value() (driver.Value, error) {

	if nt.Valid {
		return nt.Time, nil
	} else {
		return 0, nil
	}

}
