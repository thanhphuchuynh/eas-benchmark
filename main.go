package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("./template/template-daily.xlsx")
	if err != nil {
		log.Fatal(err)
	}
	f.SetCellValue("Sheet1", "B5", "phucht5")
	f.SetCellValue("Sheet1", "B6", "DOP")
	f.SetCellValue("Sheet1", "B7", "50.000.000")
	f.SetCellValue("Sheet1", "B8", "20.000.000")
	f.SetCellValue("Sheet1", "B9", "21/02/2023")

	for idx, row := range [][]interface{}{
		{"1", "21/2/22", "21/2/23", "1327192", "PDL", "chuyen tien", "tech", "301278", "2000000", "100000000", "800000000", "xxx"},
		{"2", "21/2/22", "21/2/23", "1327192", "PDL", "chuyen tien", "tech", "301278", "2000000", "100000000", "800000000", "xxx"},
		{"3", "21/2/22", "21/2/23", "1327192", "PDL", "chuyen tien", "tech", "301278", "2000000", "100000000", "800000000", "xxx"},
		{"4", "21/2/22", "21/2/23", "1327192", "PDL", "chuyen tien", "tech", "301278", "2000000", "100000000", "800000000", "xxx"},
		{"5", "21/2/22", "21/2/23", "1327192", "PDL", "chuyen tien", "tech", "301278", "2000000", "100000000", "800000000", "xxx"},
		{"6", "21/2/22", "21/2/23", "1327192", "PDL", "chuyen tien", "tech", "301278", "2000000", "100000000", "800000000", "xxx"},
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+14)
		fmt.Println(cell)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow("Sheet1", cell, &row)
	}

	if err := f.SaveAs("Book1.xlsx", excelize.Options{
		Password: "123",
	}); err != nil {
		fmt.Println(err)
	}
}

// const (
// 	SheetName = "Expense Report"
// )

// func main() {
// 	// f := excelize.NewFile()
// 	// defer func() {
// 	// 	if err := f.Close(); err != nil {
// 	// 		fmt.Println(err)
// 	// 	}
// 	// }()
// 	// // Create a new sheet.
// 	// index, err := f.NewSheet("Sheet2")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// // Set value of a cell.
// 	// f.SetCellValue("Sheet2", "A2", "Hello world.")
// 	// f.SetCellValue("Sheet1", "B2", 100)
// 	// // Set active sheet of the workbook.
// 	// f.SetActiveSheet(index)
// 	// Save spreadsheet by the given path.
// 	// if err := f.SaveAs("Book1.xlsx", excelize.Options{
// 	// 	Password: "234",
// 	// }); err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// ff, err := os.Create("data.xlsx")

// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// f.WriteTo()
// 	// buf, _ := f.WriteToBuffer()
// 	// buf.WriteTo(ff)

// 	// f, buf := NewFile(), bytes.Buffer{}
// 	// f.ContentTypes, f.Path = nil, filepath.Join("test", "TestWriteTo.xlsx")
// 	// f.Pkg.Store(defaultXMLPathContentTypes, MacintoshCyrillicCharset)
// 	// _, err := f.WriteTo(bufio.NewWriter(&buf))

// 	var err error
// 	f := excelize.NewFile()
// 	index, _ := f.NewSheet("Sheet1")
// 	f.SetActiveSheet(index)
// 	f.SetSheetName("Sheet1", SheetName)

// 	err = f.SetColWidth(SheetName, "A", "A", 6)
// 	err = f.SetColWidth(SheetName, "H", "H", 6)
// 	err = f.SetColWidth(SheetName, "B", "B", 12)
// 	err = f.SetColWidth(SheetName, "C", "C", 16)
// 	err = f.SetColWidth(SheetName, "D", "D", 13)
// 	err = f.SetColWidth(SheetName, "E", "E", 15)
// 	err = f.SetColWidth(SheetName, "F", "F", 22)
// 	err = f.SetColWidth(SheetName, "G", "G", 13)

// 	err = f.SetRowHeight(SheetName, 1, 12)
// 	err = f.MergeCell(SheetName, "A1", "H1")

// 	err = f.SetRowHeight(SheetName, 2, 25)
// 	err = f.MergeCell(SheetName, "B2", "D2")

// 	style, err := f.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 20, Color: "6d64e8"}})
// 	err = f.SetCellStyle(SheetName, "B2", "D2", style)
// 	err = f.SetSheetRow(SheetName, "B2", &[]interface{}{"Gigashots Inc."})
// 	err = f.MergeCell(SheetName, "B3", "D3")
// 	err = f.SetSheetRow(SheetName, "B3", &[]interface{}{"3154 N Richardt Ave"})

