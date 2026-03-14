package json

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/arashn0uri/go-server/internal/models"
)

func Write(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func Read(r *http.Request, dest any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dest)
}

func WriteError(w http.ResponseWriter, status int, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error: err,
	})
}

// ReadMultipart decodes a multipart/form-data request into dest using struct tags.
//
// Supported tags:
//
//	form:"fieldName"        -> r.FormValue("fieldName") for string, int, float, bool
//	form:"fieldName,file"   -> r.FormFile, field must be multipart.File
//	form:"fieldName,header" -> r.FormFile header, field must be *multipart.FileHeader
func ReadMultipart(r *http.Request, dest any) error {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		return fmt.Errorf("parse multipart form: %w", err)
	}

	type fileResult struct {
		file   multipart.File
		header *multipart.FileHeader
		err    error
	}
	fileCache := map[string]*fileResult{}
	getFile := func(key string) *fileResult {
		if cached, ok := fileCache[key]; ok {
			return cached
		}
		f, h, err := r.FormFile(key)
		result := &fileResult{f, h, err}
		fileCache[key] = result
		return result
	}

	v := reflect.ValueOf(dest).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("form")
		if tag == "" {
			continue
		}

		parts := strings.SplitN(tag, ",", 2)
		key, modifier := parts[0], ""
		if len(parts) == 2 {
			modifier = parts[1]
		}

		fieldVal := v.Field(i)

		switch modifier {
		case "file":
			res := getFile(key)
			if res.err != nil {
				return fmt.Errorf("field %s: %w", key, res.err)
			}
			fieldVal.Set(reflect.ValueOf(res.file))
		case "header":
			res := getFile(key)
			if res.err != nil {
				return fmt.Errorf("field %s header: %w", key, res.err)
			}
			fieldVal.Set(reflect.ValueOf(res.header))
		default:
			if err := setField(fieldVal, r.FormValue(key)); err != nil {
				return fmt.Errorf("field %s: %w", key, err)
			}
		}
	}

	return nil
}

func setField(v reflect.Value, raw string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(raw)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer: %w", err)
		}
		v.SetInt(n)
	case reflect.Float32, reflect.Float64:
		n, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return fmt.Errorf("invalid float: %w", err)
		}
		v.SetFloat(n)
	case reflect.Bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return fmt.Errorf("invalid bool: %w", err)
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported type: %s", v.Kind())
	}
	return nil
}
