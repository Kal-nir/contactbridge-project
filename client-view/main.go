package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ClientContactResults struct {
	Data []ClientContact
}

type AddClientContact struct {
	ClientEmailAddress string `json:"client_email_address"`
	ClientPhoneNumber  string `json:"client_phone_number"`
	ClientNote         string `json:"client_note"`
	ClientName         string `json:"client_name"`
	ClientCompany      string `json:"client_company"`
}

type ClientContact struct {
	ClientID           int    `json:"client_id"`
	LeadID             int    `json:"lead_id"`
	ClientEmailAddress string `json:"client_email_address"`
	ClientPhoneNumber  string `json:"client_phone_number"`
	ClientNote         string `json:"client_note"`
	ClientName         string `json:"client_name"`
	ClientCompany      string `json:"client_company"`
}

type LeadConversionResults struct {
	Data []LeadConversion
}

type ViewLeadResults struct {
	Data []ViewLead
}

type AddLeadConversion struct {
	ConversionStatus  string `json:"conversion_status"`
	ConversionSource  string `json:"conversion_source"`
	ConversionRemarks string `json:"conversion_remarks"`
}

type LeadConversion struct {
	LeadID            int    `json:"lead_id"`
	ConversionStatus  string `json:"conversion_status"`
	ConversionSource  string `json:"conversion_source"`
	ConversionRemarks string `json:"conversion_remarks"`
}

type ViewLead struct {
	LeadID            int    `json:"lead_id"`
	ClientName        string `json:"client_name"`
	ClientCompany     string `json:"client_company"`
	ConversionStatus  string `json:"conversion_status"`
	ConversionSource  string `json:"conversion_source"`
	ConversionRemarks string `json:"conversion_remarks"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		CallClear()
		fmt.Println("ContactBridge Client View")
		fmt.Println("[1] Contact Management")
		fmt.Println("[2] Lead Management")
		fmt.Println("[3] Exit (Or Ctrl + Z)")
		fmt.Print("-> ")
		text, err := reader.ReadString('\n')
		checkErr(err)
		option, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
		if err != nil {
			continue
		}

		if option == 1 {
			contactManagement(reader)
			continue
		}

		if option == 2 {
			leadManagement(reader)
			continue
		}

		if option == 3 {
			break
		}
	}
}

func contactManagement(reader *bufio.Reader) {
	CallClear()
	for {
		fmt.Println("\n[1] Get Contact")
		fmt.Println("[2] Add Contact")
		fmt.Println("[3] Update Contact")
		fmt.Println("[4] Delete Contact")
		fmt.Println("[5] Go Back")
		fmt.Print("->")

		text, err := reader.ReadString('\n')
		checkErr(err)
		option, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
		if err != nil {
			continue
		}

		if option == 1 {
			getContact()
		}

		if option == 2 {
			addContact(reader)
		}

		if option == 3 {
			updateContact(reader)
		}

		if option == 4 {
			deleteContact(reader)
		}

		if option == 5 {
			break
		}
	}
}

func leadManagement(reader *bufio.Reader) {
	CallClear()
	for {
		fmt.Println("\n[1] Get Lead")
		fmt.Println("[2] Add Lead")
		fmt.Println("[3] Update Lead")
		fmt.Println("[4] Delete Lead")
		fmt.Println("[5] Go Back")
		fmt.Print("->")

		text, err := reader.ReadString('\n')
		checkErr(err)
		option, err := strconv.Atoi(strings.Replace(text, "\r\n", "", -1))
		if err != nil {
			continue
		}

		if option == 1 {
			getLead()
		}

		if option == 2 {
			addLead(reader)
		}

		if option == 3 {
			updateLead(reader)
		}

		if option == 4 {
			deleteLead(reader)
		}

		if option == 5 {
			break
		}
	}
}
