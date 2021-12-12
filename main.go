package main

import (
	"fmt"
	"net/http"
	"restandxorm/misc"

	"github.com/sirupsen/logrus"
)

func main() {

	e, err := misc.EngineInit()
	if err != nil {
		fmt.Println("Create Engine failed: ", err)
	}

	crdClient, err := misc.InitCrdClient()
	if err != nil {
		fmt.Println("Create crdclient failed: ", err)
	}

	err = misc.SyncDMetaFromAPIServer(e, crdClient)
	if err != nil {
		fmt.Println("Sync DeviceMeta failed: ", err)
	}
	//
	err = misc.SyncDeviceModelFromDB(e, crdClient)
	if err != nil {
		fmt.Println("Sync DeviceModelMeta failed: ", err)
	}

	//http://luo980.japanwest.cloudapp.azure.com
	fmt.Printf("Starting server at port 80\n")
	//Query
	http.HandleFunc("/edges/addEdge", misc.AddEdge)
	http.HandleFunc("/edges/showEdges", misc.ShowEdges)
	http.HandleFunc("/device/addDevice", misc.AddDevice)
	http.HandleFunc("/device/showDevices", misc.ShowDevices)
	http.HandleFunc("/device/deleteDevice", misc.DeleteDevice)
	http.HandleFunc("/device/queryDevice", misc.QueryDevices)
	// http.HandleFunc("/command", command)

	http.HandleFunc("/get", misc.Get)
	http.HandleFunc("/post", misc.Post)
	// http.HandleFunc("/put", Put)
	// http.HandleFunc("/delete", Delete)

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		logrus.Fatal("ListenAndServe err:", err)
	}
}
