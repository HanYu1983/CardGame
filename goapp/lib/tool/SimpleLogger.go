package tool

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type SimpleLogger struct {
	Writer http.ResponseWriter
}

func (l *SimpleLogger) Log(msg interface{}) {
	str, _ := json.Marshal(msg)
	fmt.Fprintf(l.Writer, "%s\n", str)
}