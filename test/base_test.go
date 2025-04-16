/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="base_test.go">
*   Copyright (c) 2025 Aspose.Cells Cloud
* </copyright>
* <summary>
*   Permission is hereby granted, free of charge, to any person obtaining a copy
*  of this software and associated documentation files (the "Software"), to deal
*  in the Software without restriction, including without limitation the rights
*  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  copies of the Software, and to permit persons to whom the Software is
*  furnished to do so, subject to the following conditions:
* 
*  The above copyright notice and this permission notice shall be included in all
*  copies or substantial portions of the Software.
* 
*  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*  SOFTWARE.
* </summary> 
-------------------------------------------------------------------------------------------------------------------- **/

package asposecellscloudtest

import (
	"os"
	"strconv"
	. "asposecellscloud"
)

var BaseTestInstance *BaseTest

type BaseTest struct {
	remoteFolder        string
	localTestDataFolder string
	CellsApi            *CellsApiService
	TestNumber          int
}

func (bt *BaseTest) UploadFile(name string) (err error) {
	args := new(UploadFileRequest)
	args.Path = GetBaseTest().remoteFolder + "/" + name
	args.UploadFiles = make(map[string]string)
	args.UploadFiles[name] = "TestData\\" + name
	_, _, err = GetBaseTest().CellsApi.UploadFile( args)
	return err
}
func (bt *BaseTest) UploadFileToOtherStorage(name string, storageName string) (err error) {
	args := new(UploadFileRequest)
	args.StorageName = storageName
	args.Path = GetBaseTest().remoteFolder + "/" + name
	args.UploadFiles = make(map[string]string)
	args.UploadFiles[name] = "TestData\\" + name

	_, _, err = GetBaseTest().CellsApi.UploadFile( args)
	return err
}

func (bt *BaseTest) FromBealoonToString(v bool) string {
	return strconv.FormatBool(v)
}

func (bt *BaseTest) FromIntToString(v int) string {
	return strconv.FormatInt(int64(v), 10)
}

func (bt *BaseTest) GetTestNumber() int {
	bt.TestNumber++
	return bt.TestNumber
}

func NewBaseTest() *BaseTest {
	bt := &BaseTest{
		remoteFolder:        "GoTest",
		localTestDataFolder: "TestData/",
		TestNumber:          0,
		// Get Client Secret and Client Id from https://aspose.cloud
		CellsApi:      NewCellsApiService(os.Getenv("CellsCloudTestClientId"), os.Getenv("CellsCloudTestClientSecret"), os.Getenv("CellsCloudTestApiBaseUrl"), "v3.0"),
	}
	return bt
}

func GetBaseTest() *BaseTest {
	if BaseTestInstance == nil {
		BaseTestInstance = NewBaseTest()
	}
	return BaseTestInstance
}
