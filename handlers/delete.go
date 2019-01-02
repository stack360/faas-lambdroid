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

func MakeDeleteHandler(messageSender mq.MessageSender) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        defer r.Body.Close()
        body, _ := ioutil.ReadAll(r.Body)

        request := requests.DeleteFunctionRequest{}
        err := json.Unmarshal(body, &request)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if len(request.FunctionName) == 0 {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        service, getErr := messageSender.GetServiceByName(request.FunctionName)
        if getErr != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        } else if service == nil {
            w.WriteHeader(http.StatusNotFound)
            return
        }
        delErr := messageSender.DeleteService(request.FunctionName)
        if delErr != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }
        w.WriteHeader(http.StatusOK)
    }
}
