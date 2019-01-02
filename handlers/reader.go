// Copyright (c) Alex Ellis 2016-2018, OpenFaaS Author(s) 2018, Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package handlers

import (
    "bytes"
    "encoding/gob"
    "net/http"

    "github.com/stack360/faas-mq/mq"
)

func MakeFunctionReader(messageSender mq.MessageSender) VarsWrapper {
    return func(w http.ResponseWriter, r *http.Request, vars map[string]string) {
        functions, err := messageSender.ListServices()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        buf := &bytes.Buffer{}
        gob.NewEncoder(buf).Encode(functions)
        functionBytes := buf.Bytes()
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(functionBytes)
    }
}
