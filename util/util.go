package util

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

func GetJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	return body, err
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	status := resp.StatusCode
	if !(status == 200 || status == 302) {
		return errors.New("could not download jar file")
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func IsOpened(host string, port int) bool {

	timeout := 5 * time.Second
	target := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}