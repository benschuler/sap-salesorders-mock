package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type SalesOrdersResponse struct {
	D SalesOrders `json:"d"`
}

type SalesOrderResponse struct {
	D SalesOrder `json:"d"`
}

type SalesOrders struct {
	Results []SalesOrder `json:"results"`
}

type SalesOrderMetadata struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Type string `json:"type"`
	Etag string `json:"etag"`
}

type SalesOrder struct {
	Metadata                       SalesOrderMetadata `json:"__metadata"`
	SalesOrder                     string
	SalesOrderType                 string
	SalesOrganization              string
	DistributionChannel            string
	OrganizationDivision           string
	SalesGroup                     string
	SalesOffice                    string
	SalesDistrict                  string
	SoldToParty                    string
	CreationDate                   string
	CreatedByUser                  string
	LastChangeDate                 string
	SenderBusinessSystemName       string
	ExternalDocumentID             string
	LastChangeDateTime             string
	ExternalDocLastChangeDateTime  string
	PurchaseOrderByCustomer        string
	PurchaseOrderByShipToParty     string
	CustomerPurchaseOrderType      string
	CustomerPurchaseOrderDate      string
	SalesOrderDate                 string
	TotalNetAmount                 string
	TransactionCurrency            string
	SDDocumentReason               string
	PricingDate                    string
	PriceDetnExchangeRate          string
	RequestedDeliveryDate          string
	ShippingCondition              string
	CompleteDeliveryIsDefined      string
	ShippingType                   string
	HeaderBillingBlockReason       string
	DeliveryBlockReason            string
	DeliveryDateTypeRule           string
	IncotermsClassification        string
	IncotermsTransferLocation      string
	IncotermsLocation1             string
	IncotermsLocation2             string
	IncotermsVersion               string
	CustomerPriceGroup             string
	PriceListType                  string
	CustomerPaymentTerms           string
	PaymentMethod                  string
	FixedValueDate                 string
	AssignmentReference            string
	ReferenceSDDocument            string
	ReferenceSDDocumentCategory    string
	AccountingDocExternalReference string
	CustomerAccountAssignmentGroup string
	AccountingExchangeRate         string
	CustomerGroup                  string
	AdditionalCustomerGroup1       string
	AdditionalCustomerGroup2       string
	AdditionalCustomerGroup3       string
	AdditionalCustomerGroup4       string
	AdditionalCustomerGroup5       string
	SlsDocIsRlvtForProofOfDeliv    string
	CustomerTaxClassification1     string
	CustomerTaxClassification2     string
	CustomerTaxClassification3     string
	CustomerTaxClassification4     string
	CustomerTaxClassification5     string
	CustomerTaxClassification6     string
	CustomerTaxClassification7     string
	CustomerTaxClassification8     string
	CustomerTaxClassification9     string
	TaxDepartureCountry            string
	VATRegistrationCountry         string
	SalesOrderApprovalReason       string
	SalesDocApprovalStatus         string
	OverallSDProcessStatus         string
	TotalCreditCheckStatus         string
	OverallTotalDeliveryStatus     string
	OverallSDDocumentRejectionSts  string
	ToItem                         DeferredReference `json:"to_Item"`
	ToPartner                      DeferredReference `json:"to_Partner"`
	ToPaymentPlanItemDetails       DeferredReference `json:"to_PaymentPlanItemDetails"`
	ToPricingElement               DeferredReference `json:"to_PricingElement"`
	ToText                         DeferredReference `json:"to_Text"`
}

type DeferredReference struct {
	Deferred Ref `json:"__deferred"`
}

type Ref struct {
	Uri string `json:"uri"`
}

