// Copyright (c) Xicheng Chang 2018-2019. All rights reserved.
// Licensed under the MIT license.

package mq

import (
    "errors"

    "github.com/streadway/amqp"
)

// MessageSender acts like a client that triggers actions by sending messages to mq
type MessageSender interface {
    ListServices() ([]string, error)
    GetServiceByName(serviceName string) (map[string]interface{}, error)
    AddService(serviceSpec map[string]interface{}) (map[string]interface{}, error)
    DeleteService(serviceName string) error
    UpdateService(serviceName string, updatedServiceSpec map[string]interface{}) (map[string]interface{}, error)
    InvokeService(serviceName string, serviceParams []byte) (string, error)
}

type Sender struct {
    config        *Config
}

func NewSenderFromConfig(config *Config) (MessageSender, error) {
    s := Sender{
        config:    config,
    }
    return &s, nil
}

// TODO: Most of these functions are placeholders for now. Implement these after Android master app has the functionalities.
func (s *Sender) ListServices() ([]string, error) {
    services := []string{"labelmaker2"}
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

func (s *Sender) InvokeService(serviceName string, serviceParams []byte) (string, error) {
    body := serviceParams
    conn, err := amqp.Dial(s.config.RabbitURL)
    if err != nil {
        return err.Error(), err
    }
    defer conn.Close()

    ch, sErr := conn.Channel()
    if sErr != nil {
        return err.Error(), sErr
    }
    q, qErr := ch.QueueDeclare(
        serviceName,
        false,
        false,
        false,
        false,
        nil,
    )
    if qErr != nil {
        return err.Error(), errors.New(err.Error())
    }

    publishErr := ch.Publish(
        "",
        q.Name,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        body,
        },
    )
    if publishErr != nil {
        return "Error encountered when publishing message.", errors.New(err.Error())
    }
    return "Successfuly published message.", nil
}
