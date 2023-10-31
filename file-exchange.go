package YaDiskAPIClient

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(filepath string, url string) (statusCode int, err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	statusCode = resp.StatusCode
	if statusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %d", statusCode)
		return
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}

	return
}

func upload(filepath string, url string) (statusCode int, err error) {
	data, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer data.Close()

	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	statusCode = res.StatusCode
	if err != nil {
		return
	}
	defer res.Body.Close()

	return
}
