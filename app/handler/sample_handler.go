package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/takahiroaoki/go-env/service"
)

type SampleHandler struct {
	ctx           context.Context
	sampleService service.SampleServiceInterface
}

func (h *SampleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, _ := h.sampleService.GetUserByUserId(h.ctx, "1")
	json.NewEncoder(w).Encode(user)
}

func NewSampleHandler(ctx context.Context, sampleService service.SampleServiceInterface) *SampleHandler {
	return &SampleHandler{
		ctx:           ctx,
		sampleService: sampleService,
	}
}
