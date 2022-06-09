package sap_core

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"

	"github.com/google/go-querystring/query"
)

type GetDataProductByPrincipalRequestToSAP struct {
	StartNo   int    `url:"StartNo"`
	EndNo     int    `url:"EndNo"`
	Principal string `url:"Principal"`
	Customer  string `url:"Customer"`
	ZZILEV1   string `url:"ZZILEV1"`
	ZZILEV2   string `url:"ZZILEV2"`
	ZZILEV3   string `url:"ZZILEV3"`
	ZZILEV4   string `url:"ZZILEV4"`
	Search    string `url:"Search"`
}

type SAPXMLDataProductByPrincipalResponseFromSAP struct {
	XMLName xml.Name `xml:"string"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
}

type SAPDataProductByPrincipalResponseToService struct {
	Ims1Desc           string `json:"IMS1DESC"`
	Ims2Desc           string `json:"IMS2DESC"`
	Ims3Desc           string `json:"IMS3DESC"`
	Ims4Desc           string `json:"IMS4DESC"`
	Principal          string `json:"Principal"`
	PrincipalName      string `json:"PrincipalName"`
	ProductCategory    string `json:"ProductCategory"`
	ProductCode        string `json:"ProductCode"`
	ProductDescription string `json:"ProductDescription"`
	TotalRow           int64  `json:"TotalRow"`
	Uom                string `json:"Uom"`
	Zzilev1            string `json:"Zzilev1"`
	Zzilev2            string `json:"Zzilev2"`
	Zzilev3            string `json:"Zzilev3"`
	Zzilev4            string `json:"Zzilev4"`
	RowNumber          int64  `json:"row_number"`
	EDStatus           string `json:"EDStatus"`
	StartDate          string `json:"StartDate"`
	EndDate            string `json:"EndDate"`
	ValidFor           string `json:"ValidFor"`
}

func GetProductByPrincipal(reqparam GetDataProductByPrincipalRequestToSAP) (
	[]SAPDataProductByPrincipalResponseToService, string, string, int) {

	endpoint := GetSAPCoreBaseURL() + "/GetDataProduct_V2"
	query, _ := query.Values(reqparam)

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(query.Encode()))
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, err.Error(), fileLocation, fileLine
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(req)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, err.Error(), fileLocation, fileLine
	}
	defer response.Body.Close()

	srv, err := ioutil.ReadAll(response.Body)
	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, err.Error(), fileLocation, fileLine
	}

	if err != nil {
		_, fileLocation, fileLine, _ := runtime.Caller(0)
		return nil, err.Error(), fileLocation, fileLine
	}

	// get xml data from SAP
	sapXMLDataProductByPrincipalResponseFromSAP := SAPXMLDataProductByPrincipalResponseFromSAP{}
	xml.Unmarshal(srv, &sapXMLDataProductByPrincipalResponseFromSAP)

	// convert xml data to json
	resMap := []SAPDataProductByPrincipalResponseToService{}
	json.Unmarshal([]byte(sapXMLDataProductByPrincipalResponseFromSAP.Text), &resMap)
	return resMap, "", "", 0
}
