package helpers

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func PostForm(url string, path string, token string, content string) ([]byte, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	form, err := writer.CreateFormFile("xml", "xml")
	if err != nil {
		log.Fatal(err)
	}
	form.Write([]byte(content))
	writer.Close()
	request, err := http.NewRequest(http.MethodPost, url+path, body)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("Content-Type", "multipart/form-data; boundary="+writer.Boundary())
	request.Header.Add("Authorization", "bearer "+token)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseBytes, err
}
