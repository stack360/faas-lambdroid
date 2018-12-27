// Copyright (c) Alex Ellis 2018, Xicheng Chang 2018. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/openfaas/faas-provider"
	bootTypes "github.com/openfaas/faas-provider/types"
	"github.com/stack360/faas-mq/handlers"
	"github.com/stack360/faas-mq/mq"
)


func main() {
	rabbitURL := os.Getenv("RABBIT_URL")


	// creates the rancher client config
	config, err := mq.NewMQConfig(
		rabbitURL,
	)
	if err != nil {
		panic(err.Error())
	}

	// create the rancher REST client
	mqSender, err := mq.NewSenderFromConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Created New MQ Sender.")

	proxyClient := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   3 * time.Second,
				KeepAlive: 0,
			}).DialContext,
			MaxIdleConns:          1,
			DisableKeepAlives:     true,
			IdleConnTimeout:       120 * time.Millisecond,
			ExpectContinueTimeout: 1500 * time.Millisecond,
		},
	}
	bootstrapHandlers := bootTypes.FaaSHandlers{
		FunctionProxy:  handlers.MakeProxy(&proxyClient, config.FunctionsStackName).ServeHTTP,
		DeleteHandler:  handlers.MakeDeleteHandler(mqSender).ServeHTTP,
		DeployHandler:  handlers.MakeDeployHandler(mqSender).ServeHTTP,
		FunctionReader: handlers.MakeFunctionReader(mqSender).ServeHTTP,
		ReplicaReader:  handlers.MakeReplicaReader(mqSender).ServeHTTP,
		ReplicaUpdater: handlers.MakeReplicaUpdater(mqSender).ServeHTTP,
	}
	var port int
	port = 8080
	bootstrapConfig := bootTypes.FaaSConfig{
		ReadTimeout:  time.Second * 8,
		WriteTimeout: time.Second * 8,
		TCPPort:      &port,
	}

	bootstrap.Serve(&bootstrapHandlers, &bootstrapConfig)

}
