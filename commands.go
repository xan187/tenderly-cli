package main

import (
	"github.com/spf13/cobra"
	"github.com/tenderly/tenderly-cli/proxy"
)

var targetHost string
var targetPort string
var targetSchema string
var proxyHost string
var proxyPort string
var path string
var network string

func init() {
	proxyCmd.PersistentFlags().StringVar(&targetSchema, "target-schema", "http", "Blockchain rpc schema.")
	proxyCmd.PersistentFlags().StringVar(&targetHost, "target-host", "127.0.0.1", "Blockchain rpc host.")
	proxyCmd.PersistentFlags().StringVar(&targetPort, "target-port", "8545", "Blockchain rpc port.")
	proxyCmd.PersistentFlags().StringVar(&proxyHost, "proxy-host", "127.0.0.1", "Proxy host.")
	proxyCmd.PersistentFlags().StringVar(&proxyPort, "proxy-port", "9545", "Proxy port.")
	proxyCmd.PersistentFlags().StringVar(&path, "path", "", "Path to the project build folder.")
	proxyCmd.PersistentFlags().StringVar(&network, "network", "", "Network id.")

	rootCmd.AddCommand(proxyCmd)
}

var rootCmd = &cobra.Command{
	Use:   "tenderly",
	Short: "Tenderly helps you observe your contracts in any environment.",
	Long:  "Tenderly is a development tool for smart contract.",
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Proxy",
	Run: func(cmd *cobra.Command, args []string) {
		proxy.Start(targetSchema, targetHost, targetPort, proxyHost, proxyPort, path, network)
	},
}
