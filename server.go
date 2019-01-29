// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/openfaas/faas-provider"
	bootTypes "github.com/openfaas/faas-provider/types"
	"github.com/stack360/faas-lambdroid/handlers"
	"github.com/stack360/faas-lambdroid/lambdroid"
)


func main() {
	towerURL := os.Getenv("LAMBDROID_TOWER_URL")
	stackName := os.Getenv("STACK_NAME")


	// creates the tower client config object
	config, err := lambdroid.NewClientConfig(
		stackName,
		towerURL,
	)
	if err != nil {
		panic(err.Error())
	}

	// create the lambdroid lambdroid tower client
	towerClient, err := lambdroid.NewTowerClientFromConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Created New Tower Client.")

/*
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
	*/
	bootstrapHandlers := bootTypes.FaaSHandlers{
		FunctionProxy:  handlers.MakeProxy(towerClient, config.FunctionsStackName).ServeHTTP,
		DeleteHandler:  handlers.MakeDeleteHandler(towerClient).ServeHTTP,
		DeployHandler:  handlers.MakeDeployHandler(towerClient).ServeHTTP,
		FunctionReader: handlers.MakeFunctionReader(towerClient).ServeHTTP,
		ReplicaReader:  handlers.MakeReplicaReader(towerClient).ServeHTTP,
		ReplicaUpdater: handlers.MakeReplicaUpdater(towerClient).ServeHTTP,
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
