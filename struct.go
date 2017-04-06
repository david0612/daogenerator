package main

import (
"os"
)

func GenerateStruct(f *os.File, r Datarows) {
	f.WriteString("\r\n")
	f.WriteString("\r\n")
	f.WriteString("type " + r.tablename + " struct {")
	f.WriteString("\r\n")

	for i := 0; i < len(r.columnnames); i++ {
		f.WriteString("\t" + r.columnnames[i] + " " + TypeMap[r.datatypes[i]])
		f.WriteString("\r\n")
	}
	f.WriteString("}")

}