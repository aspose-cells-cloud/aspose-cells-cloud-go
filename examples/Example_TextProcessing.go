package main

import (
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v26"
)

func main() {
	BookTextXlsx := "BookText.xlsx"
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))

	_, _, err := instance.TrimCharacter(&TrimCharacterRequest{Spreadsheet: BookTextXlsx, TrimLeading: true, TrimTrailing: true, TrimSpaceBetweenWordTo1: true}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.TrimCharacter(&TrimCharacterRequest{Spreadsheet: BookTextXlsx, TrimLeading: true, TrimTrailing: true, TrimSpaceBetweenWordTo1: true, Worksheet: "Text", Range_: "J6:J6"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out1.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.TrimCharacter(&TrimCharacterRequest{Spreadsheet: BookTextXlsx, TrimLeading: true, TrimTrailing: false, TrimSpaceBetweenWordTo1: true, RemoveExtraLineBreaks: true, Worksheet: "Text", Range_: "J7:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out2.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.UpdateWordCase(&UpdateWordCaseRequest{Spreadsheet: BookTextXlsx, WordCaseType: "UpperCase", Worksheet: "Text", Range_: "J7:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out3.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.UpdateWordCase(&UpdateWordCaseRequest{Spreadsheet: "out3.xlsx", WordCaseType: "LowerCase", Worksheet: "Text", Range_: "J7:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out4.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.UpdateWordCase(&UpdateWordCaseRequest{Spreadsheet: "out4.xlsx", WordCaseType: "ProperCase", Worksheet: "Text", Range_: "J7:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out5.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.UpdateWordCase(&UpdateWordCaseRequest{Spreadsheet: "out5.xlsx", WordCaseType: "SentenceCase", Worksheet: "Text", Range_: "J7:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out6.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}

	_, _, err = instance.AddText(&AddTextRequest{Spreadsheet: BookTextXlsx, Text: "Hello, ", Position: "AtTheBeginning"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out7.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}

	_, _, err = instance.AddText(&AddTextRequest{Spreadsheet: BookTextXlsx, Text: "!", Position: "AfterText", SelectText: "hello world"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out8.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.AddText(&AddTextRequest{Spreadsheet: BookTextXlsx, Text: "!", Position: "AtTheEnd", Worksheet: "Text", Range_: "A1:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out9.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}

	_, _, err = instance.ConvertText(&ConvertTextRequest{Spreadsheet: BookTextXlsx, ConvertTextType: "ConvertNumberToText"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out10.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.ConvertText(&ConvertTextRequest{Spreadsheet: BookTextXlsx, ConvertTextType: "ConvertCharacters", SourceCharacters: "1", TargetCharacters: "Hello World", Worksheet: "Text", Range_: "A1:J7"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out11.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.ConvertText(&ConvertTextRequest{Spreadsheet: BookTextXlsx, ConvertTextType: "ConvertLinebreak", SourceCharacters: "", TargetCharacters: "Linebreak"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out12.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}
	_, _, err = instance.ConvertText(&ConvertTextRequest{Spreadsheet: BookTextXlsx, ConvertTextType: "ConvertWriteSpace", SourceCharacters: "", TargetCharacters: "WriteSpace"}, CellsCloudOption{OptionName: "LocalOutPath", OptionValue: "out13.xlsx"})
	if err != nil {
		println("Error:", err.Error())
	}

}
