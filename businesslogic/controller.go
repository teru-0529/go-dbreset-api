/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package businesslogic

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// type TruncateBody []string

// type TruncateBody struct {
// 	Tables []string `json:"tables"`
// }

type SeqResetBody struct {
	Sequences []string `json:"sequences"`
}

// FUNCTION:
func Truncate(c echo.Context) error {
	schema := c.Param("schema")
	body := []string{}
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	for _, table := range body {
		execTruncate(schema, table)
	}

	return c.NoContent(http.StatusNoContent)
}

// FUNCTION:
func SequenceReset(c echo.Context) error {
	schema := c.Param("schema")
	body := []string{}
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	for _, sequence := range body {
		execSeqReset(schema, sequence)
	}

	return c.NoContent(http.StatusNoContent)
}
