package main

import "fmt"
import "time"
import "github.com/laktek/Stack-on-Go/stackongo"

func main() {
	show_pages_upto := 5

	fmt.Printf("Questions tagged with Web:\n")

	session := stackongo.NewSession("stackoverflow")

	params := make(stackongo.Params)
	params.Sort("votes")
	params.Add("tagged", "web")
	params.Pagesize(100)

	total := 0

	for cur_page := 0; cur_page < show_pages_upto; cur_page++ {
		params.Set("page", cur_page+1)
		questions, err := session.AllQuestions(params)

		if err != nil {
			fmt.Printf(err.Error())
		}

		for _, question := range questions.Items {
			fmt.Printf("%v\n", question.Title)
			fmt.Printf("Asked By: %v on %v\n", question.Owner.Display_name, time.Unix(question.Creation_date, 0).UTC())
			fmt.Printf("Link: %v\n\n", question.Link)
		}

		total += len(questions.Items)

		//break the loop if there aren't any more results
		if !questions.Has_more {
			break
		}

		fmt.Printf("Total Questions Fetched: %v\n\n", total)
	}

}
