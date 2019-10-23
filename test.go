package main

import (
  "log"
  "sync"
  "github.com/nsqio/go-nsq"
)

func main() {

  wg := &sync.WaitGroup{}
  wg.Add(5000);
  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer("stream", "gotest", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      log.Printf("Got a message: %v", message)
      return nil
  }))
  err := q.ConnectToNSQD("real3.ktbfuso.id:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  wg.Wait()

}
