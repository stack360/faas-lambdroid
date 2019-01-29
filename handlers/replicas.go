// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "encoding/json"
    "net/http"
    "strings"

    "github.com/stack360/faas-lambdroid/lambdroid"
    "github.com/openfaas/faas/gateway/requests"
)

func MakeReplicaUpdater(towerClient lambdroid.LambdroidTowerClient) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        serviceName := vars["name"]
        updates := make(map[string]interface {})
        _, err := towerClient.UpdateService(serviceName, updates)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }
        w.WriteHeader(http.StatusAccepted)
    }
}

func MakeReplicaReader(towerClient lambdroid.LambdroidTowerClient) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        urlPath := r.URL.Path
        if urlPath == "" {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        s := strings.Split(urlPath, "/")
        functionName := s[len(s)-1]
        functions, err := getFunctions(towerClient)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        var found *requests.Function
        for _, function := range functions {
            if function.Name == functionName {
                found = &function
                break
            }
        }

        if found == nil {
            w.WriteHeader(http.StatusNotFound)
            return
        }
        functionBytes, _ := json.Marshal(found)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(functionBytes)
    }
}
