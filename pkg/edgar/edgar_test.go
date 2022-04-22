package edgar

import (
	"fmt"
	"testing"
)

var EdgarClient *Edgar

func TestNew(t *testing.T) {
	EdgarClient = New()
}

func TestEdgar_GetSubmissions(t *testing.T) {

	TestNew(t)

	var err error

	submissions, err := EdgarClient.GetSubmissions("0001326380")
	if err != nil {
		t.Errorf("Error getting submissions: %v", err)
	}
	out := fmt.Sprintf("Cik: %s\n", submissions.Cik)
	print(out)
}

func TestEdgar_GetCompanyConcept(t *testing.T) {

	TestNew(t)

	var err error

	companyConcept, err := EdgarClient.GetCompanyConcept("0001326380")
	if err != nil {
		t.Errorf("Error getting submissions: %v", err)
	}
	out := fmt.Sprintf("Cik: %s\n", companyConcept.EntityName)
	print(out)
}
