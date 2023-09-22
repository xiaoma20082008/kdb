//
// File: kdb_admin.go
// Project: server
// File Created: 2023-09-08
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-08 10:53:19
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

package server

import (
	"context"
	"fmt"
	"io"
	"kdb/conf"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type KdbAdmin struct {
	io.Closer

	name string
	port int32
	addr string

	server *http.Server

	closed chan bool
}

func (admin *KdbAdmin) Start() error {
	go func() {
		err := admin.server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (admin *KdbAdmin) Close() error {
	log.Println("KdbAdmin closing ...")
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = admin.server.Shutdown(ctx)
	err = admin.server.Close()
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("KdbAdmin closed !!!")
	admin.closed <- true
	return err
}

func (admin *KdbAdmin) Await() {
	<-admin.closed
}

func createRouter() http.Handler {
	engine := gin.Default()
	prometheus := promhttp.Handler()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "welcome to kdb !!!")
	})
	engine.GET("/metrics", func(ctx *gin.Context) {
		prometheus.ServeHTTP(ctx.Writer, ctx.Request)
	})
	return engine
}

func createAdmin(conf *conf.ServerConf) *KdbAdmin {
	admin := new(KdbAdmin)
	admin.name = conf.MainName
	admin.port = conf.CtrlPort
	admin.addr = fmt.Sprintf(":%d", conf.CtrlPort)
	admin.server = &http.Server{
		Addr:    admin.addr,
		Handler: createRouter(),

		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	admin.closed = make(chan bool)
	return admin
}
