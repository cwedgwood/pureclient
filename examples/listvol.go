//

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	pureclient "github.com/cwedgwood/pureclient/client"
	"gopkg.in/yaml.v2"
)

const (
	apiEndpoint = "https://1.2.3.4:443" // Usually the ct0/ct1 "vip" IP
	apiToken    = "secret-token-here"   // Obtain this from from Pure UI
)

// list volumes (show fairly raw output)
func main() {
	client := getApiClient()
	vols, httpRes, err := client.VolumesApi.Api28VolumesGet(context.Background(), nil)
	fmt.Println("ERR:", err)
	yamlp(vols.Items)
	fmt.Println(httpRes)
}

func getXAuthToken(httpClient http.Client) string {
	xAuthTokenReq, err := http.NewRequest("POST", apiEndpoint+"/api/2.8/login", nil)
	if err != nil {
		log.Fatal(err)
	}
	xAuthTokenReq.Header.Add("Content-Type", "application/json")
	xAuthTokenReq.Header.Add("api-token", apiToken)
	resp, err := httpClient.Do(xAuthTokenReq)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	fmt.Println("X-Auth-Token: ", resp.Header.Get("X-Auth-Token"))
	return resp.Header.Get("X-Auth-Token")
}

func getApiClient() *pureclient.APIClient {
	httpClient := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	c := pureclient.NewConfiguration()
	c.BasePath = apiEndpoint
	c.DefaultHeader["x-auth-token"] = getXAuthToken(httpClient)
	c.HTTPClient = &httpClient
	apiClient := pureclient.NewAPIClient(c)
	return apiClient
}

//

func yamlp(i interface{}) {
	d, _ := yaml.Marshal(i)
	d = append(d, []byte("\n\n")...)
	os.Stdout.Write(d)
}
