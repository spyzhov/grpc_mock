package manager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type HttpRequest struct {
	Request  json.RawMessage `json:"request"`
	Response json.RawMessage `json:"response"`
	Error    Error           `json:"error"`
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	var (
		uerr error
		err  error
		data []byte
	)
	defer func(start time.Time) {
		log.Printf("HTTP request %s for %s lasted for %s", request.Method, request.RequestURI, time.Since(start))
	}(time.Now())
	defer func() {
		if request.Body != nil {
			_ = request.Body.Close()
		}
	}()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error: %v", r)
			if err != nil {
				_, _ = writer.Write([]byte(err.Error()))
				writer.WriteHeader(http.StatusInternalServerError)
			}
		}
	}()
	switch request.Method {
	case http.MethodDelete:
		if request.RequestURI != "/" {
			uerr = fmt.Errorf("partitial clean of mock is not supported")
		} else {
			Reset()
		}
	case http.MethodPost:
		parts := strings.Split(strings.Trim(request.RequestURI, "/"), "/")
		if len(parts) != 2 {
			uerr = fmt.Errorf("URI should be like: package.Service/MethodName")
			break
		}
		method := parts[1]
		parts = strings.Split(parts[0], ".")
		if len(parts) < 2 {
			uerr = fmt.Errorf("URI should be like: package.Service/MethodName")
			break
		}
		service := parts[len(parts)-1]
		pkg := strings.Join(parts[:len(parts)-1], ".")
		data, uerr = ioutil.ReadAll(request.Body)
		if uerr != nil {
			break
		}
		req := new(HttpRequest)
		if uerr = json.Unmarshal(data, &req); uerr != nil {
			break
		}
		Set(
			PackageName(pkg),
			ServiceName(service),
			MethodName(method),
			Request(req.Request),
			Response(req.Response),
			req.Error)
	case http.MethodGet:
		// todo: not implemented
	}
	if err != nil {
		_, _ = writer.Write([]byte(err.Error()))
		writer.WriteHeader(http.StatusInternalServerError)
	} else if uerr != nil {
		_, _ = writer.Write([]byte(uerr.Error()))
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		writer.WriteHeader(http.StatusOK)
	}
}
