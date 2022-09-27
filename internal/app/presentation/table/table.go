package table

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// table is a struct that represents a table
type table struct {
	headers []string
	rows    [][]string
}

// addRow adds a row to the table
func (t *table) addRow(row []string) {
	t.rows = append(t.rows, row)
}

// render renders the table to the console
func (t *table) render() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, strings.Join(t.headers, "\t"))
	for _, row := range t.rows {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}
