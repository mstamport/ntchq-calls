package main

import (
    _ "github.com/denisenkom/go-mssqldb"
    "database/sql"
    "log"
    "fmt"
    "flag"
    "github.com/junhsieh/goexamples/fieldbinding/fieldbinding"
    "os"
    "strconv"
)

var dbServer= os.Getenv("NTC_REPORT_DB_SERVER")
var dbDatabase = os.Getenv("NTC_REPORT_DB_DATABASE")
var dbUser = os.Getenv("NTC_REPORT_DB_USER")
var dbPassword = os.Getenv("NTC_REPORT_DB_PASSWORD")
var dbPort = os.Getenv("NTC_REPORT_DB_PORT")


var server = flag.String("server", dbServer, "the database server")
var user = flag.String("user", dbUser, "the database user")
var password = flag.String("password", dbPassword, "the database password")
var port  = flag.String("port", dbPort, "the database port")
var database = flag.String("d", dbDatabase, "db_name")
var connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;encrypt=disable;database=%s", *server, *user, *password, *port, *database)
var _conn *sql.DB = nil

func getConn() *sql.DB {
    if _conn != nil {
        fmt.Println("getConn: return existing connection.")
        return _conn
    }

    var err error
    fmt.Println("getConn: create new connection.")
    _conn, err = sql.Open("mssql", connString)
    if err != nil {
        log.Panic("Open connection failed:", err.Error())
    }
    return _conn
}

func dalExecuteProcedureWhichReturnsGenericDataSet(query string, args ...interface{})[]interface{} {
   var conn = getConn();

    rows, err := conn.Query(query, args...)
    if err != nil {
        log.Panic(err)
    }

    // create a fieldbinding object.
     var fArr []string
     fb := fieldbinding.NewFieldBinding()

     if fArr, err = rows.Columns(); err != nil {
       log.Panic(err)
     }

     fb.PutFields(fArr)
     outArr := []interface{}{}

     for rows.Next() {
       if err := rows.Scan(fb.GetFieldPtrArr()...); err != nil {
         log.Panic(err)
       }

       outArr = append(outArr, fb.GetFieldArr())
     }

    if err := rows.Err(); err != nil {
        log.Panic(err)
    }

    return outArr;
}

func dalGetReportDataSet(requestedPageNumber string, rowsCount string, columnToSort string, sortDirection string, filters string) RawDataSearchResults {
   var conn = getConn();

    data := RawDataSearchResults{
    }

   rows, err := conn.Query("exec dbo.sHeadquarterCallReport_GetTotalRowsCount ?", filters)
    if err != nil {
        log.Panic(err)
    }
    for rows.Next() {
      var rowsCount string
      if err := rows.Scan(&rowsCount); err != nil {
            log.Panic(err)
        }
      data.TotalRecords = rowsCount
    }
    if (requestedPageNumber == "") {
	requestedPageNumber = "1"
    }
    data.Page = requestedPageNumber
    totalRecords, err := strconv.Atoi(data.TotalRecords)
    pageRowsCount, err := strconv.Atoi(rowsCount)
    if(pageRowsCount <= 0){
	pageRowsCount = 20
    }

    data.TotalPages =  strconv.Itoa(totalRecords / pageRowsCount)
    data.Rows = dalExecuteProcedureWhichReturnsGenericDataSet("exec dbo.sHeadquarterCallReport ?, ?, ?, ?, ?", requestedPageNumber, rowsCount, columnToSort,  sortDirection,  filters)

    return data;
}


func dalGetExcelDataSet(requestedPageNumber string, rowsCount string, columnToSort string, sortDirection string, filters string) ExcelDataSearchResults {
   var conn = getConn();

    data := ExcelDataSearchResults {
    }

   rows, err := conn.Query("exec dbo.sHeadquarterCallReport_GetTotalRowsCount ?", filters)
    if err != nil {
        log.Panic(err)
    }
    for rows.Next() {
      var rowsCount string
      if err := rows.Scan(&rowsCount); err != nil {
            log.Panic(err)
        }
      data.TotalRecords = rowsCount
    }
    if (requestedPageNumber == "") {
	requestedPageNumber = "1"
    }

    totalRecords, err := strconv.Atoi(data.TotalRecords)
    pageRowsCount, err := strconv.Atoi(rowsCount)
    if(pageRowsCount <= 0){
	pageRowsCount = 20
    }

    data.TotalPages =  strconv.Itoa(totalRecords / pageRowsCount)
    data.Rows = dalExecuteProcedureWhichReturnsGenericDataSet("exec dbo.sHeadquarterCallReport ?, ?, ?, ?, ?", requestedPageNumber, rowsCount, columnToSort,  sortDirection,  filters)

    return data;
}