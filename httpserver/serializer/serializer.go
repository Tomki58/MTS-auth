package serializer

import (
	"encoding/json"
)

type ResponseJSON struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func SerializeResponseJSON(data interface{}) []byte {
	response := ResponseJSON{}

	switch data := data.(type) {
	case error:
		response.Error = data.Error()

	default:
		response.Success = true
		response.Data = data
	}
	jRepr, err := json.Marshal(response)
	if err != nil {
		return nil
	}

	return jRepr
}
