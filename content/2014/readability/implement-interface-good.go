// +build OMIT

package sample // OMIT

import "scan" // OMIT

// ColumnWriter is a writer to write ...
type ColumnWriter struct {
	tmpDir string
	// some other fields
}

var _ scan.Writer = (*ColumnWriter)(nil) // HL
