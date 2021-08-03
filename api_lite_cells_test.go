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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.DeleteMetadata(fileMap, nil)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostAssemble(fileMap, postAssembleOpts)
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

	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostClearObjects(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostClearObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsPostExport(t *testing.T) {
	name := GetDataSource()
	var fileMap map[string]string
	fileMap = make(map[string]string)
	fileMap[name] = "TestData\\" + name
	name = GetAssemblyTest()
	fileMap[name] = "TestData\\" + name
	postOpts := new(PostExportOpts)
	postOpts.Format = "pdf"
	postOpts.ObjectType = "chart"
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostExport(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\tTestCellsPostExport \n", GetBaseTest().GetTestNumber())
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

	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostMerge(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostMetadata(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostProtect(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostSearch(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostSearch(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostSplit(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostUnlock(fileMap, postOpts)
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
	_, httpResponse, err := GetBaseTest().LiteCellsAPI.PostWatermark(fileMap, postOpts)
	if err != nil {
		t.Error(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		t.Fail()
	} else {
		fmt.Printf("%d\t TestCellsPostWatermark \n", GetBaseTest().GetTestNumber())
	}
}
