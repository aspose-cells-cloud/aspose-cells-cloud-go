/**
 *
 * Copyright (c) 2020 Aspose.Cells Cloud
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */
package asposecellscloud

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var BaseTestInstance *BaseTest

type BaseTest struct {
	remoteFolder        string
	localTestDataFolder string
	CellsAPI            *CellsApiService
	TestNumber          int
}

func (bt *BaseTest) UploadFile(name string) (err error) {
	args := new(UploadFileOpts)
	args.Path = GetBaseTest().remoteFolder + "/" + name
	file, err := os.Open(bt.localTestDataFolder + "/" + name)
	if err != nil {
		return err
	}

	_, _, err = GetBaseTest().CellsAPI.UploadFile(file, args)
	return err
}

func (bt *BaseTest) GetTestNumber() int {
	bt.TestNumber++
	return bt.TestNumber
}

func NewBaseTest() *BaseTest {
	bt := &BaseTest{
		remoteFolder:        "GoTest",
		localTestDataFolder: "../TestData/",
		TestNumber:          0,
		// Get App key and App SID from https://aspose.cloud
		CellsAPI: NewCellsApiService("91A2FD07-BBA1-4B32-9112-ABFB1FE8AEBD", "0fbf678c5ecabdb5caca48452a736dd0", "https://api-qa.aspose.cloud", "v3.0"),
	}
	return bt
}

func GetBaseTest() *BaseTest {
	if BaseTestInstance == nil {
		BaseTestInstance = NewBaseTest()
	}
	return BaseTestInstance
}
func GetMyDoc() string {
	return "myDocument.xlsx"
}
func GetImportData() string {
	return "TestImportData.xlsx"
}
func GetBook1() string {
	return "Book1.xlsx"
}

func GetWorkbook() []byte {
	file, err := os.Open("../TestData/Book1.xlsx")
	if err != nil {
		return nil
	}
	fbs, _ := ioutil.ReadAll(file)
	return fbs
}

func GetPng() string {
	f, err := ioutil.ReadFile("../TestData/WaterMark.png")
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f)
}
func GetNewXLSXFile() string {
	now := time.Now()
	return "New" + now.Format("20060102150405") + ".xlsx"
}

func GetPivTestFile() string {
	return "TestCase.xlsx"
}

func GetReportDataXml() string {
	return "ReportData.xml"
}

func GetWaterMarkPng() string {
	return "WaterMark.png"
}

func GetOLEDoc() string {
	return "OLEDoc.docx"
}

func GetWordJpg() string {
	return "word.jpg"
}

func GetTestImportDataCSV() string {
	return "TestImportDataCSV.csv"
}

func GetSheet1() string {
	return "Sheet1"
}

func GetSheet2() string {
	return "Sheet2"
}

func GetSheet3() string {
	return "Sheet3"
}

func GetSheet4() string {
	return "Sheet4"
}

func GetSheet5() string {
	return "Sheet5"
}

func GetSheet6() string {
	return "Sheet6"
}

func GetSheet7() string {
	return "Sheet7"
}

func ToInt32(index int) int32 {
	return int32(index)
}

func GetRange() string {
	return "A1:C10"
}
