package main

import (
	"amt-proxy/internal/amt"
	"amt-proxy/pkg/digesthttp"
	"amt-proxy/pkg/utils"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var listenAddr string
	flag.StringVar(&listenAddr, "listen", "0.0.0.0:26992", "listen address")
	flag.Parse()

	amtCommand := amt.NewAMTCommand()
	rc, err := amtCommand.Initialize()
	if rc != utils.Success || err != nil {
		log.Fatalf("AmtNotDetected: %d/%v", rc, err)
	}

	lsa, err := amtCommand.GetLocalSystemAccount()
	if err != nil {
		log.Fatalf("GetLocalSystemAccount: %v", err)
	}

	log.Printf("USERNAME: %s", lsa.Username)
	log.Printf("PASSWORD: %s", lsa.Password)

	httpClient := &http.Client{
		Transport: http.DefaultTransport,
	}
	session := digesthttp.New(httpClient)
	session.SetAuth(lsa.Username, lsa.Password)

	server := http.Server{
		Addr: listenAddr,
	}
	var handleFunc http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
		log.Println("REQUEST: " + req.Method + " " + req.RequestURI)

		var requestUrl url.URL = *req.URL
		requestUrl.Scheme = "http"
		requestUrl.Host = "127.0.0.1:16992"

		req.URL = &requestUrl
		req.RequestURI = ""

		localRes, err := session.Do(req)

		var body []byte
		if err == nil {
			body, err = io.ReadAll(localRes.Body)
		}
		if err != nil {
			log.Println("\tREQUEST ERROR: ", err)
			res.WriteHeader(500)
			res.Write([]byte{})
		} else {
			log.Printf("\tRESPONSE CODE: %d (%s)", localRes.StatusCode, localRes.Status)
			for key, values := range localRes.Header {
				for _, value := range values {
					localRes.Header.Add(key, value)
				}
			}
			res.WriteHeader(localRes.StatusCode)
			res.Write(body)

			if localRes.StatusCode == 401 {
				lsa, err := amtCommand.GetLocalSystemAccount()
				if err != nil {
					log.Fatalf("GetLocalSystemAccount: %v", err)
				}
				session.SetAuth(lsa.Username, lsa.Password)
			}
		}
	}
	server.Handler = handleFunc
	log.Println("Listen ", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
