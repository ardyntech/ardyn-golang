package responses

//-------------------------------------------------------------

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

//-------------------------------------------------------------
