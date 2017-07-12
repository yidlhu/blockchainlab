

package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"chaincodes/BCDEMO/db"
	"chaincodes/BCDEMO/model"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
	logging "github.com/op/go-logging"
)

var logger = logging.MustGetLogger("BCDEMO")

type Chaincode struct {
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting  Chaincode: %s", err)
	}
}

//template header end

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "init" {
		if len(args) > 1 {
			return nil, errors.New("[Init][init]Incorrect number of arguments. Expecting 0 or 1")
		}
		return t.init(stub, args)
	}
	return nil, errors.New("Received unknown function Init")
}

func (t *Chaincode) init(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//sub.f init start
	logger.Debug("[Init] start.")

	db.CreateTableIOTInfo(stub)

	db.CreateTableLogisticsInfo(stub)

	logger.Debug("[Init] end.")
	return nil, nil
	//sub.f init end
}

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateIOTInfo" {
		if len(args) != 3 {
			return nil, errors.New("[Invoke][CreateIOTInfo]Incorrect number of arguments. Expecting 3")
		}
		return t.CreateIOTInfo(stub, args)

	} else if function == "CreateLogisticsInfo" {
		if len(args) != 3 {
			return nil, errors.New("[Invoke][CreateIOTInfo]Incorrect number of arguments. Expecting 3")
		}
		return t.CreateLogisticsInfo(stub, args)
	}

	return nil, errors.New("Received unknown function Invoke")
}

func (t *Chaincode) CreateLogisticsInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	CURRENT_TIME := args[0]
	logger.Debugf("[createLogisticsInfo] invoke time: %s", CURRENT_TIME)

	logger.Debug("CreateLogisticsInfo args==================>",args, len(args), " |args[0]| ",args[0]," |args[1]| ",args[1]," |args[2]| ",args[2])

	tableModel := new(model.LogisticsInfo)
	tableModel.CREATE_TIME = CURRENT_TIME
	tableModel.LOGISTICS_JSON = args[1]
	tableModel.BATCH = args[2]

	_, err := db.InsertLogisticsInfo(stub, tableModel)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *Chaincode) CreateIOTInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	CURRENT_TIME := args[0]
	logger.Debugf("[createIOTInfo] invoke time: %s", CURRENT_TIME)

	tableModel := new(model.IOTInfo)
	tableModel.CREATE_TIME = CURRENT_TIME
	tableModel.IOT_JSON = args[1]
	tableModel.BATCH = args[2]

	_, err := db.InsertIOTInfo(stub, tableModel)
	if err != nil {
		return nil, err
	}
	return nil, nil
}


func (t *Chaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "GetIOTInfo" {
		if len(args) != 1 {
			return nil, errors.New("[Query][GetIOTInfo]Incorrect number of arguments. Expecting 1")
		}
		return t.GetIOTInfo(stub, args)

	} else if function == "ListIOTInfo" {
		if len(args) != 2 {
			return nil, errors.New("[Query][ListIOTInfo]Incorrect number of arguments. Expecting 2")
		}
		return t.ListIOTInfo(stub, args)

	} else if function == "ListLogisticsInfo" {
		if len(args) != 2 {
			return nil, errors.New("[Query][ListLogisticsInfo]Incorrect number of arguments. Expecting 2")
		}
		return t.ListLogisticsInfo(stub, args)
	}

	return nil, errors.New("Received unknown function Query")
}


func (t *Chaincode) GetIOTInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	IOT_NO := args[0]

	var err error

	rows, err := db.GetIOTInfo(stub, IOT_NO)
	if err != nil {
		return nil, err
	}
	//生成JSON
	IOTInfos, err := db.ParseIOTInfoRows(stub, rows)
	if err != nil {
		return nil, err
	}

	if len(IOTInfos) == 0 {
		err = errors.New("[GetIOTInfo]Data Not Found.")
		return nil, err
	}
	return json.Marshal(IOTInfos[0])

}

func (t *Chaincode) ListIOTInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error

	rows, err := db.SelectIOTInfo(stub, args)
	if err != nil {
		return nil, err
	}
	//生成JSON
	IOTInfos, err := db.ParseIOTInfoRows(stub, rows)
	if err != nil {
		return nil, err
	}

	if len(IOTInfos) == 0 {
		err = errors.New("[ListIOTInfo]Data Not Found.")
		return nil, err
	}
	return json.Marshal(IOTInfos)
}

func (t *Chaincode) ListLogisticsInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var err error

	rows, err := db.SelectLogisticsInfo(stub, args)
	if err != nil {
		return nil, err
	}
	//生成JSON
	LogisticsInfos, err := db.ParseLogisticsInfoRows(stub, rows)
	if err != nil {
		return nil, err
	}

	if len(LogisticsInfos) == 0 {
		err = errors.New("[ListLogisticsInfo]Data Not Found.")
		return nil, err
	}
	return json.Marshal(LogisticsInfos)
}
