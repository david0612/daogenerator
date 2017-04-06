package main

import (
"os"
)

func GenerateHeader(f *os.File) {
	f.WriteString("package main")
	f.WriteString("\r\n\r\n")
	f.WriteString("import (")
	f.WriteString("\r\n")
	f.WriteString("\t\"database/sql\"")
	f.WriteString("\r\n")
	f.WriteString("\t_ \"github.com/lib/pq\"")
	f.WriteString("\r\n")
	f.WriteString("\t\"bytes\"")
	f.WriteString("\r\n")
	f.WriteString("\t\"strconv\"")
	f.WriteString("\r\n")
	f.WriteString(")")
}