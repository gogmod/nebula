package main

import (
	"log"

	nebula "github.com/vesoft-inc/nebula-go"
	graph "github.com/vesoft-inc/nebula-go/nebula/graph"
)

func main() {
	client, err := nebula.NewClient("127.0.0.1:3699")
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Connect("user", "password"); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()
	checkResp := func(prefix string, resp *graph.ExecutionResponse) {
		if resp.GetErrorCode() != graph.ErrorCode_SUCCEEDED {
			log.Fatalf("%s, ErrorCode: %v, ErrorMsg: %s", prefix, resp.GetErrorCode(), resp.GetErrorMsg())
		}
	}
	if resp, err := client.Execute("SHOW HOSTS;"); err != nil {
		log.Fatal(err)
	} else {
		checkResp("show hosts", resp)
		log.Println("success")
	}

	if resp, err := client.Execute("CREATE SPACE client_test(partition_num=1024, replica_factor=1);"); err != nil {
		log.Fatalf(err.Error())
	} else {
		checkResp("create space", resp)
	}
}