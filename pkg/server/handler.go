package server

import (
	"bytes"
	"fmt"
	"net/http"
)

const meta = `
<meta http-equiv="refresh" content="2" />
`

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(s.render())
	return
}

func (s *Server) render() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(meta)
	for _, log := range s.logs {
		buffer.WriteString(fmt.Sprintf("<div>%s</div>", log))
	}

	return buffer.Bytes()
}
