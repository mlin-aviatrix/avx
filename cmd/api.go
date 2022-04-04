package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/mlin-aviatrix/avx/color"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Make a v2.5 API call to the controller.",
	Args:  cobra.MinimumNArgs(2),
	RunE:  apiFunc,
}

func apiFunc(cmd *cobra.Command, args []string) error {
	client, err := getClient()
	if err != nil {
		return jsonErr("could not get client", err)
	}

	method := args[0]
	endpoint := args[1]
	Url := fmt.Sprintf("https://%s/v2.5/api/%s", client.ControllerIP, endpoint)

	var data map[string]interface{}

	if len(args) > 2 {
		data = make(map[string]interface{})

		for _, v := range args[2:] {
			parts := strings.Split(v, "=")
			if len(parts) != 2 {
				return jsonErr(fmt.Sprintf("invalid format for API params, expected 'key=value', got %q", v), nil)
			}
			data[parts[0]] = parts[1]
		}
	}

	var dataBuffer bytes.Buffer
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf(color.Sprint("marshalling json data: %v", color.Red), err)
	}

	err = json.Indent(&dataBuffer, jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf(color.Sprint("indenting json data: %v", color.Red), err)
	}
	if !JsonOnly {
		fmt.Printf("controller IP: %s\n", client.ControllerIP)
		fmt.Printf("request url: %s\n", Url)
		fmt.Printf("request body:\n"+color.Sprint("%s\n", color.Green), dataBuffer.String())
	}

	start := time.Now()

	resp, err := client.RequestContext25(context.Background(), strings.ToUpper(method), Url, nil)

	end := time.Now()
	if !JsonOnly {
		fmt.Printf("latency: %dms\n", end.Sub(start).Milliseconds())
	}
	if err != nil {
		return fmt.Errorf(color.Sprint("non-nil error from API: %v", color.Red), err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf(color.Sprint("error reading response body %q failed: %v", color.Red), endpoint, err)
	}
	b := buf.Bytes()
	if JsonOnly {
		fmt.Println(string(b))
	} else {
		var pp bytes.Buffer
		err = json.Indent(&pp, b, "", "  ")
		fmt.Printf("response status code: %d\n", resp.StatusCode)
		fmt.Printf("response body:\n%s\n", color.Sprint(pp.String(), color.Green))
	}

	return nil
}
