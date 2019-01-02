// Copyright (c) 2018-2019 Xicheng Chang
//
// This software is released under the MIT License.

package handlers

import (
    "net/http"

    "github.com/gorilla/mux"
)

type VarsWrapper func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func (varsWrapper VarsWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    varsWrapper(w, r, vars)
}
