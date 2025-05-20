package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var clientCmd = cobra.Command{
	Use:   "client",
	Short: "Start the file transfer client",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("file path is required")
		}

		filePath := args[0]

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", filePath)
		}

		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %s, error: %v", filePath, err)
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			return fmt.Errorf("failed to create form file: %v", err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return fmt.Errorf("failed to copy file content: %v", err)
		}

		err = writer.Close()
		if err != nil {
			return fmt.Errorf("failed to close writer: %v", err)
		}

		serverAddress := "http://test.hayao0819.com/upload"
		req, err := http.NewRequest("PUT", serverAddress, body)
		if err != nil {
			return fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := http.DefaultClient
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("failed to send request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("server returned non-OK status: %v", resp.Status)
		}

		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}

		fmt.Println(string(respBody))

		return nil
	},
}

func init(){
	rootCmd.AddCommand(&clientCmd)
}
