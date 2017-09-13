package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "os"
    "io"
    "strings"
    "bytes"
    "encoding/csv"
    "reflect"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the NTC HQ Calls Data API Home Page!")
    fmt.Println("Endpoint Hit: homePage")
}

func HQCallExcelReport(w http.ResponseWriter, r *http.Request){
    fmt.Println("URL: " + r.URL.Host + r.URL.Path)
    sendHQCallToExcelReport(w, r)
}


func getReportDataSet(w http.ResponseWriter, r *http.Request){
    fmt.Println("URL: " + r.URL.Host + r.URL.Path)
    data := getReportRawDataResultSet(w, r)
    fmt.Println("Endpoint Hit: getReportDataSet")
    json.NewEncoder(w).Encode(data);
}

func getReportDataSetRawData(w http.ResponseWriter, r *http.Request){
    fmt.Println("URL: " + r.URL.Host + r.URL.Path)
    data := getReportRawDataResultSet(w, r)
    fmt.Println("Endpoint Hit: getReportDataSetRawData")

    b := &bytes.Buffer{} // creates IO Writer
	 writer := csv.NewWriter(b) 

  var record []string
    record = append(record, "Zone")
    record = append(record, "Region")
    record = append(record, "Territory_Owner")
    record = append(record, "Type")
    record = append(record, "Territory")
    record = append(record, "Headquarter_ID")
    record = append(record, "CallHeadquarterID")
    record = append(record, "FirstCall")
    record = append(record, "LastCall")
    record = append(record, "HeadquarterStatus")
    record = append(record, "HeadquarterNumber")
    record = append(record, "Address")
    record = append(record, "City")
    record = append(record, "State")
    record = append(record, "Zip")
    record = append(record, "Headquarter")
    record = append(record, "PrimaryWholesaler")
    record = append(record, "PWCustomerNumber")
    record = append(record, "HeadquarterID")
    record = append(record, "Frequency")
    record = append(record, "Year")
    record = append(record, "Week")
    record = append(record, "Quarter")
    record = append(record, "Assigned")
    record = append(record, "CallDate")
    record = append(record, "CallType")
    record = append(record, "LastDistributionCall")
    record = append(record, "TerritoryCoverage")
    record = append(record, "Contact")
    record = append(record, "Stokers Cans") 
    record = append(record, "Stokers Tubs")
    record = append(record, "Stokers 3 oz")
    record = append(record, "Stokers 8 oz")
    record = append(record, "Stokers 16 oz")
    record = append(record, "Beech-Nut/Tro/Dur/HB")
    record = append(record, "Wind River Loose Leaf/Twist")
    record = append(record, "V2 Closed System - PP")
    record = append(record, "V2 Closed System - NPP")
    record = append(record, "V2 Vapor")
    record = append(record, "V2 Pro Series 3")
    record = append(record, "V2 Pro Series 7")
    record = append(record, "ZZ Vapor")
    record = append(record, "ZZ Cigarillos - 2 PK NPP")
    record = append(record, "ZZ Cigarillos - 2 PK PP")
    record = append(record, "ZZ Cigarillos - 3 PK PP")
    record = append(record, "ZZ Slo Burn")
    record = append(record, "ZZ Papers")
    record = append(record, "ZZ Wraps")
    record = append(record, "ZZ Wraps - Rillo Size")
    record = append(record, "ZZ Wraps - Xtra Wide")
    record = append(record, "ZZ Cones")
    record = append(record, "Primal - Cones")
    record = append(record, "Primal - Wraps")
    record = append(record, "Primal - Shisha")
    record = append(record, "Competitive")
    record = append(record, "MYO Tobacco")
    record = append(record, "Pipe Tobacco")
    writer.Write(record)

    fmt.Println("Header created")

    for i := 0; i < len(data.Rows); i++ {
      var elementRawData = data.Rows[i]
      resultReflectionValue := reflect.ValueOf(elementRawData)
      resultInterface := resultReflectionValue.Interface()
      resultMap := resultInterface.(map[string]interface{})
      /*
      if(i == 0){
          keys := reflect.ValueOf(resultMap).MapKeys()
          var record []string
          for j := 0; j < len(keys); j++ {
            record = append(record, keys[j].String())
         }
         writer.Write(record)
      }*/

     var record []string
      record = append(record, resultMap["Zone"].(string))
      record = append(record, resultMap["Region"].(string))
      record = append(record, resultMap["Territory_Owner"].(string))
      record = append(record, resultMap["Type"].(string))
      record = append(record, resultMap["Territory"].(string))
      record = append(record, resultMap["Headquarter_ID"].(string))
      record = append(record, resultMap["CallHeadquarterID"].(string))
      record = append(record, resultMap["FirstCall"].(string))
      record = append(record, resultMap["LastCall"].(string))
      record = append(record, resultMap["HeadquarterStatus"].(string))
      record = append(record, resultMap["HeadquarterNumber"].(string))
      record = append(record, resultMap["Address"].(string))
      record = append(record, resultMap["City"].(string))
      record = append(record, resultMap["State"].(string))
      record = append(record, resultMap["Zip"].(string))
      record = append(record, resultMap["Headquarter"].(string))
      record = append(record, resultMap["PrimaryWholesaler"].(string))
      record = append(record, resultMap["PWCustomerNumber"].(string))
      record = append(record, resultMap["HeadquarterID"].(string))
      record = append(record, resultMap["Frequency"].(string))
      record = append(record, resultMap["Year"].(string))
      record = append(record, resultMap["Week"].(string))
      record = append(record, resultMap["Quarter"].(string))
      record = append(record, resultMap["Assigned"].(string))
      record = append(record, resultMap["CallDate"].(string))
      record = append(record, resultMap["CallType"].(string))
      record = append(record, resultMap["LastDistributionCall"].(string))
      record = append(record, resultMap["TerritoryCoverage"].(string))
      record = append(record, resultMap["Contact"].(string))
      record = append(record, resultMap["Stokers Cans"].(string)) 
      record = append(record, resultMap["Stokers Tubs"].(string))
      record = append(record, resultMap["Stokers 3 oz"].(string))
      record = append(record, resultMap["Stokers 8 oz"].(string))
      record = append(record, resultMap["Stokers 16 oz"].(string))
      record = append(record, resultMap["Beech-Nut/Tro/Dur/HB"].(string))
      record = append(record, resultMap["Wind River Loose Leaf/Twist"].(string))
      record = append(record, resultMap["V2 Closed System - PP"].(string))
      record = append(record, resultMap["V2 Closed System - NPP"].(string))
      record = append(record, resultMap["V2 Vapor"].(string))
      record = append(record, resultMap["V2 Pro Series 3"].(string))
      record = append(record, resultMap["V2 Pro Series 7"].(string))
      record = append(record, resultMap["ZZ Vapor"].(string))
      record = append(record, resultMap["ZZ Cigarillos - 2 PK NPP"].(string))
      record = append(record, resultMap["ZZ Cigarillos - 2 PK PP"].(string))
      record = append(record, resultMap["ZZ Cigarillos - 3 PK PP"].(string))
      record = append(record, resultMap["ZZ Slo Burn"].(string))
      record = append(record, resultMap["ZZ Papers"].(string))
      record = append(record, resultMap["ZZ Wraps"].(string))
      record = append(record, resultMap["ZZ Wraps - Rillo Size"].(string))
      record = append(record, resultMap["ZZ Wraps - Xtra Wide"].(string))
      record = append(record, resultMap["ZZ Cones"].(string))
      record = append(record, resultMap["Primal - Cones"].(string))
      record = append(record, resultMap["Primal - Wraps"].(string))
      record = append(record, resultMap["Primal - Shisha"].(string))
      record = append(record, resultMap["Competitive"].(string))
      record = append(record, resultMap["MYO Tobacco"].(string))
      record = append(record, resultMap["Pipe Tobacco"].(string))
      writer.Write(record)
    }

    writer.Flush()

    w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv

	 w.Header().Set("Content-Type", "text/csv")
	 w.Header().Set("Content-Disposition", "attachment;filename=HQCalls.csv")
	 w.Write(b.Bytes())
}

