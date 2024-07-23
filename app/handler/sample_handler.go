package handler

import (
	"encoding/json"
	"net/http"

	"github.com/takahiroaoki/go-env/service"
)

type SampleHandler struct {
	sampleService service.SampleServiceInterface
}

func (h *SampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := h.sampleService.GetUserByUserId("1")
	json.NewEncoder(w).Encode(user)
}

func NewSampleHandler(sampleService service.SampleServiceInterface) *SampleHandler {
	return &SampleHandler{
		sampleService: sampleService,
	}
}
