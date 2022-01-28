package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-order-confirmation-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToConfirmation(raw []byte, l *logger.Logger) ([]Confirmation, error) {
	pm := &responses.Confirmation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Confirmation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	confirmation := make([]Confirmation, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		confirmation = append(confirmation, Confirmation{
			ConfirmationGroup:              data.ConfirmationGroup,
			ConfirmationCount:              data.ConfirmationCount,
			OrderID:                        data.OrderID,
			Sequence:                       data.Sequence,
			OrderOperation:                 data.OrderOperation,
			OrderSuboperation:              data.OrderSuboperation,
			OrderType:                      data.OrderType,
			OrderOperationInternalID:       data.OrderOperationInternalID,
			ConfirmationText:               data.ConfirmationText,
			Language:                       data.Language,
			Material:                       data.Material,
			OrderPlannedTotalQty:           data.OrderPlannedTotalQty,
			ProductionUnit:                 data.ProductionUnit,
			FinalConfirmationType:          data.FinalConfirmationType,
			IsFinalConfirmation:            data.IsFinalConfirmation,
			OpenReservationsIsCleared:      data.OpenReservationsIsCleared,
			IsReversed:                     data.IsReversed,
			IsReversal:                     data.IsReversal,
			APIConfHasNoGoodsMovements:     data.APIConfHasNoGoodsMovements,
			OrderConfirmationRecordType:    data.OrderConfirmationRecordType,
			ConfirmationEntryDate:          data.ConfirmationEntryDate,
			ConfirmationEntryTime:          data.ConfirmationEntryTime,
			EnteredByUser:                  data.EnteredByUser,
			LastChangeDate:                 data.LastChangeDate,
			LastChangedByUser:              data.LastChangedByUser,
			ConfirmationExternalEntryDate:  data.ConfirmationExternalEntryDate,
			ConfirmationExternalEntryTime:  data.ConfirmationExternalEntryTime,
			EnteredByExternalUser:          data.EnteredByExternalUser,
			ExternalSystemConfirmation:     data.ExternalSystemConfirmation,
			Plant:                          data.Plant,
			WorkCenterTypeCode:             data.WorkCenterTypeCode,
			WorkCenter:                     data.WorkCenter,
			CapacityCategoryCode:           data.CapacityCategoryCode,
			CapacityRequirementSplit:       data.CapacityRequirementSplit,
			Personnel:                      data.Personnel,
			TimeRecording:                  data.TimeRecording,
			EmployeeWageType:               data.EmployeeWageType,
			EmployeeWageGroup:              data.EmployeeWageGroup,
			BreakDurationUnit:              data.BreakDurationUnit,
			ConfirmedBreakDuration:         data.ConfirmedBreakDuration,
			EmployeeSuitability:            data.EmployeeSuitability,
			NumberOfEmployees:              data.NumberOfEmployees,
			PostingDate:                    data.PostingDate,
			ConfirmedExecutionStartDate:    data.ConfirmedExecutionStartDate,
			ConfirmedExecutionStartTime:    data.ConfirmedExecutionStartTime,
			ConfirmedSetupEndDate:          data.ConfirmedSetupEndDate,
			ConfirmedSetupEndTime:          data.ConfirmedSetupEndTime,
			ConfirmedProcessingStartDate:   data.ConfirmedProcessingStartDate,
			ConfirmedProcessingStartTime:   data.ConfirmedProcessingStartTime,
			ConfirmedProcessingEndDate:     data.ConfirmedProcessingEndDate,
			ConfirmedProcessingEndTime:     data.ConfirmedProcessingEndTime,
			ConfirmedTeardownStartDate:     data.ConfirmedTeardownStartDate,
			ConfirmedTeardownStartTime:     data.ConfirmedTeardownStartTime,
			ConfirmedExecutionEndDate:      data.ConfirmedExecutionEndDate,
			ConfirmedExecutionEndTime:      data.ConfirmedExecutionEndTime,
			ConfirmationUnit:               data.ConfirmationUnit,
			ConfirmationYieldQuantity:      data.ConfirmationYieldQuantity,
			ConfirmationScrapQuantity:      data.ConfirmationScrapQuantity,
			ConfirmationReworkQuantity:     data.ConfirmationReworkQuantity,
			VarianceReasonCode:             data.VarianceReasonCode,
			OpWorkQuantityUnit1:            data.OpWorkQuantityUnit1,
			OpConfirmedWorkQuantity1:       data.OpConfirmedWorkQuantity1,
			NoFurtherOpWorkQuantity1IsExpd: data.NoFurtherOpWorkQuantity1IsExpd,
			OpWorkQuantityUnit2:            data.OpWorkQuantityUnit2,
			OpConfirmedWorkQuantity2:       data.OpConfirmedWorkQuantity2,
			NoFurtherOpWorkQuantity2IsExpd: data.NoFurtherOpWorkQuantity2IsExpd,
			OpWorkQuantityUnit3:            data.OpWorkQuantityUnit3,
			OpConfirmedWorkQuantity3:       data.OpConfirmedWorkQuantity3,
			NoFurtherOpWorkQuantity3IsExpd: data.NoFurtherOpWorkQuantity3IsExpd,
			OpWorkQuantityUnit4:            data.OpWorkQuantityUnit4,
			OpConfirmedWorkQuantity4:       data.OpConfirmedWorkQuantity4,
			NoFurtherOpWorkQuantity4IsExpd: data.NoFurtherOpWorkQuantity4IsExpd,
			OpWorkQuantityUnit5:            data.OpWorkQuantityUnit5,
			OpConfirmedWorkQuantity5:       data.OpConfirmedWorkQuantity5,
			NoFurtherOpWorkQuantity5IsExpd: data.NoFurtherOpWorkQuantity5IsExpd,
			OpWorkQuantityUnit6:            data.OpWorkQuantityUnit6,
			OpConfirmedWorkQuantity6:       data.OpConfirmedWorkQuantity6,
			NoFurtherOpWorkQuantity6IsExpd: data.NoFurtherOpWorkQuantity6IsExpd,
			BusinessProcessEntryUnit:       data.BusinessProcessEntryUnit,
			BusinessProcessConfirmedQty:    data.BusinessProcessConfirmedQty,
			NoFurtherBusinessProcQtyIsExpd: data.NoFurtherBusinessProcQtyIsExpd,
			ToMaterialMovements:            data.ToMaterialMovements.Deferred.URI,
		})
	}

	return confirmation, nil
}

