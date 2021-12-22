package responses

type BatchCharacteristic struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ConfirmationGroup    string `json:"ConfirmationGroup"`
			ConfirmationCount    string `json:"ConfirmationCount"`
			MaterialDocument     string `json:"MaterialDocument"`
			MaterialDocumentItem string `json:"MaterialDocumentItem"`
			MaterialDocumentYear string `json:"MaterialDocumentYear"`
			Plant                string `json:"Plant"`
			Material             string `json:"Material"`
			Batch                string `json:"Batch"`
			CharcInternalID      string `json:"CharcInternalID"`
			Characteristic       string `json:"Characteristic"`
			CharcValue           string `json:"CharcValue"`
		} `json:"results"`
	} `json:"d"`
}
