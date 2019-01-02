// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "encoding/json"
    "io/ioutil"
    "net/http"

    "github.com/openfaas/faas/gateway/requests"
    "github.com/stack360/faas-mq/mq"
)

func MakeDeployHandler(messageSender mq.MessageSender) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        defer r.Body.Close()
        body, _ := ioutil.ReadAll(r.Body)

        request := requests.CreateFunctionRequest{}
        err := json.Unmarshal(body, &request)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        servicSpec := map[string]interface{} {}
        _, err := messageSender.AddService(serviceSpec)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }
        w.WriteHeader(http.StatusAccepted)
    }
}