func ConvertToMaterialMovements(raw []byte, l *logger.Logger) ([]MaterialMovements, error) {
	pm := &responses.MaterialMovements{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to MaterialMovements. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	materialMovements := make([]MaterialMovements, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		materialMovements = append(materialMovements, MaterialMovements{
			ConfirmationGroup:          data.ConfirmationGroup,
			ConfirmationCount:          data.ConfirmationCount,
			MaterialDocument:           data.MaterialDocument,
			MaterialDocumentItem:       data.MaterialDocumentItem,
			MaterialDocumentYear:       data.MaterialDocumentYear,
			OrderType:                  data.OrderType,
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			ManufacturingOrderCategory: data.ManufacturingOrderCategory,
			Material:                   data.Material,
			Plant:                      data.Plant,
			Reservation:                data.Reservation,
			ReservationItem:            data.ReservationItem,
			StorageLocation:            data.StorageLocation,
			ProductionSupplyArea:       data.ProductionSupplyArea,
			Batch:                      data.Batch,
			InventoryValuationType:     data.InventoryValuationType,
			GoodsMovementType:          data.GoodsMovementType,
			GoodsMovementRefDocType:    data.GoodsMovementRefDocType,
			InventoryUsabilityCode:     data.InventoryUsabilityCode,
			InventorySpecialStockType:  data.InventorySpecialStockType,
			SalesOrder:                 data.SalesOrder,
			SalesOrderItem:             data.SalesOrderItem,
			WBSElementExternalID:       data.WBSElementExternalID,
			Supplier:                   data.Supplier,
			Customer:                   data.Customer,
			ReservationIsFinallyIssued: data.ReservationIsFinallyIssued,
			IsCompletelyDelivered:      data.IsCompletelyDelivered,
			ShelfLifeExpirationDate:    data.ShelfLifeExpirationDate,
			ManufactureDate:            data.ManufactureDate,
			StorageType:                data.StorageType,
			StorageBin:                 data.StorageBin,
			MaterialDocumentItemText:   data.MaterialDocumentItemText,
			EntryUnit:                  data.EntryUnit,
			QuantityInEntryUnit:        data.QuantityInEntryUnit,
			ToBatchCharacteristic:      data.ToBatchCharacteristic.Deferred.URI,
		})
	}

	return materialMovements, nil
}

func ConvertToBatchCharacteristic(raw []byte, l *logger.Logger) ([]BatchCharacteristic, error) {
	pm := &responses.BatchCharacteristic{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BatchCharacteristic. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	batchCharacteristic := make([]BatchCharacteristic, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		batchCharacteristic = append(batchCharacteristic, BatchCharacteristic{
			ConfirmationGroup:    data.ConfirmationGroup,
			ConfirmationCount:    data.ConfirmationCount,
			MaterialDocument:     data.MaterialDocument,
			MaterialDocumentItem: data.MaterialDocumentItem,
			MaterialDocumentYear: data.MaterialDocumentYear,
			Plant:                data.Plant,
			Material:             data.Material,
			Batch:                data.Batch,
			CharcInternalID:      data.CharcInternalID,
			Characteristic:       data.Characteristic,
			CharcValue:           data.CharcValue,
		})
	}

	return batchCharacteristic, nil
}

func ConvertToToMaterialMovements(raw []byte, l *logger.Logger) ([]ToMaterialMovements, error) {
	pm := &responses.ToMaterialMovements{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToMaterialMovements. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toMaterialMovements := make([]ToMaterialMovements, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toMaterialMovements = append(toMaterialMovements, ToMaterialMovements{
			ConfirmationGroup:          data.ConfirmationGroup,
			ConfirmationCount:          data.ConfirmationCount,
			MaterialDocument:           data.MaterialDocument,
			MaterialDocumentItem:       data.MaterialDocumentItem,
			MaterialDocumentYear:       data.MaterialDocumentYear,
			OrderType:                  data.OrderType,
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			ManufacturingOrderCategory: data.ManufacturingOrderCategory,
			Material:                   data.Material,
			Plant:                      data.Plant,
			Reservation:                data.Reservation,
			ReservationItem:            data.ReservationItem,
			StorageLocation:            data.StorageLocation,
			ProductionSupplyArea:       data.ProductionSupplyArea,
			Batch:                      data.Batch,
			InventoryValuationType:     data.InventoryValuationType,
			GoodsMovementType:          data.GoodsMovementType,
			GoodsMovementRefDocType:    data.GoodsMovementRefDocType,
			InventoryUsabilityCode:     data.InventoryUsabilityCode,
			InventorySpecialStockType:  data.InventorySpecialStockType,
			SalesOrder:                 data.SalesOrder,
			SalesOrderItem:             data.SalesOrderItem,
			WBSElementExternalID:       data.WBSElementExternalID,
			Supplier:                   data.Supplier,
			Customer:                   data.Customer,
			ReservationIsFinallyIssued: data.ReservationIsFinallyIssued,
			IsCompletelyDelivered:      data.IsCompletelyDelivered,
			ShelfLifeExpirationDate:    data.ShelfLifeExpirationDate,
			ManufactureDate:            data.ManufactureDate,
			StorageType:                data.StorageType,
			StorageBin:                 data.StorageBin,
			MaterialDocumentItemText:   data.MaterialDocumentItemText,
			EntryUnit:                  data.EntryUnit,
			QuantityInEntryUnit:        data.QuantityInEntryUnit,
			ToBatchCharacteristic:      data.ToBatchCharacteristic.Deferred.URI,
		})
	}

	return toMaterialMovements, nil
}

func ConvertToToBatchCharacteristic(raw []byte, l *logger.Logger) ([]ToBatchCharacteristic, error) {
	pm := &responses.ToBatchCharacteristic{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToBatchCharacteristic. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toBatchCharacteristic := make([]ToBatchCharacteristic, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toBatchCharacteristic = append(toBatchCharacteristic, ToBatchCharacteristic{
			ConfirmationGroup:    data.ConfirmationGroup,
			ConfirmationCount:    data.ConfirmationCount,
			MaterialDocument:     data.MaterialDocument,
			MaterialDocumentItem: data.MaterialDocumentItem,
			MaterialDocumentYear: data.MaterialDocumentYear,
			Plant:                data.Plant,
			Material:             data.Material,
			Batch:                data.Batch,
			CharcInternalID:      data.CharcInternalID,
			Characteristic:       data.Characteristic,
			CharcValue:           data.CharcValue,
		})
	}

	return toBatchCharacteristic, nil
}
