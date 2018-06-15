package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/timaa/trafficStat/daemon"
	"github.com/timaa/trafficStat/repository"
	"log"
	"github.com/timaa/trafficStat/rabbit"
	"encoding/json"
	"github.com/timaa/trafficStat/DTO"
	"github.com/timaa/trafficStat/worker"
)

const (
	WorkerCount = 10
	WorkerInputChanSize = 4
)

func init() {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	} else {
		fmt.Println("Service RUN on PROD mode")
	}

}

func assembleConfig() (*daemon.Config){
	cfg := &daemon.Config{}
	cfg.Db.ConnectionString = fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.dbname"))
	fmt.Print(cfg.Db.ConnectionString)
	cfg.Rabbit.ConnectionString = fmt.Sprintf("amqp://%s:%s@%s:%s",
		viper.GetString("rabbit.user"),
		viper.GetString("rabbit.password"),
		viper.GetString("rabbit.host"),
		viper.GetString("rabbit.port"))

	return cfg
}

func main() {
	cfg := assembleConfig()
	db, err := daemon.Run(cfg);
	if err != nil {
		log.Printf("error %v \n", err)
	}

	deviceRepository := repository.New(db)

	devices, err := deviceRepository.GetAll()
	for _, device := range devices {
		fmt.Println(device)
	}
	msgs, err := rabbit.Init(cfg.Rabbit)
	if err != nil {
		fmt.Printf("rabbit error:%v \n", err)
	}



	workerInput := make(chan *DTO.TrafficDto, WorkerInputChanSize)

	for i := 0; i < WorkerCount; i++ {
		w := &worker.Worker{}
		go w.Run(db, workerInput, i)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			 b := &DTO.TrafficDto{}
			if err := json.Unmarshal(d.Body, b); err != nil {
				log.Printf("%v \n", err)
			}
			workerInput <- b
			d.Ack(false)
		}
	}()

	fmt.Println("main \n")
	<-forever
}
