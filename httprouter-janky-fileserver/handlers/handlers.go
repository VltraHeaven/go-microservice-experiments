package handlers

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os/exec"
)

func getCommandOutput(cmd string, args ...string) string {
	out, _ := exec.Command(cmd, args...).Output()
	return string(out)
}

func GetGoVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := getCommandOutput("/usr/bin/go", "version")
	io.WriteString(w, res)
	return
}

func GetFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	io.WriteString(w, getCommandOutput("/bin/cat", params.ByName("name")))
	return
}
