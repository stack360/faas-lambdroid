// Copyright (c) Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "bytes"
    "net/http"

    "github.com/stack360/faas-mq/mq"

    "io/ioutil"
)

func MakeProxy(messageSender mq.MessageSender, stackName string) Vars {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {

        if r.Method != "POST" {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        serviceName := vars["name"]
        serviceParams, _ := ioutil.ReadAll(r.Body)
        defer r.Body.Close()
        status, err := messageSender.InvokeService(serviceName, serviceParams)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write(status)
        }

        w.WriteHeader(http.StatusOK)
        w.Write(status)
    }
}
