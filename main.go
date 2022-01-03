package main

import (
	sap_api_caller "sap-api-integrations-production-order-confirmation-reads/SAP_API_Caller"
	"sap-api-integrations-production-order-confirmation-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Production_Order_Confirmation_Conf_By_OrderID_Seq_Op_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {

		accepter = []string{
			"ConfByOrderID", "MaterialMovements", "BatchCharacteristic",
			"ConfByOrderIDConfGroup", "ConfByOrderIDSeqOp",
		}
	}

	caller.AsyncGetProductionOrderConfirmation(
		inoutSDC.ProductionOrderConfirmation.OrderID,
		inoutSDC.ProductionOrderConfirmation.MaterialMovements.Batch,
		inoutSDC.ProductionOrderConfirmation.ConfirmationGroup,
		inoutSDC.ProductionOrderConfirmation.Sequence,
		inoutSDC.ProductionOrderConfirmation.OrderOperation,
		accepter,
	)
}
