package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync"
)

type (
	PackageName string
	ServiceName string
	MethodName  string
	Request     []byte
	Response    []byte
	Error       string
	RequestHash interface{}
	Mock        struct {
		Hash     RequestHash
		Request  Request
		Response Response
		Error    Error
		Called   int
	}
	Mocks    []*Mock
	Stub     map[MethodName]Mocks
	Stubs    map[ServiceName]Stub
	Packages map[PackageName]Stubs
	mock     struct {
		mu       sync.Mutex
		Packages Packages
	}
)

var mocks = &mock{Packages: make(Packages)}

func (m *mock) get(pkg PackageName, service ServiceName, method MethodName, hash RequestHash) *Mock {
	if packages, ok := m.Packages[pkg]; !ok {
		return nil
	} else if stubs, ok := packages[service]; !ok {
		return nil
	} else if mocks, ok := stubs[method]; !ok {
		return nil
	} else {
		for _, mock := range mocks {
			if isEqualHash(mock.Hash, hash) {
				return mock
			}
		}
	}
	return nil
}

func Reset() {
	mocks.mu.Lock()
	defer mocks.mu.Unlock()
	mocks.Packages = make(Packages, 0)
}

func Set(pkg PackageName, service ServiceName, method MethodName, request Request, response Response, error Error) {
	mocks.mu.Lock()
	defer mocks.mu.Unlock()
	hash := request.Hash()
	if mock := mocks.get(pkg, service, method, hash); mock != nil {
		mock.Response = response
		mock.Error = error
		mock.Called = 0
	} else {
		if _, ok := mocks.Packages[pkg]; !ok {
			mocks.Packages[pkg] = make(Stubs, 0)
		}
		if _, ok := mocks.Packages[pkg][service]; !ok {
			mocks.Packages[pkg][service] = make(Stub, 0)
		}
		if _, ok := mocks.Packages[pkg][service][method]; !ok {
			mocks.Packages[pkg][service][method] = make(Mocks, 0)
		}
		mocks.Packages[pkg][service][method] = append(mocks.Packages[pkg][service][method], &Mock{
			Hash:     hash,
			Request:  request,
			Response: response,
			Error:    error,
			Called:   0,
		})
	}
}

//func Get(pkg PackageName, service ServiceName, method MethodName, request Request) *Mock {
//	mocks.mu.Lock()
//	defer mocks.mu.Unlock()
//	return mocks.get(pkg, service, method, request.Hash())
//}

func Call(pck PackageName, service ServiceName, method MethodName, in interface{}, out interface{}) (ret interface{}, err error) {
	mocks.mu.Lock()
	defer mocks.mu.Unlock()
	defer func() {
		if r := recover(); r != nil {
			ret = nil
			err = fmt.Errorf("%v", r)
		}
	}()

	var request []byte
	request, err = json.Marshal(in)
	if err != nil {
		log.Panicf("request couldn't be marshaled, for \"%s.%s/%s\": %s", pck, service, method, err)
	}
	mock := mocks.get(pck, service, method, Request(request).Hash())
	mock.Used()
	if mock == nil {
		log.Panicf("unregistered request for \"%s.%s/%s\" with the request: %s", pck, service, method, string(request))
	}
	if mock.Error != "" {
		return out, fmt.Errorf("%s", mock.Error)
	}
	if err = json.Unmarshal(mock.Response, &out); err != nil {
		log.Panicf("response couldn't be marshaled, for \"%s.%s/%s\": %s", pck, service, method, err)
	}
	return out, nil
}

func (r Request) Hash() RequestHash {
	data := new(RequestHash)
	if err := json.Unmarshal(r, &data); err != nil {
		log.Panicf("error on parse request: %s\nInput:\n%s", err, string(r))
	}
	return *data
}

func isEqualHash(left, right RequestHash) bool {
	return reflect.DeepEqual(left, right)
}

func (m *Mock) Used() {
	if m == nil {
		return
	}
	m.Called++
}
