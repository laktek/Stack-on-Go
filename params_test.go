package stackongo

import (
	"testing"
	"time"
)

func TestParamsAdd(t *testing.T) {
	params := make(Params)
	params.Add("key", 5)

	if params["key"] != "5" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsDel(t *testing.T) {
	params := make(Params)
	params.Add("key", "test")

	params.Del("key")

	if params["key"] == "test" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsPage(t *testing.T) {
	params := make(Params)
	params.Page(5)

	if params["page"] != "5" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsPagesize(t *testing.T) {
	params := make(Params)
	params.Pagesize(25)

	if params["pagesize"] != "25" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsFromdate(t *testing.T) {
	params := make(Params)
	date, _ := time.Parse(time.UnixDate, "Thu Feb 4 21:00:57 PST 2012")
	params.Fromdate(date)

	if params["fromdate"] != "1328389257" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsTodate(t *testing.T) {
	params := make(Params)
	date, _ := time.Parse(time.UnixDate, "Thu Feb 4 21:00:57 PST 2012")
	params.Todate(date)

	if params["todate"] != "1328389257" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsSort(t *testing.T) {
	params := make(Params)
	params.Sort("reputation")

	if params["sort"] != "reputation" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsOrder(t *testing.T) {
	params := make(Params)
	params.Order("desc")

	if params["order"] != "desc" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsMin(t *testing.T) {
	params := make(Params)
	params.Min(1)

	if params["min"] != "1" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsMax(t *testing.T) {
	params := make(Params)
	params.Max(1000)

	if params["max"] != "1000" {
		t.Errorf("%v doesn't match expectation", params)
	}
}

func TestParamsAddVectorized(t *testing.T) {
	params := make(Params)
	params.AddVectorized("tagged", []string{"ruby", "java", "go"})

	if params["tagged"] != "ruby;java;go" {
		t.Errorf("%v doesn't match expectation", params)
	}
}
