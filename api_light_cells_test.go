/*
 *  Copyright (c) 2021 Aspose.Cells Cloud
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
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
 *
 */

package asposecellscloud

import (
	"fmt"
	"testing"
)

func TestCellsDeleteMetadata(t *testing.T) {
	name := GetBook1()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	_, httpResponse, err := GetBaseTest().LightCellsApi.DeleteMetadata(fileMap, nil)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tDeleteMetadata \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostAssemble(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postAssembleOpts := new(PostAssembleOpts)
	postAssembleOpts.Datasource = "ds"
	postAssembleOpts.Format = "xlsx"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostAssemble(fileMap, postAssembleOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostAssemble \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostClearObjects(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostClearObjectsOpts)
	postOpts.Objecttype = "chart"

	_, httpResponse, err := GetBaseTest().LightCellsApi.PostClearObjects(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostClearObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostExport_chart(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "pdf"
	postOpts.ObjectType = "chart"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostExport_listobject(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "pdf"
	postOpts.ObjectType = "listobject"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}
func TestCellsPostExport_picture(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "png"
	postOpts.ObjectType = "picture"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}
func TestCellsPostExport_shape(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "png"
	postOpts.ObjectType = "shape"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}
func TestCellsPostExport_oleobject(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "png"
	postOpts.ObjectType = "oleobject"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}
func TestCellsPostExport_worksheet(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "pdf"
	postOpts.ObjectType = "worksheet"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostExport_workbook(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "pdf"
	postOpts.ObjectType = "workbook"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostImport(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name

	importIntArrayOption := new(ImportIntArrayOption)
	importIntArrayOption.Data = []int64{1, 2, 3}
	importIntArrayOption.DestinationWorksheet = GetSheet1()
	importIntArrayOption.FirstColumn = 2
	importIntArrayOption.FirstRow = 1
	importIntArrayOption.IsVertical = true
	importIntArrayOption.ImportDataType = "IntArray"
	postOpts := new(PostImportOpts)
	postOpts.ImportOption = &importIntArrayOption

	_, httpResponse, err := GetBaseTest().LightCellsApi.PostImport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostImport \n", GetBaseTest().GetTestNumber())
	}
}
func TestCellsPostMerge(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostMergeOpts)
	postOpts.Format = "pdf"

	_, httpResponse, err := GetBaseTest().LightCellsApi.PostMerge(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostMetadata(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostMetadataOpts)
	cellsDocumentProperty := new(CellsDocumentProperty)
	cellsDocumentProperty.Name = "test"
	cellsDocumentProperty.Value = "test"
	postOpts.DocumentProperties = []CellsDocumentProperty{*cellsDocumentProperty}
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostMetadata(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostMetadata \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostProtect(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostProtectOpts)
	postOpts.Password = "123456"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostProtect(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostProtect \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostSearch(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostSearchOpts)
	postOpts.Text = "1"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostSearch(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostSearch_Sheet(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostSearchOpts)
	postOpts.Text = "1"
	postOpts.Sheetname = "Sheet1"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostSearch(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostSplit(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostSplitOpts)
	postOpts.Format = "png"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostSplit(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostSplit \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostUnlock(t *testing.T) {
	name := GetNeedUnlock()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name

	postOpts := new(PostUnlockOpts)
	postOpts.Password = "123456"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostUnlock(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostUnlock \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostWatermark(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostWatermarkOpts)
	postOpts.Color = "#fff"
	postOpts.Text = "aspose.cells cloud"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostWatermark(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostWatermark \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostCompress(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostCompressOpts)
	postOpts.CompressLevel = 80
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostCompress(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostCompress \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostReplace(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostReplaceOpts)
	postOpts.Text = "1"
	postOpts.NewText = "aspose.cells cloud"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostReplace(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t func TestCellsPostReplace	\n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostReverse(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostReverseOpts)
	postOpts.RotateType = "row"
	postOpts.Format = "pdf"
	_, httpResponse, err := GetBaseTest().LightCellsApi.PostReverse(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t func TestCellsPostReverse	\n", GetBaseTest().GetTestNumber())
	}
}
