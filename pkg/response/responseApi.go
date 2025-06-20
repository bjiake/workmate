package responseApi

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"workmate/internal/domain/models/response"
)

func WriteJson(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, response.Success{
		Data: data,
	})
}

func WriteError(w http.ResponseWriter, r *http.Request, status int, err interface{}) {
	render.Status(r, status)

	var result map[string]interface{}

	// for map validator errors
	if errMap, ok := err.(map[string]string); ok {
		result = map[string]interface{}{
			"error": errMap,
		}
	} else {
		// for simple errors
		result = map[string]interface{}{
			"error": fmt.Sprintf("%v", err),
		}
	}

	render.JSON(w, r, result)
}
