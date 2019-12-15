package main

import (
	"fmt"
	docxt "github.com/carmel/go-docx-templates"
)

type TestStruct struct {
	FileName string
	Items    []TestItemStruct
}

type TestItemStruct struct {
	Column1  string
	Column2  string
	SubItems []TestItemStruct2
}

type TestItemStruct2 struct {
	Column1 string
	Column4 string
}

func main() {
	template, err := docxt.OpenTemplate("./example.docx")
	if err != nil {
		fmt.Println(err)
		return
	}
	test := new(TestStruct)
	test.FileName = "example.docx"
	test.Items = []TestItemStruct{
		TestItemStruct{"1", "2", []TestItemStruct2{TestItemStruct2{"你好", "4"}, TestItemStruct2{"5", "中国"}}},
		TestItemStruct{"3", "4", []TestItemStruct2{TestItemStruct2{"7", "8"}, TestItemStruct2{"9", "10"}}},
	}

	// test := map[string]interface{}{
	// 	"FileName": "example.docx",
	// 	"Items": []map[string]interface{}{
	// 		map[string]interface{}{
	// 			"Column1": "1",
	// 			"Column2": "2",
	// 			"SubItems": []map[string]interface{}{
	// 				map[string]interface{}{
	// 					"Column1": "你好",
	// 					"Column4": "4",
	// 				},
	// 				map[string]interface{}{
	// 					"Column1": "5",
	// 					"Column4": "中国",
	// 				},
	// 			},
	// 		},
	// 		map[string]interface{}{
	// 			"Column1": "3",
	// 			"Column2": "4",
	// 			"SubItems": []map[string]interface{}{
	// 				map[string]interface{}{
	// 					"Column1": "7",
	// 					"Column4": "8",
	// 				},
	// 				map[string]interface{}{
	// 					"Column1": "9",
	// 					"Column4": "10",
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	if err := template.RenderTemplate(test); err != nil {
		fmt.Println(err)
		return
	}
	if err := template.Save("result.docx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
