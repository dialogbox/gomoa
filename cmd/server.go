// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/zeromq/goczmq"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "ZeroMQ Server Tester",

	Run: func(cmd *cobra.Command, args []string) {
		sockType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}
		endpoint, err := cmd.Flags().GetString("endpoint")
		if err != nil {
			log.Fatal(err)
		}
		server(sockType, endpoint)
	},
}

func server(serverType string, endpoint string) {
	var server *goczmq.Channeler
	switch serverType {
	case "pull":
		server = goczmq.NewPullChanneler(endpoint)
	case "sub":
		server = goczmq.NewSubChanneler(endpoint)
		server.Subscribe("")
	case "router":
		server = goczmq.NewRouterChanneler(endpoint)
	default:
		log.Fatalf("invalid type %s", serverType)
	}
	defer server.Destroy()

	log.Println("server created and bound")

	for {
		message := <-server.RecvChan
		log.Printf("router received '%s'", message[0])
		time.Sleep(500 * time.Millisecond)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().StringP("type", "t", "pull", "pull, sub, router")
	serverCmd.Flags().StringP("endpoint", "e", "tcp://127.0.0.1:5555", "endpoint url")
}
