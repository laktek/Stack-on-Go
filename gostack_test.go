package gostack

import (
	"testing"
	"fmt"
)

func TestGetInfo(t *testing.T) {
	site := NewSite("stackoverflow")
	info, err := site.GetInfo()

	if err != nil {
		t.Error(err.String())
	}

	print(info.Total_questions)
}

func TestGetPrivileges(t *testing.T) {
	site := NewSite("stackoverflow")
	privileges, err := site.GetPrivileges()

	if err != nil {
		t.Error(err.String())
	}

	for _, elem := range privileges {
		fmt.Printf("short description: %v \n", elem.Short_description)
		fmt.Printf("description: %v \n", elem.Description)
		fmt.Printf("reputation: %v \n", elem.Reputation)
	}
}

func TestErrorRequest(t *testing.T) {
	site := NewSite("stackoverflo")
	_, err := site.GetInfo()

	if err != nil {
		t.Error(err.String())
	}
}
