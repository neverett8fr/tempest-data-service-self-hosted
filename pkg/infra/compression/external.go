package compression

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// "convert" - convert either to compressed, or uncompressed
func (cp *CompressionProvider) callCompress(data []byte) ([]byte, error) {

	route := fmt.Sprintf("%s/%s", cp.Path, compressSubPath)

	request, err := http.NewRequest(http.MethodPost, route, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request, err %v", err)
	}

	request.Header.Add("Accept", "*/*") // change to consts
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error calling compression service, err %v", err)
	}

	return readBody(*resp)
}

// "convert" - convert either to compressed, or uncompressed
func (cp *CompressionProvider) callDecompress(data []byte) ([]byte, error) {

	route := fmt.Sprintf("%s/%s", cp.Path, decompressSubPath)

	request, err := http.NewRequest(http.MethodPost, route, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request, err %v", err)
	}

	request.Header.Add("Accept", "*/*") // change to consts
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error calling compression service, err %v", err)
	}

	return readBody(*resp)
}

func readBody(resp http.Response) ([]byte, error) {

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body, err %v", err)
	}

	return body, nil
}
