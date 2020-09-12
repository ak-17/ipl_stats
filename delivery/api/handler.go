package api

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type HandlerFunc func(rw http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := Response{}
	start := time.Now()

	data, err := fn(w, r)
	response.Header.ProcessTime = time.Since(start).String()
	w.Header().Set("Content-Type", "application/json")
	if data != nil && err == nil {
		response.Data = data
		if buf, err := json.Marshal(response); err == nil {
			_, err := w.Write(buf)
			if err != nil {
				log.Errorf("[API] [ServeHTTP] [Error] in writing http response. err: %s", err.Error())
				return
			}
		}
	}

	if err != nil {
		response.Header.ErrorMessage = []string{err.Error()}
		log.Errorf("[API] [ServeHTTP] [Error] handler error. err: %v", err)
	} else {
		return
	}

	if buf, err := json.Marshal(response); err == nil {
		_, err := w.Write(buf)
		if err != nil {
			log.Errorf("[API] [ServeHTTP] [Error] in writing http response. err: %s", err.Error())
			return
		}
	}

}
