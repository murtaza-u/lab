package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	timeout := time.Second * 5

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: timeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err = cli.Put(ctx, "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.Get(ctx, "foo", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}

	for _, kv := range resp.Kvs {
		fmt.Printf("%s\n", string(kv.Value))
	}
}
