package main

import "fmt"
import "time"
import "github.com/laktek/Stack-on-Go/stackongo"

func main() {
	tags := []string{"Go", "Erlang", "Ruby", "Python", "JavaScript"}

	fmt.Printf("No.of Questions for Technology (Last 30 days):\n")

	from_date := time.Now().Unix() - (60 * 60 * 24 * 30)

	session := stackongo.NewSession("stackoverflow")

	//set the common params
	params := make(stackongo.Params)
	params.Add("filter", "total")
	params.Add("fromdate", from_date)

	for _, tag := range tags {
		params.Add("tagged", tag)
		results, err := session.AllQuestions(params)

		if err != nil {
			fmt.Printf(err.Error())
		}

		fmt.Printf("%v  %v\n", tag, results.Total)

	}
}
