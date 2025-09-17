package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v25"
)

func main() {
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	RemoteFolder := "SDKGo"
	CompanySalesXlsx := "CompanySales.xlsx"
	EmployeeSalesSummaryXlsx := "EmployeeSalesSummary.xlsx"

	// Check if file exists.
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
	// Upload source files.
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: CompanySalesXlsx, Path: RemoteFolder + "/" + CompanySalesXlsx})
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: EmployeeSalesSummaryXlsx, Path: RemoteFolder + "/" + EmployeeSalesSummaryXlsx})
	// Create folder on cloud storage.
	httpResponse, err = instance.CreateFolder(&CreateFolderRequest{Path: RemoteFolder + "/CellsCloud"})
	// Move file to other folder on cloud storage.
	httpResponse, err = instance.MoveFile(&MoveFileRequest{SrcPath: RemoteFolder + "/" + CompanySalesXlsx, DestPath: RemoteFolder + "/CellsCloud/" + CompanySalesXlsx})
	// Copy a file to other file on cloud storage.
	httpResponse, err = instance.CopyFile(&CopyFileRequest{SrcPath: RemoteFolder + "/" + EmployeeSalesSummaryXlsx, DestPath: RemoteFolder + "/CellsCloud/" + EmployeeSalesSummaryXlsx})
	// Move folder on cloud storage.
	httpResponse, err = instance.MoveFolder(&MoveFolderRequest{SrcPath: RemoteFolder + "/CellsCloud", DestPath: RemoteFolder + "/CellsCloud2"})
	// Get files list in a folder.
	fileList, _, _ := instance.GetFilesList(&GetFilesListRequest{Path: RemoteFolder + "/CellsCloud2"})
	for index, storageFile := range fileList.Value {
		print("Get File")
		print(index)
		println(storageFile.Name)
	}
	// Delete folder on cloud storage.
	httpResponse, err = instance.DeleteFolder(&DeleteFolderRequest{Path: RemoteFolder + "/CellsCloud2"})
	// Delete file on cloud storage.
	httpResponse, err = instance.DeleteFile(&DeleteFileRequest{Path: RemoteFolder + "/" + CompanySalesXlsx})
}
