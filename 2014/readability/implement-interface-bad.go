// +build OMIT

package sample // OMIT

import "scan" // OMIT

// Column writer implements the scan.Writer interface.
type ColumnWriter struct {
	scan.Writer // HL
	tmpDir      string
	// some other fields
}
