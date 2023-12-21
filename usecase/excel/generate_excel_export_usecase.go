package excel

import (
	"fmt"

	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

type ExcelUseCase struct{

}

func (e *ExcelUseCase) GenerateExcelSheet(data [][]string, filePath string) (string, error){
	sheetName := "Sheet1"

	// Create a new Excel file
	f := excelize.NewFile()

	// Add a new sheet to the file
	index,err := f.NewSheet(sheetName)
    if err != nil{
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GenerateExcelFile, zap.Error(err))
		return "", fmt.Errorf(global.GvaLoggerMessage["log"].GenerateExcelFile)
	}
	// Define the cell styles
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 13,
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#808080","#dddddd"},
			Pattern: 1,
		},
		Alignment: &excelize.Alignment{
			Vertical: "center",
			Horizontal: "center",
		},
	})
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GenerateExcelFile, zap.Error(err))
	}

	alternateStyle, err := f.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#f2f2f2"},
			Pattern: 1,
		},
	})
	if err != nil {
		global.GvaLog.Error(global.GvaLoggerMessage["log"].GenerateExcelFile, zap.Error(err))
	}

	// Set the header row and apply the header style
	header := data[0]
	for i, cell := range header {
		col,_ := excelize.ColumnNumberToName(i + 1)
		f.SetCellValue(sheetName, fmt.Sprintf("%s1", col), cell)
		f.SetCellStyle(sheetName, fmt.Sprintf("%s1", col), fmt.Sprintf("%s1", col), headerStyle)
	}

	// Set the data rows and apply the alternate row style
	for i, row := range data[1:] {
		for j, cell := range row {
			col,_ := excelize.ColumnNumberToName(j + 1)
			f.SetCellValue(sheetName, fmt.Sprintf("%s%d", col, i+2), cell)
			if i%2 == 0 {
				f.SetCellStyle(sheetName, fmt.Sprintf("%s%d", col, i+2), fmt.Sprintf("%s%d", col, i+2), alternateStyle)
			}
		}
	}

	// Set the column widths
	for i := 0; i < len(data[0]); i++ {
		col,_ := excelize.ColumnNumberToName(i + 1)
		f.SetColWidth(sheetName, col, col, 25)
	}

	// Set the active sheet
	f.SetActiveSheet(index)

	// Save the file
	err = f.SaveAs(filePath)
	
	return filePath, err
}