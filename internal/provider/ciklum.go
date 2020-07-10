package provider

import (
	"encoding/json"
	"fmt"
	"github.com/igorkichuk/ciklum/internal/models"
	"io/ioutil"
	"net/http"
)

type ciklumProvider struct{}

func NewCiklumProvider() CiklumProvider {
	return &ciklumProvider{}
}

func (p *ciklumProvider) GetResponse(url string) (models.CiklumResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return models.CiklumResponse{}, err
	}

	if resp.StatusCode >= 300 {
		return models.CiklumResponse{}, fmt.Errorf("invalid response: %v", resp)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.CiklumResponse{}, err
	}

	var ciklumResponse models.CiklumResponse
	err = json.Unmarshal(body, &ciklumResponse)
	if err != nil {
		return models.CiklumResponse{}, err
	}

	return ciklumResponse, nil
}
