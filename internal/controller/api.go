package controller

import (
	"encoding/json"
	"github.com/igorkichuk/ciklum/internal/usecase"
	"net/http"
)

type StatusResp struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type ApiController struct {
	articleUsecase usecase.Article
}

func NewApiController(articleUsecase usecase.Article) *ApiController {
	return &ApiController{
		articleUsecase: articleUsecase,
	}
}

func (c *ApiController) TwoResponses(w http.ResponseWriter, r *http.Request) {
	resp, err := c.articleUsecase.SplitResponses()
	if err != nil {
		http.Error(w, Err500, http.StatusInternalServerError)
	}

	printResult(w, resp)
}

func printResult(w http.ResponseWriter, jsonAnswer interface{}) {
	body, err := json.Marshal(jsonAnswer)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		http.Error(w, Err500, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
