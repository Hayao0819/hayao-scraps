package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serverCmd = cobra.Command{

	Use:   "server",
	Short: "Start the file transfer server",
	RunE: func(cmd *cobra.Command, args []string) error {

		e := gin.Default()
		e.PUT("/upload", func(c *gin.Context) {

			tmpdir, err := os.MkdirTemp("", "file-transfer")
			if err != nil {
				c.JSON(500, gin.H{"error": "failed to create temp dir"})
				return
			}
			defer os.RemoveAll(tmpdir)
			file, err := c.FormFile("file")
			if err != nil {
				c.JSON(400, gin.H{"error": "failed to get file"})
				return
			}
			err = c.SaveUploadedFile(file, filepath.Join(tmpdir, file.Filename))
			if err != nil {
				c.JSON(500, gin.H{"error": "failed to save file"})
				return
			}

			// show file content to stdout
			f, err := os.Open(filepath.Join(tmpdir, file.Filename))
			if err != nil {
				c.JSON(500, gin.H{"error": "failed to open file"})
				return
			}
			defer f.Close()
			content, err := io.ReadAll(f)
			if err != nil {
				c.JSON(500, gin.H{"error": "failed to read file"})
				return
			}
			fmt.Println(string(content))
			c.JSON(200, gin.H{"message": "file uploaded successfully"})
		})

		// start the server
		return e.Run(":8080")
	},
}

func init() {
	rootCmd.AddCommand(&serverCmd)
}
