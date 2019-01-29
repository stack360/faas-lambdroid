// Copyright (c) Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"
    "strings"
    "github.com/stack360/faas-lambdroid/lambdroid"

    "io/ioutil"
)

func MakeProxy(towerClient lambdroid.LambdroidTowerClient, stackName string) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        if r.Method != "POST" {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        urlPath := r.URL.Path
        if urlPath == "" {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        s := strings.Split(urlPath, "/")
        functionName := s[len(s)-1]

        serviceParams, _ := ioutil.ReadAll(r.Body)
        paramsMap := make(map[string]interface{})
        unmarshalErr := json.Unmarshal(serviceParams, &paramsMap)
        if unmarshalErr != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        paramsMap["action_type"] = "run_app"
        paramsMap["action_name"] = functionName
        defer r.Body.Close()

        params, marshalErr := json.Marshal(paramsMap)
        if marshalErr != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        fmt.Println("Marshal successful: ", params)
        status, err := towerClient.InvokeService(stackName, params)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(status))
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte(status))
    }
}
