package main

import (
	"context"
	"os"
	"restandxorm/misc"
	"restandxorm/v1alpha2"
	"strings"
)

type CloudDeviceMeta struct {
	ID          string `xorm:"pk"`
	Name        string
	Model       string
	Description string
	How         string
	Where       string
	Node        string
}

func main() {

	e, err := misc.EngineInit()
	if err != nil {
		print("Create Engine failed: ", err)
	}
	//
	//dht11 := CloudDeviceMeta{
	//	ID:          "dht11-sensor-2",
	//	Name:        "dht11-sensor-1",
	//	Model:       "dht11-sensor-model",
	//	Description: "a-temperature-sensor-for-industrial",
	//	How:         "collect",
	//	Where:       "office",
	//	Node:        "kubeedge-node1",
	//}
	//
	////dht12 := CloudDeviceMeta{
	////	ID:          "dht11-sensor-1",
	////}
	//
	////DeviceTableName := "Device"
	errsyc := e.Sync2(new(CloudDeviceMeta))
	if errsyc != nil {
		print("Create table failed. err: ", errsyc)
	}
	//
	// result, err := e.Insert(dht11)
	// if err != nil {
	//	print("Insert Error, err:", err)
	// }
	// print("Insert Result is :", result)
	// //
	// //result, err := e.ID("dht11-sensor-1").Delete(dht12)
	// //if err != nil {
	//	// print("Delete error. err:", err)
	// //}
	// //print("Delete result is ", result)

	ConfigPath := "./config.yaml"
	connectionConfig, err := misc.ParseConfig(ConfigPath)
	if err != nil {
		print("parse config failed.")
		os.Exit(-1)
	}
	crdClient := misc.InitRestClient(connectionConfig.KubeAPIConfig)

	DL := v1alpha2.DeviceList{}
	err = crdClient.RESTClient().Get().Resource("devices").Do(context.Background()).Into(&DL)

	if err != nil {
		print(err)
	}

	//i := v1alpha2.Device{}
	for key, item := range DL.Items {
		description := ""
		how := ""
		where := ""
		//logrus.WithFields(logrus.Fields{
		// "\nindex"		: 	key,
		// "\nName"		:	item.Name,
		// "\nNamespace"	:	item.Namespace,
		// "\nDeviceModel":	item.Spec.DeviceModelRef.Name,
		// "\nNodeName"	:	item.Spec.NodeSelector.NodeSelectorTerms[0].MatchExpressions[0].Values[0],
		// "\nTopic"		:	item.Spec.Data.DataTopic,
		//}).Infof("\nHere's the device cached")
		print(
			"\nindex:\t\t", key,
			"\nName:\t\t", item.Name,
			"\nNamespace:\t", item.Namespace,
			"\nDeviceModel:\t", item.Spec.DeviceModelRef.Name,
			"\nNodeName:\t", item.Spec.NodeSelector.NodeSelectorTerms[0].MatchExpressions[0].Values[0])

		for k, v := range item.ObjectMeta.Labels {
			if strings.Compare(k, "description") == 0 {
				description = v
			} else if strings.Compare(k, "how") == 0 {
				how = v
			} else if strings.Compare(k, "where") == 0 {
				where = v
			}
			print("\n", k, "\t", v)
		}
		for k, v := range item.Spec.Data.DataProperties {
			//logrus.WithFields(logrus.Fields{
			// "\nindex"		:	k,
			// "\nName"		: 	v.PropertyName,
			// "\nInfo"		:	v.Metadata,
			//}).Infof("\nHere's the device property")
			print(
				"\nindex\t", k,
				"\nName\t", v.PropertyName,
				"\nInfo\t", v.Metadata)
		}

		for k, v := range item.Status.Twins {
			//logrus.WithFields(logrus.Fields{
			// "\nindex"		:	k,
			// "  Name"		:	v.PropertyName,
			// "  type"		:	v.Reported.Metadata["0"],
			//}).Infof("\nHere's the device twins")
			print(
				"\n", k,
				"\t", v.PropertyName,
				"\t", v.Reported.Metadata["0"])
		}
		NewDevice := CloudDeviceMeta{
			ID:          item.Name,
			Name:        item.Name,
			Model:       item.Spec.DeviceModelRef.Name,
			Description: description,
			How:         how,
			Where:       where,
			Node:        item.Spec.NodeSelector.NodeSelectorTerms[0].MatchExpressions[0].Values[0],
		}
		result, err := e.Insert(NewDevice)
		if err != nil {
			print("Insert Error, err:", err)
		}
		print("\nInsert Result is :", result)
	}
	//print("Request err is :", err.Error())
	//DeviceList, err := crdClient.Devices("default").List(context.Background(), v1.ListOptions{})
	//print("The return value is ", DeviceList)
	//print("The result is :", err)
	//if err != nil {
	// print("Crdclient get device list failed. err:", err)
	//}
	//for item := range DeviceList.Items {
	//print(item)
	//}
}

//func PullDeviceList(){
//
//}
