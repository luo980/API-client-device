package misc

import (
	_ "github.com/mattn/go-sqlite3"
	"os"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

func EngineInit() (*xorm.Engine, error) {
	//master := "master"
	//slave1 := "slave1"
	//slave2 := "slave2"
	//
	//deviceDriver := "Device"
	//
	//deviceSource := []string{master, slave1, slave2}
	//engineGroup, err := xorm.NewEngineGroup(deviceDriver, deviceSource)

	//masterEngine, err := xorm.NewEngine(deviceDriver, master)
	//slave1Engine, err := xorm.NewEngine(deviceDriver, slave1)
	//slave2Engine, err := xorm.NewEngine(deviceDriver, slave2)
	//engineGroup, err := xorm.NewEngineGroup(masterEngine, []*Engine{slave1Engine, slave2Engine})

	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	f, err := os.Create("sql.log")
	if err != nil {
		print("open file error:", err)
	}
	engine.SetLogger(log.Logger(log.NewSimpleLogger(f)))
	return engine, err
}

func EngineGroupInit() (*xorm.EngineGroup, error) {
	master := "./master.db"
	slave1 := "./slave1.db"
	slave2 := "./slave2.db"

	deviceDriver := "sqlite3"

	deviceSource := []string{master, slave1, slave2}
	engineGroup, err := xorm.NewEngineGroup(deviceDriver, deviceSource)

	//masterEngine, err := xorm.NewEngine(deviceDriver, master)
	//slave1Engine, err := xorm.NewEngine(deviceDriver, slave1)
	//slave2Engine, err := xorm.NewEngine(deviceDriver, slave2)
	//engineGroup, err := xorm.NewEngineGroup(masterEngine, []*Engine{slave1Engine, slave2Engine})

	//engine, err := xorm.NewEngine("sqlite3", "./test.db")

	return engineGroup, err
}

func InsertRecord(e *xorm.Engine, i interface{}) {
	e.Insert(&i)
}

func TableCreate(e *xorm.Engine, name struct{}) bool {
	isexist, err := e.IsTableExist(name)
	if err != nil {
		print("Engine got errors.")
	}
	if isexist {
		return true
	} else {
		err := e.CreateTables(name)
		if err != nil {
			print("Create table failed, err:", err)
			return false
		} else {
			return true
		}
	}
}
