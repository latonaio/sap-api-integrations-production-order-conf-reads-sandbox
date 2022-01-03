package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-production-order-confirmation-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetProductionOrderConfirmation(orderID, batch, confirmationGroup, sequence, orderOperation string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ConfByOrderID":
			func() {
				c.ConfByOrderID(orderID)
				wg.Done()
			}()
		case "MaterialMovements":
			func() {
				c.MaterialMovements(batch)
				wg.Done()
			}()
		case "BatchCharacteristic":
			func() {
				c.BatchCharacteristic(batch)
				wg.Done()
			}()
		case "ConfByOrderIDConfGroup":
			func() {
				c.ConfByOrderIDConfGroup(orderID, confirmationGroup)
				wg.Done()
			}()
		case "ConfByOrderIDSeqOp":
			func() {
				c.ConfByOrderIDSeqOp(orderID, sequence, orderOperation)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) ConfByOrderID(orderID string) {
	confbyOrderIDData, err := c.callProductionOrderConfirmationSrvAPIRequirementConfByOrderID("ProdnOrdConf2", orderID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(confbyOrderIDData)

	materialMovementsData, err := c.callToMaterialMovements(confbyOrderIDData[0].ToMaterialMovements)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialMovementsData)

	batchCharacteristicData, err := c.callToBatchCharacteristic(materialMovementsData[0].ToBatchCharacteristic)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(batchCharacteristicData)

}

func (c *SAPAPICaller) callProductionOrderConfirmationSrvAPIRequirementConfByOrderID(api, orderID string) ([]sap_api_output_formatter.Confirmation, error) {
	url := strings.Join([]string{c.baseURL, "API_PROD_ORDER_CONFIRMATION_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithConfByOrderID(req, orderID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToConfirmation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToMaterialMovements(url string) ([]sap_api_output_formatter.ToMaterialMovements, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToMaterialMovements(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBatchCharacteristic(url string) ([]sap_api_output_formatter.ToBatchCharacteristic, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBatchCharacteristic(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) MaterialMovements(batch string) {
	materialMovementsData, err := c.callProductionOrderConfirmationSrvAPIRequirementMaterialMovements("ProdnOrdConfMatlDocItm", batch)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialMovementsData)

	batchCharacteristicData, err := c.callToBatchCharacteristic2(materialMovementsData[0].ToBatchCharacteristic)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(batchCharacteristicData)

}

func (c *SAPAPICaller) callProductionOrderConfirmationSrvAPIRequirementMaterialMovements(api, batch string) ([]sap_api_output_formatter.MaterialMovements, error) {
	url := strings.Join([]string{c.baseURL, "API_PROD_ORDER_CONFIRMATION_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithMaterialMovements(req, batch)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToMaterialMovements(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToBatchCharacteristic2(url string) ([]sap_api_output_formatter.ToBatchCharacteristic, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToBatchCharacteristic(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) BatchCharacteristic(batch string) {
	batchCharacteristicData, err := c.callProductionOrderConfirmationSrvAPIRequirementBatchCharacteristic("ProdnOrderConfBatchCharc", batch)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(batchCharacteristicData)

}

func (c *SAPAPICaller) callProductionOrderConfirmationSrvAPIRequirementBatchCharacteristic(api, batch string) ([]sap_api_output_formatter.BatchCharacteristic, error) {
	url := strings.Join([]string{c.baseURL, "API_PROD_ORDER_CONFIRMATION_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithBatchCharacteristic(req, batch)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToBatchCharacteristic(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ConfByOrderIDConfGroup(orderID, confirmationGroup string) {
	confbyOrderIDConfGroupData, err := c.callProductionOrderConfirmationSrvAPIRequirementConfByOrderIDConfGroup("ProdnOrdConf2", orderID, confirmationGroup)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(confbyOrderIDConfGroupData)

	materialMovementsData, err := c.callToMaterialMovements(confbyOrderIDConfGroupData[0].ToMaterialMovements)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialMovementsData)

	batchCharacteristicData, err := c.callToBatchCharacteristic(materialMovementsData[0].ToBatchCharacteristic)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(batchCharacteristicData)

}

func (c *SAPAPICaller) callProductionOrderConfirmationSrvAPIRequirementConfByOrderIDConfGroup(api, orderID, confirmationGroup string) ([]sap_api_output_formatter.Confirmation, error) {
	url := strings.Join([]string{c.baseURL, "API_PROD_ORDER_CONFIRMATION_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithConfByOrderIDConfGroup(req, orderID, confirmationGroup)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToConfirmation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ConfByOrderIDSeqOp(orderID, sequence, orderOperation string) {
	confbyOrderIDSeqOpData, err := c.callProductionOrderConfirmationSrvAPIRequirementConfByOrderIDSeqOp("ProdnOrdConf2", orderID, sequence, orderOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(confbyOrderIDSeqOpData)

	materialMovementsData, err := c.callToMaterialMovements(confbyOrderIDSeqOpData[0].ToMaterialMovements)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(materialMovementsData)

	batchCharacteristicData, err := c.callToBatchCharacteristic(materialMovementsData[0].ToBatchCharacteristic)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(batchCharacteristicData)

}

func (c *SAPAPICaller) callProductionOrderConfirmationSrvAPIRequirementConfByOrderIDSeqOp(api, orderID, sequence, orderOperation string) ([]sap_api_output_formatter.Confirmation, error) {
	url := strings.Join([]string{c.baseURL, "API_PROD_ORDER_CONFIRMATION_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithConfByOrderIDSeqOp(req, orderID, sequence, orderOperation)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToConfirmation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithConfByOrderID(req *http.Request, orderID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("OrderID eq '%s'", orderID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithMaterialMovements(req *http.Request, batch string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Batch eq '%s'", batch))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithBatchCharacteristic(req *http.Request, batch string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Batch eq '%s'", batch))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithConfByOrderIDConfGroup(req *http.Request, orderID, confirmationGroup string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("OrderID eq '%s' and ConfirmationGroup eq '%s'", orderID, confirmationGroup))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithConfByOrderIDSeqOp(req *http.Request, orderID, sequence, orderOperation string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("OrderID eq '%s' and Sequence eq '%s' and OrderOperation eq '%s'", orderID, sequence, orderOperation))
	req.URL.RawQuery = params.Encode()
}
