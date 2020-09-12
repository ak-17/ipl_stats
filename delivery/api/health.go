package api

import "net/http"

func (api *Api) Health(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "OK", nil
}
