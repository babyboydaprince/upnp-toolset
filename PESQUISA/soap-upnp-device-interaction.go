package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const soapRequest = `<?xml version="1.0"?>
<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/" s:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
  <s:Body>
    <u:GetExternalIPAddress xmlns:u="urn:schemas-upnp-org:service:WANIPConnection:1" />
  </s:Body>
</s:Envelope>`

func main() {
	controlURL := "http://192.168.1.1:1900/control?WANIPConn1" // Replace with actual URL

	req, err := http.NewRequest("POST", controlURL, bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"urn:schemas-upnp-org:service:WANIPConnection:1#GetExternalIPAddress"`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body)) // The response contains the external IP address
}
