// Copyright (c) Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package lambdroid

type Config struct {
    // RabbitMQ service URL
    FunctionsStackName string
    LambdroidTowerURL string
}

func NewClientConfig(sn string, url string) (*Config, error) {
    config := Config{
        FunctionsStackName: sn,
        LambdroidTowerURL:  url,
    }
    return &config, nil
}
