package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"json-bin/bins"
	"json-bin/config"
	"json-bin/output"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Api struct {
	config *config.Config
}

type Response struct {
	Metadata bins.Bin       `json:"metadata"`
	Record   bins.BinRecord `json:"record"`
}

const baseUrl = "https://api.jsonbin.io/v3/b"

func New(config *config.Config) *Api {
	return &Api{
		config: config,
	}
}

func (api *Api) request(method string, url string, body []byte) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	var reader io.Reader
	if body == nil {
		reader = nil
	} else {
		reader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, baseUrl+url, reader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("X-Master-Key", api.config.XMasterKey)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		// fmt.Println(response)
		return nil, fmt.Errorf("status not 200: %s", response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func handleBinResponse(method string, response []byte) (*bins.Bin, error) {
	var data Response
	err := json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}

	color.Cyan(method + " bin record\n")
	output.PrintJson(data.Record)

	if method != "PUT" {
		color.Cyan(method + " bin metadata\n")
		output.PrintJson(data.Metadata)
	}

	return &data.Metadata, err
}

func (api *Api) Get(id string) error {
	if id == "" {
		return errors.New("empty id")
	}
	response, err := api.request("GET", "/"+id, nil)
	if err != nil {
		return err
	}

	_, err = handleBinResponse("GET", response)
	if err != nil {
		return err
	}

	return nil
}

func binRecordToJson(record *bins.BinRecord) ([]byte, error) {
	data, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (api *Api) Create(record *bins.BinRecord) (*bins.Bin, error) {

	body, err := binRecordToJson(record)
	if err != nil {
		return nil, err
	}

	response, err := api.request("POST", "", body)
	if err != nil {
		return nil, err
	}

	bin, err := handleBinResponse("POST", response)
	if err != nil {
		return nil, err
	}

	return bin, nil
}

func (api *Api) Update(id string, record *bins.BinRecord) (*bins.Bin, error) {
	if id == "" {
		return nil, errors.New("empty id")
	}

	body, err := binRecordToJson(record)
	if err != nil {
		return nil, err
	}

	response, err := api.request("PUT", "/"+id, body)
	if err != nil {
		return nil, err
	}

	bin, err := handleBinResponse("PUT", response)
	if err != nil {
		return nil, err
	}

	return bin, nil
}

func (api *Api) Delete(id string) error {
	if id == "" {
		return errors.New("empty id")
	}
	_, err := api.request("DELETE", "/"+id, nil)
	if err != nil {
		return err
	}

	return nil
}
