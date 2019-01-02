// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "net/http"

    "github.com/openfaas/faas/gateway/requests"
    "github.com/stack360/faas-mq/mq"
)

func MakeFunctionReader(messageSender mq.MessageSender) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        functions, err := messageSender.ListServices()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        functionBytes := string(functions[:])
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(functionBytes)
    }
}
