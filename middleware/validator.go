package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type validateErr map[string]string

func BodyValidator[T any](handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := new(T)
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode([]validateErr{{"key": "body",
				"value": "Invalid Body"}})
			return
		}
		validateErrs := validator.New().Struct(body)
		if validateErrs != nil {
			errors := validateErrs.Error()
			w.WriteHeader(http.StatusBadRequest)
			errorArr := strings.Split(errors, "Key: ")
			errorsObj := []validateErr{}
			for _, element := range errorArr {
				if strings.TrimSpace(element) == "" {
					continue
				}
				elementArr := strings.Split(element, "Error:")
				newError := validateErr{"key": strings.TrimSpace(elementArr[0]), "value": strings.ReplaceAll(elementArr[len(elementArr)-1], "\n", "")}
				errorsObj = append(errorsObj, newError)
			}
			json.NewEncoder(w).Encode(errorsObj)
			return
		}
		ctx := context.WithValue(r.Context(), "body", body)
		req := r.WithContext(ctx)
		handler(w, req)
	}
}
