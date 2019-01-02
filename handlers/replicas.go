// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "net/http"

    "github.com/stack360/faas-mq/mq"
)

func MakeReplicaUpdater(messageSender mq.MessageSender) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        serviceName := vars["name"]
        updates := make(map[string]string)
        _, err := messageSender.UpdateService(serviceName, updates)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }
        w.WriteHeader(http.StatusAccepted)
    }
}

func MakeReplicaReader(messageSender mq.MessageSender) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        serviceName := vars["name"]
        services, err := messageSender.ListServices()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        replicaCounter := 0
        for _, sName := range services {
            if serviceName == sName {
                counter ++
            }
        }
        if counter == 0 {
            w.WriterHeader(http.StatusNotFound)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriterHeader(http.StatusOK)
        w.Write("Found Service")
    }
}
