/**
 * Created by huyi on 2017/7/4.
 */
var bcOps = require('./../blockchainOps');

module.exports.measuresAddOne = function(req, res){
    console.log("getting into CreateIOTInfo............")
    insertRow("CreateIOTInfo",req, res);
};

module.exports.logisticsAddOne = function(req, res){
    insertRow("CreateLogisticsInfo",req, res);
};



function insertRow(funcName, req, res){
    var today = new Date();

    var reqJSONString = JSON.stringify(req.body) ;
    var batch = req.body.shipment_id || req.body.batch;

    if (batch == undefined || batch == null || batch.length == 0) {
        res
            .status(400)
            .json({message:"Required key item 'batch' missing from request."});
        return;
    }


    if (reqJSONString.length > 0) {
        var bodystring = "-------->>>>" +  reqJSONString +"<<<<-------"
        console.log(bodystring );

        bcOps.invoke(funcName, today.toString(), reqJSONString, batch);

        res
            .status(201)
            .json(req.body);

    } else {
        res
            .status(400)
            .json({message:"Required data missing from body."});
    }
}


module.exports.measuresList = function(req, res){
    SelectRows("ListIOTInfo",req, res);
};

module.exports.logisticsList = function(req, res){
    console.log("getting into ListLogisticsInfo............")
    SelectRows("ListLogisticsInfo", req, res);
};


function SelectRows(funcName, req, res) {
    var args =[];
    console.log("getting into SelectRows............");
    console.log("funcName--->",funcName);
    console.log("req.params --->",req.params);

    if (funcName == "ListIOTInfo"){
        args.push("");
        args.push(req.params.batch||"");
    } else if (funcName == "ListLogisticsInfo") {
        args.push('');
        args.push('');

    } else if (funcName == "measuresRetrieveOne") {
        args.push("");
        args.push(req.params.recordId);
    } else if (funcName == "measuresRetrieveAll") {
        args.push("");
        args.push("");
    } else return;

    bcOps.query( funcName, args , function(result) {
            res.send(result);
    });
}

module.exports.measuresRetrieveOne = function(req, res){
    SelectRows("measuresRetrieveOne",req, res);
};

module.exports.measuresRetrieveAll = function(req, res){
    SelectRows("measuresRetrieveAll",req, res);
};