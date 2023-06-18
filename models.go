package openai

import (
	"fmt"
	"net/http"
)

const modelsSuffix = "/models"


// {      
// 	"id": "gpt-3.5-turbo-16k",
//     "object": "model",
//     "created": 1683758102,
//       "owned_by": "openai-internal",
//       "permission": [
//         {
//           "id": "modelperm-LMK1z45vFJF9tVUvKb3pZfMG",
//           "object": "model_permission",
//           "created": 1686799823,
//           "allow_create_engine": false,
//           "allow_sampling": true,
//           "allow_logprobs": true,
//           "allow_search_indices": false,
//           "allow_view": true,
//           "allow_fine_tuning": false,
//           "organization": "*",
//           "group": null,
//           "is_blocking": false
//         }
//       ],
//       "root": "gpt-3.5-turbo-16k",
//       "parent": null
// }


type ModelsResponse struct {

}

//unfinish
func (c *Client) ListModels() (err error) {
	var response string

	url := c.Config.apiBase + modelsSuffix

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	err = c.request(req, &response)
	fmt.Printf("response: %v\n", response)
	return
}
