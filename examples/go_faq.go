package main

import "fmt"
import "time"
import "github.com/laktek/Stack-on-Go/stackongo"

func main() {
	fmt.Printf("Frequently Asked Questions about Go:\n")

	session := stackongo.NewSession("stackoverflow")
	questions, err := session.FAQForTags([]string{"Go"}, map[string]string{})

	if err != nil {
		fmt.Printf(err.Error())
	}

	for _, question := range questions.Items {
		fmt.Printf("%v\n", question.Title)
		fmt.Printf("Asked By: %v on %v\n", question.Owner.Display_name, time.Unix(question.Creation_date, 0).UTC())
		fmt.Printf("Link: %v\n\n", question.Link)
	}
}
