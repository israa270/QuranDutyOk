package pdf

import (
	"github.com/ebedevelopment/next-gen-tms/server/global"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"go.uber.org/zap"
)

const(
	Merchant = "merchant"
	Terminal = "terminal"
    Platform = "platform"
	AppCategory ="appCategory"
    Manufacturer = "manufacturer"
	Model = "model"
	MerchantType ="merchantType"
	Org  = "org"
	Parameter ="parameter"
	AppPush = "appPush"

	Login   = "login"
	Audit = "audit"
	User = "user"

	// Maker = "maker"
	// Checker ="checker"
)

type PDFUseCase struct{

}

func (e *PDFUseCase) CreatePDF(filePath string, header []string, data [][]string , title string,typeFn string ) (string, error){
	m := pdf.NewMaroto(consts.Landscape, consts.A3)
	m.SetPageMargins(20, 10, 20)
	
	if typeFn == Terminal {
		buildTerminalList(m, header, data, title)
	}else if typeFn == Merchant{
		buildMerchantList(m, header, data, title)
	}else if typeFn == Platform{
	    buildPlatformList(m, header, data, title)
	}else if  typeFn ==AppCategory{
		buildAppCategoryList(m, header, data, title)
	}else if typeFn == Manufacturer{
		buildManufacturerList(m, header, data, title)
	}else if typeFn == Model{
		buildModelList(m, header, data, title)
	}else if typeFn == MerchantType{
		buildMerchantTypeList(m, header, data, title)
	}else if typeFn == Parameter{
		buildParameterList(m, header, data, title)
	}else if typeFn ==AppPush{
		buildAppPushList(m, header, data, title)
	}else if typeFn == Org{
		buildOrgList(m, header, data, title)
	}else if typeFn ==Login{
		buildLoginList(m, header, data, title)
	}else if typeFn == Audit{
		buildAuditList(m, header, data, title)
	}else if typeFn == User{
		buildUserList(m, header, data, title)
	}
	
	err := m.OutputFileAndClose(filePath)
	if err != nil {
		global.GvaLog.Error("failed o save PDF" , zap.Error(err))
		return "", err
	}

	return filePath, nil
}




func buildModelList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3,3,3,3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3,3,3,3},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}



func buildMerchantList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())
	
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2, 1,1,1,1,1, 1,1,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2, 1,1,1,1,1, 1,1,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}



func buildAppPushList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())
	
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2,2,2,2,2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2,2,2,2,2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}



func buildParameterList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())
	
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2,2,2,2,2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2,2,2,2,2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}


func buildOrgList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())
	
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1,1,1,1, 1,1,1,1,1, 1,1,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1,1,1,1, 1,1,1,1,1, 1,1,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}



func buildMerchantTypeList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3,3,2,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3,3,2,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}




func buildManufacturerList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2, 1,1,1,1,1, 1, 1,1,1,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2,2, 1,1,1,1,1, 1, 1,1,1,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}


func buildTerminalList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {
	
	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{2, 2,2,2,2,1,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2,2,2,2,1,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}


func buildPlatformList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())
	

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2, 1,1,1,1,1, 1, 1,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2, 1,1,1,1,1, 1, 1,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}


func buildAppCategoryList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2, 1,1,2,2,1, 1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2, 1,1,2,2,1, 1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}


func buildLoginList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3,2,2,3,2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3,2,2,3,2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}

func buildAuditList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{1, 1, 1,1,1,2,1, 1, 1,1,1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 1, 1,1,1,2,1, 1, 1,1,1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}

func buildUserList(m pdf.Maroto, tableHeadings []string, contents [][]string, title string) {

	m.RegisterHeader(func() {
	m.Row(20, func() {
			m.Col(12, func() {
				m.Text(title, props.Text{
					Size: 16,
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: getDarkPurpleColor(),
				})
			})
		})
	})

	lightPurpleColor := getLightGreyColor()

	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 2,2,2,2,2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{2, 2,2,2,2,2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})

}


func getLightGreyColor() color.Color {
	return color.Color{
		Red:   211,
		Green: 211,
		Blue:  211,
	}
}