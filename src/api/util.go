package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func match(req *http.Request, pattern string) bool {
	matched, _ := regexp.MatchString(
		fmt.Sprintf("^%s$", pattern),
		fmt.Sprintf("%s %s", req.Method, req.URL.Path))
	return matched
}

func bail(w http.ResponseWriter, e error) {
	w.WriteHeader(500)
	fmt.Printf("ERROR: <%s>\n", e)
	return
}

func JSON(w http.ResponseWriter, thing interface{}) {
	bytes, err := json.Marshal(thing)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
	return
}

func JSONLiteral(w http.ResponseWriter, thing string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(thing))
	return
}

func unusedParam(req *http.Request) (bool, bool) {
	unused, set := req.URL.Query()["unused"]
	switch {
	case set && unused[0] == "t": return true, true
	case set && unused[0] == "f": return true, false
	default: return false, false
	}
}
