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

func getContact() {
	url := "http://localhost:8080/api/client_contact/"
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

	contacts := &ClientContactResults{}
	json := json.Unmarshal([]byte(body), contacts)
	checkErr(json)

	fmt.Print("Client ID\t")
	fmt.Print("Lead ID\t")
	fmt.Print("Client Name\t")
	fmt.Print("Client Company\t")
	fmt.Print("Client Email\t")
	fmt.Print("Client Phone\t\n")
	for _, contact := range contacts.Data {
		company := contact.ClientCompany
		phone := contact.ClientPhoneNumber
		email := contact.ClientEmailAddress

		if len(contact.ClientCompany) == 0 {
			company = "None"
		}
		if len(contact.ClientPhoneNumber) == 0 {
			phone = "None"
		}
		if len(contact.ClientEmailAddress) == 0 {
			email = "None"
		}

		fmt.Printf(
			"%d\t\t%d\t%s\t%s\t%s\t%s\n\n",
			contact.ClientID,
			contact.LeadID,
			contact.ClientName,
			company,
			email,
			phone,
		)
	}
}

func getContactByID(id int) ClientContact {
	url := "http://localhost:8080/api/client_contact/"
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

	contacts := &ClientContactResults{}
	json := json.Unmarshal([]byte(body), contacts)
	checkErr(json)

	for i := 0; i < len(contacts.Data); i++ {
		if contacts.Data[i].ClientID == id {
			return contacts.Data[i]
		}
	}
	return ClientContact{ClientID: -1}
}

func addContact(reader *bufio.Reader) {
	fmt.Print("Enter client name: ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	client_name := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter company name: ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	company_name := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter email address: ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	client_email := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter phone number: ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	phone_number := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter note: ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	note := strings.Replace(text, "\r\n", "", -1)

	body := AddClientContact{
		ClientName:         client_name,
		ClientCompany:      company_name,
		ClientEmailAddress: client_email,
		ClientPhoneNumber:  phone_number,
		ClientNote:         note,
	}

	url := "http://localhost:8080/api/client_contact/"
	body_json, err := json.Marshal(body)
	checkErr(err)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(body_json))
	checkErr(err)
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}

func updateContact(reader *bufio.Reader) {
	fmt.Print("Enter ID to delete: ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	client_id, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
	if err != nil {
		fmt.Println("Wrong input!")
		return
	}

	client := getContactByID(client_id)
	if client.ClientID == -1 {
		fmt.Println("ID not found!")
		return
	}

	fmt.Print("Enter new lead ID ")
	fmt.Print("(Old: " + client.ClientName + ", 0 if no changes): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	lead_id, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
	if err != nil {
		fmt.Println("Wrong input!")
		return
	}

	fmt.Print("Enter new client name ")
	fmt.Print("(Old: " + client.ClientName + ", empty if no changes): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	client_name := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter new company name ")
	fmt.Print("(Old: " + client.ClientCompany + "): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	company_name := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter new email address ")
	fmt.Print("(Old: " + client.ClientEmailAddress + "): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	client_email := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter new phone number ")
	fmt.Print("(Old: " + client.ClientPhoneNumber + "): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	phone_number := strings.Replace(text, "\r\n", "", -1)

	fmt.Print("Enter new note: ")
	fmt.Print("(Old: " + client.ClientNote + "): ")
	text, err = reader.ReadString('\n')
	checkErr(err)
	note := strings.Replace(text, "\r\n", "", -1)

	body := ClientContact{
		LeadID:             lead_id,
		ClientName:         client_name,
		ClientCompany:      company_name,
		ClientEmailAddress: client_email,
		ClientPhoneNumber:  phone_number,
		ClientNote:         note,
	}
	if body.LeadID == 0 {
		body.LeadID = client.LeadID
	}
	if len(body.ClientName) == 0 {
		body.ClientName = client.ClientName
	}
	if len(body.ClientCompany) == 0 {
		body.ClientCompany = client.ClientCompany
	}
	if len(body.ClientPhoneNumber) == 0 {
		body.ClientPhoneNumber = client.ClientPhoneNumber
	}
	if len(body.ClientNote) == 0 {
		body.ClientNote = client.ClientNote
	}

	url := "http://localhost:8080/api/client_contact/" + strconv.Itoa(client_id)
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

func deleteContact(reader *bufio.Reader) {
	fmt.Print("Enter ID to edit: ")
	text, err := reader.ReadString('\n')
	checkErr(err)
	client_id, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
	if err != nil {
		fmt.Println("Wrong input!")
		return
	}

	client := getContactByID(client_id)
	if client.ClientID == -1 {
		fmt.Println("ID not found!")
		return
	}

	url := "http://localhost:8080/api/client_contact/" + strconv.Itoa(client_id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	checkErr(err)

	requester := &http.Client{}
	res, err := requester.Do(req)
	checkErr(err)
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
