package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "file-transfer-example",
	Short: "A simple file transfer example using Cloudflare Tunnel",
}
