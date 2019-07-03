package excel_test

import (
	"fmt"
	"testing"

	excel "github.com/n1kx/go-excel"
)

// defined a struct
type SimpleStruct struct {
	FirstName string `xlsx:"first_name"`
	LastName  string `xlsx:"last_name"`
	Age       int    `xlsx:"age"`
	FullName  string `xlsx:"full_name"`
	Age2      string `xlsx:"age2"`
}

//ResultMatching is the database table representative
type ResultMatching struct {
	Periode              string `xlsx:"periode"`
	StoreID              string `xlsx:"store_id"`
	ProductName          string `xlsx:"product_name"`
	SkuStockist          string `xlsx:"sku_stockist"`
	PriceStockist        string `xlsx:"price_stockist"`
	Stock                string `xlsx:"stock"`
	MovingAverage        string `xlsx:"moving_average"`
	DoiPercentage        string `xlsx:"doi_percentage"`
	Status               string `xlsx:"status"`
	BundlingQty          string `xlsx:"bundling_qty"`
	BundlingPrice        string `xlsx:"bundling_price"`
	UOM                  string `xlsx:"uom"`
	DrugAttrStrengthApt  string `xlsx:"drug_attr_strength_apt"`
	Value                string `xlsx:"value"`
	Matching             string `xlsx:"matching"`
	SkuOriginal          string `xlsx:"sku_original"`
	Kode                 string `xlsx:"kode"`
	SkuOriginalValidasi2 string `xlsx:"sku_original_validasi_2"`
	SkuOriginalFinal     string `xlsx:"sku_original_final"`
	UOM1                 string `xlsx:"uom1"`
	Qty1                 string `xlsx:"qty1"`
	UOM2                 string `xlsx:"uom2"`
	Qty2                 string `xlsx:"qty2"`
	DrugAttrUomBased     string `xlsx:"drug_attr_uom_based"`
	StockBased           string `xlsx:"stock_based"`
	PriceStockist1       string `xlsx:"price_stockist1"`
}

func Test_Simple(t *testing.T) {
	conn := excel.NewConnecter()

	err := conn.Open("../go-excel/testdata/result_matching.xlsx")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer conn.Close()

	rd, err := conn.NewReader("Sheet1")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer rd.Close()

	for rd.Next() {
		var r ResultMatching
		// Read a row into a struct.
		rd.Read(&r)
		fmt.Println(r)

	}
}
