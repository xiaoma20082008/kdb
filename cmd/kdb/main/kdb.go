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
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kdb",
	Short: "A simple database",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start kdb server, usage: kdb start --conf ?",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("conf")
		data, _ := os.ReadFile(file)
		conf := string(data)
		msg := "Kdb starting with config file: " + file
		slog.Info(msg)
		slog.Info(conf)
		slog.Info("Kdb started.")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop kdb server, usage: kdb stop",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Kdb stopping...")
		slog.Info("Kdb stopped.")
	},
}

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload kdb server config, usage: kdb reload",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Kdb reloading...")
		slog.Info("Kdb reload success.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(reloadCmd)
	startCmd.Flags().StringP("conf", "c", "", "config file(default is $HOME/config/.cobra.yaml)")
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
