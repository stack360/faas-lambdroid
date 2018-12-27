// Copyright (c) Xicheng Chang 2018. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package mq

import (
    "fmt"
    "github.com/streadway/amqp"
    "errors"
)

// MessageSender acts like a client that triggers actions by sending messages to mq
type MessageSender interface {
    ListServices() ([]string, error)
    GetServiceByName(serviceName string) (map[string]interface{}, error)
    AddService(serviceSpec map[string]interface{}) (map[string]interface{}, error)
    DeleteService(serviceName, string) error
    UpdateService(serviceName string, updatedServiceSpec map[string]interface{}) (map[string]interface{}, error)
    InvokeService(serviceName string, serviceParams []byte) (string, error)
}

type Sender struct {
    config        *Config
}

func NewSenderFromConfig(config *config) (MessageSender, error) {
    conn, err := amqp.Dial(config.RabbitURL)
    if err != nil {
        return nil, errors.New("Could not connect to MQ!")
    }
    defer conn.Close()

    s, sErr := conn.Channel()
    if sErr != nil {
        return nil, errors.New("Could not open channel to create sender!")
    }
    return &s, nil
}

// TODO: Most of these functions are placeholders for now. Implement these after Android master app has the functionalities.
func (s *Sender) ListServices() ([]string, error) {
    services := []string{""}
    return services, nil
}

func (s *Sender) GetServiceByName(serviceName string) (map[string]interface{}, error) {
    service := map[string]interface{} {}
    return service, nil
}

func (s *Sender) AddService(serviceSpec map[string]interface{}) (map[string]interface{}, error) {
    service := map[string]interface{} {}
    return service, nil
}

func (s *Sender) DeleteService(serviceName string) (error) {
    return nil
}

func (s *Sender) UpdateService(serviceName string, updatedServiceSpec map[string]interface{}) (map[string]interface{}, error) {
    service := map[string]interface{} {}
    return service, nil
}

func (s *Sender) InvokeService(serviceName string, serviceParams []byte]) (string, error) {
    s, err := NewSenderFromConfig()
    if err != nil {
        return "Error encountered when instantianting sender.", errors.New(err.Error())
    }
    body := serviceParams
    q, err := s.QueueDeclare(
        "android",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return "Error encountered when declaring queue.", errors.New(err.Error())
    }

    publishErr := s.Publish(
        "",
        q.Name,
        false,
        false,
        amqp.Publishing(
            ContentType: "text/plain",
            Body:        body,
        ),
    )
    if publishErr != nil {
        return "Error encountered when publishing message.", errors.New(err.Error())
    }
    return "Successfuly published message.", nil
}
