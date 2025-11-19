package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hw2/api"
	"log"
	"net/http"
	"time"
)

func main() {
	base := "http://localhost:8080"
	client := http.Client{}

	// 1) GET /version
	resp, err := client.Get(base + "/version")
	if err != nil {
		log.Println("error getting version:", err)
		return
	}
	var verResp api.VersionResponse
	json.NewDecoder(resp.Body).Decode(&verResp)
	resp.Body.Close()

	// 2) POST /decode
	original := "blablabla"
	enc := base64.StdEncoding.EncodeToString([]byte(original))
	reqBody := api.DecodeRequest{InputString: enc}
	b, _ := json.Marshal(reqBody)
	resp2, err := client.Post(base+"/decode", "application/json", bytes.NewReader(b))
	if err != nil {
		log.Println("error posting decode:", err)
		return
	}
	var decResp api.DecodeResponse
	json.NewDecoder(resp2.Body).Decode(&decResp)
	resp2.Body.Close()

	// 3) GET /hard-op
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	req3, _ := http.NewRequestWithContext(ctx, http.MethodGet, base+"/hard-op", nil)
	resp3, err := client.Do(req3)
	hardOK := false
	statusCode := 0
	if err != nil {
		if ctx.Err() != nil {
			log.Println(ctx.Err())
		} else {
			log.Println(err)
		}
	} else {
		statusCode = resp3.StatusCode
		if statusCode == http.StatusOK {
			hardOK = true
		}
		resp3.Body.Close()
	}

	fmt.Println(verResp.Version)
	fmt.Println(decResp.OutputString)
	if ctx.Err() == nil {
		fmt.Printf("%v, %d\n", hardOK, statusCode)
	}
}
