package db

import "github.com/lib/pq"

// https://www.postgresql.org/docs/9.4/static/errcodes-appendix.html
func processError(err error) error {
	pqErr, ok := err.(*pq.Error)
	if !ok {
		return err
	}

	switch pqErr.Code {
	case "23505":
		return ErrPgUniqueViolation
	case "23502":
		return ErrPgNotNullViolation
	case "23503":
		return ErrPgForeignKeyViolation
	}

	return err
}
