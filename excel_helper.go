package main

import (
    "fmt"
    "net/http"
    "github.com/tealeg/xlsx"
    "bytes"
    "math"
    "io"
    "strings"
    "encoding/json"
    "reflect"
)


func sendHQCallToExcelReport(w http.ResponseWriter, r *http.Request){
    r.ParseForm()
    requestedPageNumber := r.FormValue("page")
    rowsCount :=  r.FormValue("rows")
    columnToSort :=  r.FormValue("sidx")
    sortDirection :=  r.FormValue("sord")
    filters :=  r.FormValue("filters")
    
    fmt.Println("RequestedPageNumber : " + requestedPageNumber )
    fmt.Println("RowsCount : " + rowsCount )
    fmt.Println("ColumnToSort : " + columnToSort )
    fmt.Println("SortDirection : " + sortDirection )
    fmt.Println("Filters : " + filters )

    dec := json.NewDecoder(strings.NewReader(filters))
    var f Filters
    var filterBuffer bytes.Buffer

    for {
		if err := dec.Decode(&f); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
	}

   for index, element := range f.Rules {

       if(index == 0){
         filterBuffer.WriteString(" where ")
       }
       if(index > 0){
         filterBuffer.WriteString(" and ")
       }
       filterBuffer.WriteString(element.Field)
       switch element.Op{
          case "bw": // Begins With
            filterBuffer.WriteString( " like '" + element.Data + "%'")
         case "bn": // Does not begin with
            filterBuffer.WriteString( " not like '" + element.Data + "%'")
          case "eq": // Equal
            filterBuffer.WriteString( " = '" + element.Data + "'")
          case "ne": // Not Equal
            filterBuffer.WriteString( " != '" + element.Data + "'")       
          case "lt": // Less than
            filterBuffer.WriteString( " < '" + element.Data + "'")       
          case "le": // Less than or equal
            filterBuffer.WriteString( " <= '" + element.Data + "'")       
          case "gt": // Greater than
            filterBuffer.WriteString( " > '" + element.Data + "'")       
          case "ge": // Greater than or equal
            filterBuffer.WriteString( " >= '" + element.Data + "'")       
          case "ew": // Ends with
            filterBuffer.WriteString( " like '%" + element.Data + "'")       
          case "en": // Does not ends with
            filterBuffer.WriteString( " not like '%" + element.Data + "'")       
          case "nu": // is null
            filterBuffer.WriteString( " is null ")       
          case "nn": // is not null
            filterBuffer.WriteString( " is not null ")       
          case "in": // in
            filterBuffer.WriteString( " in (" + element.Data + ")")       
          case "ni": // not in
            filterBuffer.WriteString( " not in (" + element.Data + ")")       
          case "nc": // Does not contain
            filterBuffer.WriteString( " not like '%" + element.Data + "%'")       
          default: // contains
            filterBuffer.WriteString( " like '%" + element.Data + "%'")       
       }
       fmt.Println("SQL Filters : " + filterBuffer.String() )
   }
   

    data := dalGetExcelDataSet(requestedPageNumber, rowsCount, columnToSort, sortDirection, filterBuffer.String());
    columndata := dalGetExcelColumnsSet(requestedPageNumber, rowsCount, columnToSort, sortDirection, filterBuffer.String());


    //Generate excel content
    var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var err error
    var headerstyle *xlsx.Style
    var itemstyle *xlsx.Style
    var labelStyle *xlsx.Style

    file = xlsx.NewFile()
    headerstyle = xlsx.NewStyle()
    headerstyle.Alignment.Horizontal = "center"
    headerstyle.Alignment.WrapText = true
    headerstyle.Font = *xlsx.NewFont(10, "Arial")
    headerstyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")


    itemstyle = xlsx.NewStyle()
    itemstyle.Alignment.WrapText = true
    itemstyle.Font = *xlsx.NewFont(10, "Arial")
    itemstyle.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
    itemstyle.Border.LeftColor = "2E6E9E"
    itemstyle.Border.RightColor = "2E6E9E"
    itemstyle.Border.TopColor = "2E6E9E"
    itemstyle.Border.BottomColor = "2E6E9E"

    labelStyle = xlsx.NewStyle()
    labelStyle.Font = *xlsx.NewFont(11, "Verdana")
    labelStyle.Font.Italic = true
 
    sheet, err = file.AddSheet("HQ Calls")
    if err != nil {
        fmt.Printf(err.Error())
    }

    row = sheet.AddRow()
    addExcelCell(row, "HQ_HQID", headerstyle)
    addExcelCell(row, "CallHQ_ID", headerstyle)
    addExcelCell(row, "Zone", headerstyle)
    addExcelCell(row, "Region", headerstyle)
    addExcelCell(row, "Territory_Owner", headerstyle)
    addExcelCell(row, "Type", headerstyle)
    addExcelCell(row, "Territory", headerstyle)
    addExcelCell(row, "FirstCall", headerstyle)
    addExcelCell(row, "LastCall", headerstyle)
    addExcelCell(row, "HQName", headerstyle)
    addExcelCell(row, "HQStatus", headerstyle)
    addExcelCell(row, "HQNumber", headerstyle)
    addExcelCell(row, "Address", headerstyle)
    addExcelCell(row, "City", headerstyle)
    addExcelCell(row, "State", headerstyle)
    addExcelCell(row, "Zip", headerstyle)
    addExcelCell(row, "PrimaryWholesaler", headerstyle)
    addExcelCell(row, "PrimaryWholesalerCustNumber", headerstyle)
    addExcelCell(row, "Frequency", headerstyle)
    addExcelCell(row, "Year", headerstyle)
    addExcelCell(row, "Week", headerstyle)
    addExcelCell(row, "Quarter", headerstyle)
    addExcelCell(row, "Assigned", headerstyle)
    addExcelCell(row, "CallDate", headerstyle)
    addExcelCell(row, "CallType", headerstyle)
    addExcelCell(row, "LastDistributionCall", headerstyle)
    addExcelCell(row, "Territory Coverage", headerstyle)
    addExcelCell(row, "Notes", headerstyle)
    addExcelCell(row, "Contact", headerstyle)
    for i := 0; i < len(columndata.Rows); i++ {
      var elementRawData = columndata.Rows[i]
      resultReflectionValue := reflect.ValueOf(elementRawData)
      resultInterface := resultReflectionValue.Interface()
      resultMap := resultInterface.(map[string]interface{}) 
      addExcelCell(row, resultMap["ColumnName"].(string), itemstyle)
    }
  

    for i := 0; i < len(data.Rows); i++ {
      var elementRawData = data.Rows[i]
      resultReflectionValue := reflect.ValueOf(elementRawData)
      resultInterface := resultReflectionValue.Interface()
      resultMap := resultInterface.(map[string]interface{})

       // item level row
       row = sheet.AddRow()
       addExcelCell(row, resultMap["Headquarter_ID"].(string), itemstyle)
       addExcelCell(row, resultMap["CallHeadquarterID"].(string), itemstyle)
       addExcelCell(row, resultMap["Zone"].(string), itemstyle)
       addExcelCell(row, resultMap["Region"].(string), itemstyle)
       addExcelCell(row, resultMap["Territory_Owner"].(string), itemstyle)
       addExcelCell(row, resultMap["Type"].(string), itemstyle)
       addExcelCell(row, resultMap["Territory"].(string), itemstyle)
       addExcelCell(row, resultMap["FirstCall"].(string), itemstyle)
       addExcelCell(row, resultMap["LastCall"].(string), itemstyle)
       addExcelCell(row, resultMap["Headquarter"].(string), itemstyle)
       addExcelCell(row, resultMap["HeadquarterStatus"].(string), itemstyle)
       addExcelCell(row, resultMap["HeadquarterNumber"].(string), itemstyle)
       addExcelCell(row, resultMap["Address"].(string), itemstyle)
       addExcelCell(row, resultMap["City"].(string), itemstyle)
       addExcelCell(row, resultMap["State"].(string), itemstyle)
       addExcelCell(row, resultMap["Zip"].(string), itemstyle)
       addExcelCell(row, resultMap["PrimaryWholesaler"].(string), itemstyle)
       addExcelCell(row, resultMap["PWCustomerNumber"].(string), itemstyle)
       addExcelCell(row, resultMap["Frequency"].(string), itemstyle)
       addExcelCell(row, resultMap["Year"].(string), itemstyle)
       addExcelCell(row, resultMap["Week"].(string), itemstyle)
       addExcelCell(row, resultMap["Quarter"].(string), itemstyle)
       addExcelCell(row, resultMap["Assigned"].(string), itemstyle)
       addExcelCell(row, resultMap["CallDate"].(string), itemstyle)
       addExcelCell(row, resultMap["CallType"].(string), itemstyle)
       addExcelCell(row, resultMap["LastDistributionCall"].(string), itemstyle)
       addExcelCell(row, resultMap["TerritoryCoverage"].(string), itemstyle)
       addExcelCell(row, resultMap["Notes"].(string), itemstyle)
       addExcelCell(row, resultMap["Contact"].(string), itemstyle)
       for x := 0; x < len(columndata.Rows); x++ {
         var elementColumnData = columndata.Rows[x]
         resultReflectionColValue := reflect.ValueOf(elementColumnData )
         resultColInterface := resultReflectionColValue.Interface()
         resultColMap := resultColInterface.(map[string]interface{}) 
         addExcelCell(row, resultMap[resultColMap["ItemID"].(string)].(string), itemstyle)
       }


    }

     
    //column formatting
   sheet.SetColWidth(0, 0, 29)
   sheet.SetColWidth(1, 9, 15) 
   sheet.SetColWidth(27, 27, 150) 


    // Write to IO bytes
    b := &bytes.Buffer{} 
    file.Write(b)

    //Sent file back
    w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	 w.Header().Set("Content-Disposition", "attachment;filename=HQCalls.xlsx")
    w.Write(b.Bytes())
}

func addExcelCell(row *xlsx.Row, cellValue string, cellStyle *xlsx.Style) *xlsx.Cell{
    var cell *xlsx.Cell
    cell = row.AddCell()
    cell.Value = cellValue
    cell.SetStyle(cellStyle)
    return cell
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}