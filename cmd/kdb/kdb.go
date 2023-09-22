//
// File: kdb.go
// Project: cmd
// File Created: 2023-09-06
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-06 18:48:04
// -----
//
// Copyright (C) xiaoma20082008. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"bufio"
	"fmt"
	"kdb/conf"
	"kdb/server"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	root := &cobra.Command{}
	help := &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
		Args:  nil,
		Run: func(cmd *cobra.Command, args []string) {
			runHelp()
		},
	}
	ctrl := &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
		Args:  nil,
		Run: func(cmd *cobra.Command, args []string) {
			runExec()
		},
	}
	root.AddCommand(help, ctrl)
}

func runHelp() {

}

func runExec() {

}

func main() {
	file, err := os.Open("/Users/admin/Desktop/mysql.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			fmt.Println(scanner.Text() + ":\"" + scanner.Text() + "\",")
		} else {
			fmt.Println()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main0() {
	serverConf := conf.ServerConf{}
	file := "kdb.yaml"
	data, err := os.ReadFile(file)
	if err != nil {
		slog.Error("Read yaml file failed:", err)
		return
	}
	err = yaml.Unmarshal(data, &serverConf)
	if err != nil {
		slog.Error("Decode yaml file failed:", err)
		return
	}
	proxy := server.CreateProxy(&serverConf)
	err = proxy.Start()

	if err != nil {
		_ = proxy.Close()
	} else {
		slog.Info("Kdb proxy started")
		go func() {
			quit := make(chan os.Signal)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			proxy.Close()
		}()
		proxy.Await()
	}
}
