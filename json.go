package w

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ErrNotJSON = errors.New("not json")
)

// ReadJSON reads in a json encoded request body into the given object. If the
// object is not JSON, ErrNotJSON is returned. Other errors may be present from
// read. It is recommended that all errors other than ErrNotJSON are treated as
// Internal Server Errors.
func ReadJSON(r *http.Request, v interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(data, v); err != nil {
		return ErrNotJSON
	}

	return nil
}

// WriteJSON writes a jsonizable object out the the response writer. If the
// object is not json serializable, then WriteJSON will panic because that
// is a programming error.
func WriteJSON(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		panic("error jsonizing your data")
	}

	fmt.Fprintf(w, string(data))
}

// WriteJSONError writes a jsonized error out to the client. This will be
// presented to the client as a json blob with the error code as well as
// a list of errors. This is syntactic sugar on top of WriteJSON. The same
// semantics apply.
func WriteJSONError(w http.ResponseWriter, code int, errors []string) {
	out := map[string]interface{}{
		"code": code,
		"errors": errors,
	}
	w.WriteHeader(code)
	WriteJSON(w, out)
}
