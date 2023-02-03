package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os/exec"
)

// getCommandOutput is a wrapper function for the exec.Command function
func getCommandOutput(cmd string, args ...string) string {
	out, _ := exec.Command(cmd, args...).Output()
	return string(out)
}

// GetGoVersion is a handler that returns the specified go binary's version
func GetGoVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := getCommandOutput("/usr/bin/go", "version")
	io.WriteString(w, res)
	return
}

// GetFileContent is a handler that receives a filename and returns the text contained within
func GetFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, getCommandOutput("/bin/cat", params.ByName("name")))
	return
}
