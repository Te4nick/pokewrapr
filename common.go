package pokewrapr

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/Te4nick/pokewrapr/cache"
)

const baseURL string = "https://pokeapi.co/api/v2"

var c *cache.Cache = cache.NewCache()
var CacheTTL = cache.CacheDefaultTTL
var UseCache = true

func fetch(endpoint string, out interface{}) error {
	var apiURL string = ""
	if strings.Contains(endpoint, baseURL) {
		apiURL = endpoint
		endpoint = apiURL[len(baseURL):]
	} else {
		apiURL = baseURL
		sep := ""
		if endpoint[0] != '/' {
			sep = "/"
		}
		apiURL += sep + endpoint
	}

	if UseCache {
		data := c.Get(endpoint)
		if data != nil { // TODO: Check if can set value
			reflect.ValueOf(out).Elem().Set(reflect.ValueOf(data).Elem())
		}
	}

	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return HTTPStatusError{StatusCode: resp.StatusCode, URL: apiURL}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	if UseCache {
		if CacheTTL != 0 {
			c.Set(endpoint, out, CacheTTL)
		} else {
			c.SetDefault(endpoint, out)
		}
	}

	return nil
}

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func convertIdOrName[T Integer | string](id T) (s string) {
	switch v := interface{}(id).(type) {
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case string:
		s = v
	}
	return
}
