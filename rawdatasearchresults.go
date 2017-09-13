package main

type RawDataSearchResults struct {
    Page string `json:"page"`
    TotalPages string `json:"total"`
    TotalRecords string `json:"records"`
    Rows []interface{} `json:"rows"`
}

type ExcelDataSearchResults  struct {
    Page string `json:"page"`
    TotalPages string `json:"total"`
    TotalRecords string `json:"records"`
    Rows []interface{} `json:"rows"`
}

type ExcelColumnResults struct {
    Page string `json:"page"`
    TotalPages string `json:"total"`
    TotalRecords string `json:"records"`
    Rows []interface{} `json:"rows"`
}