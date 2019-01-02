// Copyright (c) Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package mq

type Config struct {
    // RabbitMQ service URL
    FunctionsStackName string
    RabbitURL string
}

func NewMQConfig(url string) (*Config, error) {
    config := Config{
        RabbitURL: url,
    }
    return &config, nil
}
