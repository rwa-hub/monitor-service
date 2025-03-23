package listeners

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBlockNumber() (string, error) {
	resp, err := http.Get("http://host.docker.internal:8546")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return fmt.Sprintf("BlockNumber: %s", string(body)), nil
}
