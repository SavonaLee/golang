package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func inner() error {
	return errors.Wrap(sql.ErrNoRows, "inner error")
}

func outer() error {
	if err := inner(); err != nil {
		return errors.WithMessage(err, "outer error")
	}
	return nil
}

func main() {
	err := outer()
	switch {
	case errors.Cause(err) == sql.ErrNoRows:
		fmt.Printf("data not found, %v\n", err)
		fmt.Printf("%+v\n", err)
	case err != nil:
		fmt.Println("Unkown error", err)
	}
}
