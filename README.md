# Intro | 简介

[![Go Report Card](https://goreportcard.com/badge/github.com/szyhf/go-excel)](https://goreportcard.com/report/github.com/szyhf/go-excel)
![cover.run go](https://cover.run/go/github.com/szyhf/go-excel.svg?tag=golang-1.10)
[![GoDoc](https://godoc.org/github.com/szyhf/go-excel?status.png)](https://godoc.org/github.com/szyhf/go-excel)

Expect to create a reader library to read relate-db-like excel easily.
Just like read a config.

> This library can read all xlsx file correctly from our project now.

```shell
go get github.com/szyhf/go-excel
```

## Example | 用例

Here is a simple example.

Assume you have a xlsx file like below:

|ID|NameOf|Age|Slice|UnmarshalString|
|-|-|-|-|-|
|1|Andy|1|1\|2|{"Foo":"Andy"}|
|2|Leo|2|2\|3\|4|{"Foo":"Leo"}|
|3|Ben|3|3\|4\|5\|6|{"Foo":"Ben"}|
|4|Ming|4|1|{"Foo":"Ming"}|

+ the first row is the title row.
+ other row is the data row.

> All examples list in [https://godoc.org/github.com/szyhf/go-excel#pkg-examples](https://godoc.org/github.com/szyhf/go-excel#pkg-examples).

```go
// defined a struct
type Standard struct {
	// use field name as default column name
	ID      int
	// column means to map the column name
	Name    string `xlsx:"column(NameOf)"`
	// you can map a column into more than one field
	NamePtr *string `xlsx:"column(NameOf)"`
	// omit `column` if only want to map to column name, it's equal to `column(AgeOf)`
	Age     int     `xlsx:"AgeOf"`
	// split means to split the string into slice by the `|`
	Slice   []int `xlsx:"split(|)"`
	// *Temp implement the `encoding.BinaryUnmarshaler`
	Temp    *Temp `xlsx:"column(UnmarshalString)"`
	// use '-' to ignore.
	Ignored string `xlsx:"-"`
}

// func (this Standard) GetXLSXSheetName() string {
// 	return "Some other sheet name if need"
// }

type Temp struct {
	Foo string
}

// self define a unmarshal interface to unmarshal string.
func (this *Temp) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, this)
}

func simpleUsage() {
	// will assume the sheet name as "Standard" from the struct name.
	var stdList []Standard
	err := excel.UnmarshalXLSX("./testdata/simple.xlsx", &stdList)
	if err != nil {
		panic(err)
	}
}

func defaultUsage(){
	conn := excel.NewConnecter()
	err := conn.Open("./testdata/simple.xlsx")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Generate an new reader of a sheet
	// sheetNamer: if sheetNamer is string, will use sheet as sheet name.
	//             if sheetNamer is int, will i'th sheet in the workbook, be careful the hidden sheet is counted. i ∈ [1,+inf]
	//             if sheetNamer is a object implements `GetXLSXSheetName()string`, the return value will be used.
	//             otherwise, will use sheetNamer as struct and reflect for it's name.
	// 			   if sheetNamer is a slice, the type of element will be used to infer like before.
	rd, err := conn.NewReader(stdSheetName)
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	for rd.Next() {
		var s Standard
		// Read a row into a struct.
		err:=rd.Read(&s)
		if err!=nil{
			panic(err)
		}
		fmt.Printf("%+v",s)
	}

	// Read all is also supported.
	// var stdList []Standard
	// err = rd.ReadAll(&stdList)
	// if err != nil {
	//   panic(err)
	//	 return
	// }
	// fmt.Printf("%+v",stdList)

	// map with string key is support, too.
	// if value is not string
	// will try to unmarshal to target type
	// but will skip if unmarshal failed.

	// var stdMapList []map[string]string
	// err = rd.ReadAll(&stdMapList)
	// if err != nil {
	//   panic(err)
	//	 return
	// }
	// fmt.Printf("%+v",stdList)

	// Using binary instead of file.
	// xlsxData, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// conn := excel.NewConnecter()
	// err = conn.OpenBinary(xlsxData)
}
```

> See the `simple.xlsx`.`Standard` in `testdata` and code in `./standard_test.go` and `./standard_example_test.go` for details.

## Advance 

The advance usage can make more options.

### Config | 配置

Using a config as "excel.Config":

```go
type Config struct {
	// sheet: if sheet is string, will use sheet as sheet name.
	//        if sheet is a object implements `GetXLSXSheetName()string`, the return value will be used.
	//        otherwise, will use sheet as struct and reflect for it's name.
	// 		  if sheet is a slice, the type of element will be used to infer like before.
	Sheet interface{}
	// Use the index row as title, every row before title-row will be ignore, default is 0.
	TitleRowIndex int
	// Skip n row after title, default is 0 (not skip), empty row is not counted.
	Skip int
	// Auto prefix to sheet name.
	Prefix string
	// Auto suffix to sheet name.
	Suffix string
}
```

Tips:

+ Empty row will be skipped.
+ Column larger than len(TitleRow) will be skipped.
+ Only empty cell can fill with default value, if a cell can not parse into a field it will return an error.
+ Default value can be unmarshal by `encoding.BinaryUnmarshaler`, too.
+ If no title row privoded, the default column name in exce like `'A', 'B', 'C', 'D' ......, 'XFC', 'XFD'` can be used as column name by 26-number-system.

For more details can see the code in `./test/advance_test.go` and file in `simple.xlsx`.`Advance.suffx` sheet.

## XLSX Tag | 标签使用

### column

Map to field name in title row, by default will use the field name.

### default

Set default value when no value is filled in excel cell, by default is 0 or "".

### split

Split a string and convert them to a slice, it won't work if not set.

### nil

Will not skip scan value in the cell equals to this 'nil value'

### req

Will return error if clomun title not exist in excel.