// 	err = f.MergeCell(SheetName, "B4", "D4")
// 	err = f.SetSheetRow(SheetName, "B4", &[]interface{}{"Indianapolis, IN 46276"})

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "666666"}})
// 	err = f.MergeCell(SheetName, "B5", "D5")
// 	err = f.SetCellStyle(SheetName, "B5", "D5", style)
// 	err = f.SetSheetRow(SheetName, "B5", &[]interface{}{"(317) 854-0398"})

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 32, Color: "2B4492", Bold: true}})
// 	err = f.MergeCell(SheetName, "B7", "G7")
// 	err = f.SetCellStyle(SheetName, "B7", "G7", style)
// 	err = f.SetSheetRow(SheetName, "B7", &[]interface{}{"Expense Report"})

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 13, Color: "E25184", Bold: true}})
// 	err = f.MergeCell(SheetName, "B8", "C8")
// 	err = f.SetCellStyle(SheetName, "B8", "C8", style)
// 	err = f.SetSheetRow(SheetName, "B8", &[]interface{}{"09/04/00 - 09/05/00"})

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 13, Bold: true}})
// 	err = f.SetCellStyle(SheetName, "B10", "G10", style)
// 	err = f.SetSheetRow(SheetName, "B10", &[]interface{}{"Name", "", "Employee ID", "", "Department"})
// 	err = f.MergeCell(SheetName, "B10", "C10")
// 	err = f.MergeCell(SheetName, "D10", "E10")
// 	err = f.MergeCell(SheetName, "F10", "G10")

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "666666"}})
// 	err = f.SetCellStyle(SheetName, "B11", "G11", style)
// 	err = f.SetSheetRow(SheetName, "B11", &[]interface{}{"John Doe", "", "#1B800XR", "", "Brand & Marketing"})
// 	err = f.MergeCell(SheetName, "B11", "C11")
// 	err = f.MergeCell(SheetName, "D11", "E11")
// 	err = f.MergeCell(SheetName, "F11", "G11")

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Size: 13, Bold: true}})
// 	err = f.SetCellStyle(SheetName, "B13", "G13", style)
// 	err = f.SetSheetRow(SheetName, "B13", &[]interface{}{"Manager", "", "Purpose"})
// 	err = f.MergeCell(SheetName, "B13", "C13")
// 	err = f.MergeCell(SheetName, "D13", "E13")

// 	style, err = f.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "666666"}})
// 	err = f.SetCellStyle(SheetName, "B14", "G14", style)
// 	err = f.SetSheetRow(SheetName, "B14", &[]interface{}{"Jane Doe", "", "Brand Campaign"})
// 	err = f.MergeCell(SheetName, "B14", "C14")
// 	err = f.MergeCell(SheetName, "D14", "E14")

// 	var (
// 		expenseData = [][]interface{}{
// 			{"2022-04-10", "Flight", "Trip to San Fransisco", "", "", "$3,462.00"},
// 			{"2022-04-10", "Hotel", "Trip to San Fransisco", "", "", "$1,280.00"},
// 			{"2022-04-12", "Swags", "App launch", "", "", "$862.00"},
// 			{"2022-03-15", "Marketing", "App launch", "", "", "$7,520.00"},
// 			{"2022-04-11", "Event hall", "App launch", "", "", "$2,080.00"},
// 		}
// 	)

// 	style, err = f.NewStyle(&excelize.Style{
// 		Font:      &excelize.Font{Size: 13, Bold: true, Color: "2B4492"},
// 		Alignment: &excelize.Alignment{Vertical: "center"},
// 	})
// 	err = f.SetCellStyle(SheetName, "B17", "G17", style)
// 	err = f.SetSheetRow(SheetName, "B17", &[]interface{}{"Date", "Category", "Description", "", "Notes", "Amount"})
// 	err = f.MergeCell(SheetName, "D17", "E17")
// 	err = f.SetRowHeight(SheetName, 17, 32)

// 	startRow := 18
// 	for i := startRow; i < (len(expenseData) + startRow); i++ {
// 		var fill string
// 		if i%2 == 0 {
// 			fill = "F3F3F3"
// 		} else {
// 			fill = "FFFFFF"
// 		}

// 		style, err = f.NewStyle(&excelize.Style{
// 			Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{fill}},
// 			Font:      &excelize.Font{Color: "666666"},
// 			Alignment: &excelize.Alignment{Vertical: "center"},
// 		})
// 		err = f.SetCellStyle(SheetName, fmt.Sprintf("B%d", i), fmt.Sprintf("G%d", i), style)
// 		err = f.SetSheetRow(SheetName, fmt.Sprintf("B%d", i), &expenseData[i-18])
// 		err = f.SetCellRichText(SheetName, fmt.Sprintf("C%d", i), []excelize.RichTextRun{
// 			{Text: expenseData[i-18][1].(string), Font: &excelize.Font{Bold: true}},
// 		})

// 		err = f.MergeCell(SheetName, fmt.Sprintf("D%d", i), fmt.Sprintf("E%d", i))
// 		err = f.SetRowHeight(SheetName, i, 18)
// 	}
// 	err = f.SaveAs("expense-report.xlsx")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

// package main

// import (
// 	"encoding/csv"
// 	"log"
// 	"os"
// )

// func main() {
// 	data := [][]string{
// 		{"Báo cáo sổ phụ trong ngày"},
// 		{"Tài khoản:", "phucht5"},
// 		{"Loại sản phẩm:", "dop"},

// 		{"Số dư kế toán hiện thời:", "xxxxxxxxxxxx"},
// 		{"Số dư thực tế hiện thời:", "xxxxxxxxxxxx"},
// 		{"Chi tiết tại thời điểm:", "xxxxxxxxxxxx"},

// 		{"STT", "Ngày giao dịch", "Ngày hiệu lực", "Số tham chiếu", "Đại lý", "Diễn giải", "Ngân hàng chuyển tiền đến", "Mã điện thành công", "Số tiền rút", "Sốt tiền gửi", "Số dư", "ToupId"},

// 		{"1", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"2", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"3", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"4", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"5", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"6", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"7", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 		{"8", "21-12-2022", "21-12-2022", "1312312", "PDL", "dien giai", "techcombank", "987612", "900000000000", "888888888888", "6666666666666", "Vghda9"},
// 	}

// 	// create a file
// 	file, err := os.Create("result.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// initialize csv writer
// 	writer := csv.NewWriter(file)

// 	defer writer.Flush()

// 	// write all rows at once
// 	writer.WriteAll(data)

// 	// write single row
// 	// extraData := []string{"lettuce", "raspberry"}
// 	// writer.Write(extraData)
// }
