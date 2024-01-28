package helper

import (
	"encoding/json"
	"net/http"
	"pixel8labs-test/backend/internal/constant"
)

// Middleware defines a generic middleware function.
type Middleware func(http.HandlerFunc) http.HandlerFunc

func GenericMiddleware(next func(http.ResponseWriter, *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := next(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Marshal the response data into JSON type.
		mapData := map[string]interface{}{"data": response}
		jsonResponse, err := json.Marshal(mapData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set(constant.HTTPHeaderContentType, "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse) //nolint:all
	}
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth := r.Header.Get(constant.HTTPHeaderAuthorization)

		if auth == constant.DefaultString {
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		} else {
			// Set the Content-Type header to indicate JSON response.
			w.Header().Set(constant.HTTPHeaderContentType, "application/json")

			/*tokenize := strings.Split(auth, " ")
			if len(tokenize) < 2 {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(writeErrorResponse(http.StatusInternalServerError))

				return
			}

			// Create an HTTP client.
			client := &http.Client{}

			// Create a request with a custom header.
			req, err := http.NewRequestWithContext(ctx, "GET", "https://api.github.com/user", nil)
			if err != nil {
				log.Printf("Failed to create request: %v", err)
				return
			}

			// Create request header.
			req.Header.Set(constant.HTTPHeaderAuthorization, tokenize[1])
			req.Header.Set(constant.HTTPHeaderAcceptContent, constant.DefaultAPIGitHubAcceptContentValue)

			// Do HTTP call to service provider
			res, err := client.Do(req)
			if err != nil {
				log.Printf("Failed to call service provider: %v", err)
				return
			}
			defer res.Body.Close()

			// Validate response status code.
			if res.StatusCode != http.StatusOK {
				if res.StatusCode == http.StatusUnauthorized {
					r.Header.Set(constant.HTTPHeaderAuthorization, constant.DefaultString)
				} else {
					log.Printf("Request failed with status code: %v", res.StatusCode)
					return
				}
			} */

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
	}
}

//func writeErrorResponse(statusCode int) []byte {
//	mapSts := map[string]string{
//		"error": errors.New(http.StatusText(statusCode)).Error(),
//	}
//	res, _ := json.Marshal(mapSts)
//
//	return res
//}
