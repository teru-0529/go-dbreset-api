/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package businesslogic

import (
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
)

// FUNCTION:
func execTruncate(schema string, table string) error {
	sql := fmt.Sprintf("TRUNCATE %s.%s RESTART IDENTITY CASCADE;", schema, table)
	_, err := queries.Raw(sql).Exec(boil.GetDB())
	return err
}

// FUNCTION:
func execSeqReset(schema string, sequence string) error {
	sql := fmt.Sprintf("SELECT SETVAL ('%s.%s', 1, false);", schema, sequence)
	_, err := queries.Raw(sql).Exec(boil.GetDB())
	return err
}
