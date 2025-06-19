package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	RemoteFolder := "SDKGo"
	CompanySalesXlsx := "CompanySales.xlsx"
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"

	objectExists, httpResponse, err := instance.ObjectExists(&ObjectExistsRequest{Path: RemoteFolder})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print("E")
	} else {
		if !objectExists.Exists {
			httpResponse, err = instance.CreateFolder(&CreateFolderRequest{Path: RemoteFolder})
		}
	}

	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: CompanySalesXlsx, Path: RemoteFolder + "/" + CompanySalesXlsx})
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: EmployeeSalesSummaryXlsx, Path: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	print("UploadFiles")
	httpResponse, err = instance.CreateFolder(&CreateFolderRequest{Path: RemoteFolder + "/CellsCloud"})
	httpResponse, err = instance.MoveFile(&MoveFileRequest{SrcPath: RemoteFolder + "/" + CompanySalesXlsx, DestPath: RemoteFolder + "/CellsCloud/" + CompanySalesXlsx})
	httpResponse, err = instance.CopyFile(&CopyFileRequest{SrcPath: RemoteFolder + "/" + EmployeeSalesSummaryXlsx, DestPath: RemoteFolder + "/CellsCloud/" + EmployeeSalesSummaryXlsx})
	httpResponse, err = instance.MoveFolder(&MoveFolderRequest{SrcPath: RemoteFolder + "/CellsCloud", DestPath: RemoteFolder + "/CellsCloud2"})
	fileList, _, _ := instance.GetFilesList(&GetFilesListRequest{Path: RemoteFolder + "/CellsCloud2"})
	for index, storageFile := range fileList.Value {
		print("Get File")
		print(index)
		println(storageFile.Name)
	}
	httpResponse, err = instance.DeleteFolder(&DeleteFolderRequest{Path: RemoteFolder + "/CellsCloud2"})
	httpResponse, err = instance.DeleteFile(&DeleteFileRequest{Path: RemoteFolder + "/" + CompanySalesXlsx})
}
