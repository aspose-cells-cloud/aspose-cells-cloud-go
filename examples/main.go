package main

import (
	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
	"context"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	client := asposecellscloud.NewAsposeCellsCloudClient(os.Getenv("AsposeCellsCloudClientId"), os.Getenv("AsposeCellsCloudClientSecret"), os.Getenv("CellsCloudApiBaseUrl"))
	fontSetting := models.FontSetting{
		StartIndex: asposecellscloud.Int32Ptr(0),
		Length:     asposecellscloud.Int32Ptr(3),
		Font:       &models.Font{Size: asposecellscloud.Int32Ptr(15)},
	}
	fontSettings := []models.FontSetting{fontSetting}

	responses, err := client.Do(context.Background(),
		requests.NewUploadFileRequest("BookText.xlsx", "BookText_New.xlsx", requests.WithCommonParameter("StorageName", "Cells")),
		requests.NewPostCellCharactersRequest("BookText_New.xlsx", "Text", "D4", requests.WithCommonParameter("StorageName", "Cells"),
			requests.WithCommonParameter("Options", fontSettings)),
	)
	if err != nil {
		println(err.Error())
	}
	for _, response := range responses {
		println(response.ToString())
	}

}
