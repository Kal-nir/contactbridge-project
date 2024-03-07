package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getLead() {
	url := "http://localhost:8080/api/view_lead/"
	clientContact := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err)

	res, err := clientContact.Do(req)
	checkErr(err)

	if res.Body != nil {
		defer res.Body.Close()
	}

	buffer := new(strings.Builder)

	_, err = io.Copy(buffer, res.Body)
	checkErr(err)

	body := buffer.String()

	leads := &ViewLeadResults{}
	json := json.Unmarshal([]byte(body), leads)
	checkErr(json)

	fmt.Print("Lead ID\t\t")
	fmt.Print("Client Name\t")
	fmt.Print("Client Company\t")
	fmt.Print("Conversion Status\t")
	fmt.Print("Conversion Source\t")
	fmt.Print("Conversion Remarks\t\n")
	for _, lead := range leads.Data {
		status := lead.ConversionStatus
		source := lead.ConversionSource
		remarks := lead.ConversionRemarks

		if len(lead.ConversionStatus) == 0 {
			status = "None"
		}
		if len(lead.ConversionSource) == 0 {
			source = "None"
		}
		if len(lead.ConversionRemarks) == 0 {
			remarks = "None"
		}

		fmt.Printf(
			"%d\t\t%s\t%s\t%s\t%s\t%s\n\n",
			lead.LeadID,
			lead.ClientName,
			lead.ClientCompany,
			status,
			source,
			remarks,
		)
	}
}

func getLeadByID(id int) LeadConversion {
	url := "http://localhost:8080/api/lead_conversion/"
	clientContact := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err)

	res, err := clientContact.Do(req)
	checkErr(err)

	if res.Body != nil {
		defer res.Body.Close()
	}

	buffer := new(strings.Builder)

	_, err = io.Copy(buffer, res.Body)
	checkErr(err)

	body := buffer.String()

	leads := &LeadConversionResults{}
	json := json.Unmarshal([]byte(body), leads)
	checkErr(json)

	for i := 0; i < len(leads.Data); i++ {
		if leads.Data[i].LeadID == id {
			return leads.Data[i]
		}
	}
	return LeadConversion{LeadID: -1}
}

func addLead(reader *bufio.Reader) {
	fmt.Print("Enter conversion status: ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	conversion_status := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter conversion source: ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	conversion_source := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter conversion remarks: ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	conversion_remarks := strings.Replace(text, "\r\n", "", -1)

	body := AddLeadConversion{
		ConversionStatus:  conversion_status,
		ConversionSource:  conversion_source,
		ConversionRemarks: conversion_remarks,
	}

	url := "http://localhost:8080/api/lead_conversion/"
	body_json, err := json.Marshal(body)
	checkErr(err)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body_json))
	checkErr(err)
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func updateLead(reader *bufio.Reader) {
	fmt.Print("Enter ID to delete: ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	lead_id, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
	if err != nil {
		fmt.Println("Wrong input!")
		return
	}

	lead := getLeadByID(lead_id)
	if lead.LeadID == -1 {
		fmt.Println("ID not found!")
		return
	}

	fmt.Print("Enter new lead name ")
	fmt.Print("(Old: " + lead.ConversionStatus + ", empty if no changes): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	conversion_status := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter new company name ")
	fmt.Print("(Old: " + lead.ConversionSource + "): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	conversion_source := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter new email address ")
	fmt.Print("(Old: " + lead.ConversionRemarks + "): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	conversion_remarks := strings.Replace(text, "\r\n", "", -1)

	body := LeadConversion{
		ConversionStatus:  conversion_status,
		ConversionSource:  conversion_source,
		ConversionRemarks: conversion_remarks,
	}
	if body.LeadID == 0 {
		body.LeadID = lead.LeadID
	}
	if len(body.ConversionStatus) == 0 {
		body.ConversionStatus = lead.ConversionStatus
	}
	if len(body.ConversionSource) == 0 {
		body.ConversionSource = lead.ConversionSource
	}
	if len(body.ConversionRemarks) == 0 {
		body.ConversionRemarks = lead.ConversionRemarks
	}

	url := "http://localhost:8080/api/lead_conversion/" + strconv.Itoa(lead_id)
	body_json, err := json.Marshal(body)
	checkErr(err)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body_json))
	checkErr(err)

	req.Header.Set("Content-Type", "application/json")

	requester := &http.Client{}
	res, err := requester.Do(req)
	checkErr(err)
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func deleteLead(reader *bufio.Reader) {
	fmt.Print("Enter ID to edit: ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	lead_id, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
	if err != nil {
		fmt.Println("Wrong input!")
		return
	}

	lead := getLeadByID(lead_id)
	if lead.LeadID == -1 {
		fmt.Println("ID not found!")
		return
	}

	url := "http://localhost:8080/api/lead_conversion/" + strconv.Itoa(lead_id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	checkErr(err)

	requester := &http.Client{}
	res, err := requester.Do(req)
	checkErr(err)
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
