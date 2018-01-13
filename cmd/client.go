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
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/zeromq/goczmq"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "ZeroMQ Client Tester",

	Run: func(cmd *cobra.Command, args []string) {
		sockType, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Fatal(err)
		}
		endpoint, err := cmd.Flags().GetString("endpoint")
		if err != nil {
			log.Fatal(err)
		}
		client(sockType, endpoint)
	},
}

func client(sockType string, endpoint string) {
	var client *goczmq.Channeler
	switch sockType {
	case "push":
		client = goczmq.NewPushChanneler(endpoint)
	case "pub":
		client = goczmq.NewPubChanneler(endpoint)
	case "dealer":
		client = goczmq.NewDealerChanneler(endpoint)
	default:
		log.Fatalf("invalid type %s", sockType)
	}
	defer client.Destroy()

	log.Println("client created and connected")
	for i := 0; i < 100; i++ {
		msg := []byte(fmt.Sprintf("Hello %v", time.Now()))
		client.SendChan <- [][]byte{msg}

		log.Printf("client sent %s", msg)
		time.Sleep(50 * time.Millisecond)
	}
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.Flags().StringP("type", "t", "push", "push, pub, dealer")
	clientCmd.Flags().StringP("endpoint", "e", "tcp://127.0.0.1:5555", "endpoint url")
}
