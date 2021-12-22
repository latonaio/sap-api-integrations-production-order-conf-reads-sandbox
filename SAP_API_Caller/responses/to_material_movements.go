package responses

type ToMaterialMovements struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ConfirmationGroup          string      `json:"ConfirmationGroup"`
			ConfirmationCount          string      `json:"ConfirmationCount"`
			MaterialDocument           string      `json:"MaterialDocument"`
			MaterialDocumentItem       string      `json:"MaterialDocumentItem"`
			MaterialDocumentYear       string      `json:"MaterialDocumentYear"`
			OrderType                  string      `json:"OrderType"`
			OrderID                    string      `json:"OrderID"`
			OrderItem                  string      `json:"OrderItem"`
			ManufacturingOrderCategory string      `json:"ManufacturingOrderCategory"`
			Material                   string      `json:"Material"`
			Plant                      string      `json:"Plant"`
			Reservation                string      `json:"Reservation"`
			ReservationItem            string      `json:"ReservationItem"`
			StorageLocation            string      `json:"StorageLocation"`
			ProductionSupplyArea       string      `json:"ProductionSupplyArea"`
			Batch                      string      `json:"Batch"`
			InventoryValuationType     string      `json:"InventoryValuationType"`
			GoodsMovementType          string      `json:"GoodsMovementType"`
			GoodsMovementRefDocType    string      `json:"GoodsMovementRefDocType"`
			InventoryUsabilityCode     string      `json:"InventoryUsabilityCode"`
			InventorySpecialStockType  string      `json:"InventorySpecialStockType"`
			SalesOrder                 string      `json:"SalesOrder"`
			SalesOrderItem             string      `json:"SalesOrderItem"`
			WBSElementExternalID       string      `json:"WBSElementExternalID"`
			Supplier                   string      `json:"Supplier"`
			Customer                   string      `json:"Customer"`
			ReservationIsFinallyIssued bool        `json:"ReservationIsFinallyIssued"`
			IsCompletelyDelivered      bool        `json:"IsCompletelyDelivered"`
			ShelfLifeExpirationDate    string      `json:"ShelfLifeExpirationDate"`
			ManufactureDate            string      `json:"ManufactureDate"`
			StorageType                string      `json:"StorageType"`
			StorageBin                 string      `json:"StorageBin"`
			MaterialDocumentItemText   string      `json:"MaterialDocumentItemText"`
			EntryUnit                  string      `json:"EntryUnit"`
			QuantityInEntryUnit        string      `json:"QuantityInEntryUnit"`
			ToBatchCharacteristic struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ProdnOrderConfBatchCharc"`
		} `json:"results"`
	} `json:"d"`
}