var orderCol []SalesOrder = []SalesOrder{
	{
		SalesOrder:            "9000000152",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "01",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "7",
		SoldToParty:           "10332",
		CreationDate:          "1/2/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "12/30/2020",
		TotalNetAmount:        "245.83",
		TransactionCurrency:   "USD",
		PricingDate:           "1/2/2021",
		RequestedDeliveryDate: "3/31/2021",
	},
	{
		SalesOrder:            "9000000158",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "01",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1002085",
		CreationDate:          "2/11/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "2/7/2021",
		TotalNetAmount:        "901.36",
		TransactionCurrency:   "USD",
		PricingDate:           "2/11/2021",
		RequestedDeliveryDate: "7/1/2021",
	},
	{
		SalesOrder:            "9000000173",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1004186",
		CreationDate:          "12/29/2020",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "12/29/2020",
		TotalNetAmount:        "1180.86",
		TransactionCurrency:   "USD",
		PricingDate:           "12/29/2020",
		RequestedDeliveryDate: "4/16/2021",
	},
	{
		SalesOrder:            "9000000253",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1001710",
		CreationDate:          "1/20/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "1/19/2021",
		TotalNetAmount:        "175.75",
		TransactionCurrency:   "USD",
		PricingDate:           "1/20/2021",
		RequestedDeliveryDate: "4/21/2021",
	},
	{
		SalesOrder:            "9000000348",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1001710",
		CreationDate:          "3/8/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "3/7/2021",
		TotalNetAmount:        "1576.55",
		TransactionCurrency:   "USD",
		PricingDate:           "3/8/2021",
		RequestedDeliveryDate: "8/8/2021",
	},
	{
		SalesOrder:            "9000000363",
		SalesOrderType:        "ZSMP",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "2",
		SoldToParty:           "1010004",
		CreationDate:          "2/23/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "2/23/2021",
		TotalNetAmount:        "1112.74",
		TransactionCurrency:   "USD",
		PricingDate:           "2/23/2021",
		RequestedDeliveryDate: "5/7/2021",
	},
	{
		SalesOrder:            "9000000364",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "2",
		SoldToParty:           "1010004",
		CreationDate:          "1/6/2021",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "1/5/2021",
		TotalNetAmount:        "1051.28",
		TransactionCurrency:   "USD",
		PricingDate:           "1/6/2021",
		RequestedDeliveryDate: "7/3/2021",
	},
	{
		SalesOrder:            "9000000043",
		SalesOrderType:        "OR",
		SalesOrganization:     "US1100",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "7",
		SoldToParty:           "10332",
		CreationDate:          "12/17/2020",
		CreatedByUser:         "LARRY",
		LastChangeDate:        "",
		SalesOrderDate:        "12/17/2020",
		TotalNetAmount:        "2260.6",
		TransactionCurrency:   "USD",
		PricingDate:           "12/17/2020",
		RequestedDeliveryDate: "2/14/2021",
	},
	{
		SalesOrder:            "9000000232",
		SalesOrderType:        "ZOR",
		SalesOrganization:     "2000",
		DistributionChannel:   "01",
		OrganizationDivision:  "",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1004186",
		CreationDate:          "3/18/2021",
		CreatedByUser:         "SERGEY",
		LastChangeDate:        "",
		SalesOrderDate:        "3/15/2021",
		TotalNetAmount:        "1356.08",
		TransactionCurrency:   "EUR",
		PricingDate:           "3/18/2021",
		RequestedDeliveryDate: "6/10/2021",
	},
	{
		SalesOrder:            "9000000237",
		SalesOrderType:        "ZOR",
		SalesOrganization:     "2000",
		DistributionChannel:   "01",
		OrganizationDivision:  "00",
		SalesGroup:            "",
		SalesOffice:           "",
		SalesDistrict:         "",
		SoldToParty:           "1004186",
		CreationDate:          "2/8/2021",
		CreatedByUser:         "SERGEY",
		LastChangeDate:        "",
		SalesOrderDate:        "2/5/2021",
		TotalNetAmount:        "81.86",
		TransactionCurrency:   "EUR",
		PricingDate:           "2/8/2021",
		RequestedDeliveryDate: "6/15/2021",
	},
}

func SetResponseDefaults(o *SalesOrdersResponse) {
	for i := 0; i < len(o.D.Results); i++ {
		SetOrderDefaults(&o.D.Results[i], i)
	}
}

func SetSingleResponseDefaults(o *SalesOrderResponse) {
	SetOrderDefaults(&o.D, 0)
}

func SetOrderDefaults(o *SalesOrder, i int) {

	if o.SalesOrderType == "" {
		o.SalesOrderType = "OR"
	}
	if o.SalesOrganization == "" {
		o.SalesOrganization = "1710"
	}
	if o.DistributionChannel == "" {
		o.DistributionChannel = "10"
	}

	o.RequestedDeliveryDate = "2021-04-24T11:20:58+00:00"

	if i == 0 {
		o.OverallTotalDeliveryStatus = "SCHEDULED"
	} else {
		o.OverallTotalDeliveryStatus = "FULFILLMENT"
	}

	o.Metadata = SalesOrderMetadata{Id: "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')",
		Uri:  "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')",
		Type: "API_SALES_ORDER_SRV.A_SalesOrderType",
		Etag: "W/\"datetimeoffset'2016-09-02T06%3A15%3A47.1257050Z'\""}

	o.ToItem = DeferredReference{Deferred: Ref{Uri: "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')/to_Item"}}
	o.ToPartner = DeferredReference{Deferred: Ref{Uri: "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')/to_Partner"}}
	o.ToPaymentPlanItemDetails = DeferredReference{Deferred: Ref{Uri: "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')/to_PaymentPlanItemDetails"}}
	o.ToPricingElement = DeferredReference{Deferred: Ref{Uri: "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')/to_PricingElement"}}
	o.ToText = DeferredReference{Deferred: Ref{Uri: "https://sap-orders-mock-h7pi7igbcq-ew.a.run.app/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder('" + strconv.Itoa(i+1) + "')/to_Text"}}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		r := SalesOrdersResponse{D: SalesOrders{Results: orderCol}}
		SetResponseDefaults(&r)
		j, _ := json.Marshal(r)
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &SalesOrder{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		r := SalesOrderResponse{D: *p}
		SetSingleResponseDefaults(&r)
		orderCol = append(orderCol, *p)

		j, _ := json.Marshal(r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}

func main() {
	http.HandleFunc("/sap/opu/odata/sap/API_SALES_ORDER_SRV/A_SalesOrder", orderHandler)

	log.Println("Go!")
	http.ListenAndServe(":8080", nil)
}
