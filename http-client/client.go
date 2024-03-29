package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/FUnigrad/funiverse-workspace-service/config"
	"github.com/FUnigrad/funiverse-workspace-service/model"
)

type HttpClient struct {
	Hostname string
	Client   *http.Client
}

func NewClient(config config.Config) (*HttpClient, error) {
	var hostname string
	if config.Enviroment == "local" {
		hostname = "authen.system.funiverse.world"
	} else if config.Enviroment == "prod" {
		hostname = "authen-service:8000"
	} else {
		return nil, errors.New("configuration incorrect at env")
	}

	httpClient := HttpClient{
		Hostname: hostname,
		Client:   &http.Client{},
	}
	return &httpClient, nil
}

func (client *HttpClient) GetAllWorkspace(token string) (workspaces []model.Workspace) {

	url := fmt.Sprintf("http://%s/workspace", client.Hostname)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)

	resp, err := client.Client.Do(req)

	if err != nil {
		log.Print(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	if err = json.Unmarshal(body, &workspaces); err != nil {
		log.Print(err.Error())
		return nil
	}

	return
}

func (client *HttpClient) GetWorkspaceById(id int, token string) (workspace *model.Workspace) {
	url := fmt.Sprintf("http://%s/workspace/%d", client.Hostname, id)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", token)

	resp, err := client.Client.Do(req)

	if err != nil {
		log.Print(err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print(err.Error())
		return nil
	}

	if err = json.Unmarshal(body, &workspace); err != nil {
		log.Print(err.Error())
		return nil
	}

	return
}

func (client *HttpClient) CreateWorkspace(workspace model.WorkspaceDTO, token string) (*model.Workspace, error) {
	url := fmt.Sprintf("http://%s/workspace", client.Hostname)

	request_body, _ := json.Marshal(workspace)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(request_body))

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		log.Print(string(body))
		return nil, errors.New(string(body))
	}

	var result *model.Workspace

	if err = json.Unmarshal(body, &result); err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return result, nil

}

func (client *HttpClient) DeleteWorkspace(id int, token string) bool {
	url := fmt.Sprintf("http://%s/workspace/%d", client.Hostname, id)

	req, _ := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", token)
	resp, _ := client.Client.Do(req)

	if resp.StatusCode != 200 {
		return false
	} else {
		return true
	}
}
