package main

import (
	"log/slog"

	"github.com/madhab452/collection/filter"
)

func main() {
	filter, err := filter.Parse(`id="123"`)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for _, v := range filter.AndStatements {
		slog.Info("field", "=>", v.Field)
		slog.Info("operator", "=>", v.Operator)
		slog.Info("value", "=>", v.Value.(string))
	}
}