func getReportRawDataResultSet(w http.ResponseWriter, r *http.Request) RawDataSearchResults{
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

    return dalGetReportDataSet(requestedPageNumber, rowsCount, columnToSort, sortDirection, filterBuffer.String());
}

func handleRequests(serverAddress string) {

    myRouter := mux.NewRouter().StrictSlash(true).UseEncodedPath()
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/reportdataset", getReportDataSet).Methods("GET", "POST")
    myRouter.HandleFunc("/reportdatasetrawdata", getReportDataSetRawData).Methods("GET", "POST")
    myRouter.HandleFunc("/hqCallExcelReport", HQCallExcelReport).Methods("GET", "POST")

    http.Handle("/", &MyServer{myRouter})
    http.ListenAndServe(serverAddress, nil);
}

type MyServer struct {
    r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    fmt.Println(req.Method)
    if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }
    // Stop here if its Preflighted OPTIONS request
    if req.Method == "OPTIONS" {
        return
    }
    // Lets Gorilla work
    s.r.ServeHTTP(rw, req)
}


func main() {
    var serverPort = os.Getenv("NTC_HQ_CALLS_APP_PORT")
    var serverAddress = ":" + serverPort;
    fmt.Println("Server Port:", serverPort)

    fmt.Println("NTC HQ Calls Service, Version 1.0, Port " + serverPort)
    handleRequests(serverAddress)
}
