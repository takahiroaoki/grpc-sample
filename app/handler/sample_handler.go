package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/takahiroaoki/go-env/service"
)

type SampleHandler struct {
	sampleService service.SampleService
}

func (h *SampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := h.sampleService.GetUserByUserId("1")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Println("Failed to encode user")
	}
}

func NewSampleHandler(sampleService service.SampleService) *SampleHandler {
	return &SampleHandler{
		sampleService: sampleService,
	}
}
