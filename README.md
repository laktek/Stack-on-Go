## Stack on Go

This is a wrapper written in Go(lang) for [Stack Exchange API 2.0](https://api.stackexchange.com).

This is ideal to build awesome backend apps that are based on Stack Exchange API such as notifiers, aggregators and analyzers.
*Stack on Go* is also fully compatible with Google AppEngine. 

So if you're a Stack Overflow enthusiast who'd love to learn Go, this is a good way to get a start.

### Installation

Installation of *Stack on Go* is simple. Of course you will need to have Go installed on your machine. *Stack on Go* is targeted for the current Go stable release r.60.3.

To install the package using `goinstall` run:

 ```bash
 goinstall github.com/laktek/Stack-on-Go 
 ```

Or you can [clone](https://github.com/laktek/Stack-on-Go) this repo and run:
  
  ```bash
  cd Stack-on-Go
  make install
  ```

### Basic Usage

To use *Stack on Go* first you will need to import it in your source file.

  ```go
  import "github.com/laktek/stack-on-go"
  ```

By default, package will be named as `stackongo`, but you could give an alternate name in the import.

To execute methods on specific StackExchange site, you should create a new session by passing the site name as a string.

  ```go
  session := stackongo.NewSession("stackoverflow")
  ```

Then you can call methods on the site:

  ```go
  info, err := session.Info()
  ```

For the methods that accepts parameters you can either explicitly pass in a `map[string]string` or use the special `Params` type to build the parameters map.

  ```go
  //set the params
  params := make(stackongo.Params)
  params.Add("filter", "total")
  params.AddVectorized("tagged", []string("go", "ruby", "java"))

  questions, err := session.AllQuestions(params)
  ```

Most methods returns a struct which contains a collection of items. You can traverse those items in this way:

  ```go
  for _, question := range questions.Items {
		fmt.Printf("%v\n", question.Title)
		fmt.Printf("Asked By: %v on %v\n", question.Owner.Display_name, time.SecondsToUTC(question.Creation_date))
		fmt.Printf("Link: %v\n\n", question.Link)
	}
  ```

Methods are organized into files by their return types. You can see how to use a method by checking the relavant test file. There's 100% test coverage for all tests.

*Stack on Go* implements all methods defined in the Stack Exchange API 2.0. Use its [documentation](https://api.stackexchange.com/docs) to learn more about parameters you can pass to a method, available filters and fields you can expect in a response.

### Examples

Please check the `/examples` directory to learn various ways you can use *Stack on Go*.

### Issues & Suggestions

Please report any bugs or feature requests here:
http://github.com/laktek/Stack-on-Go/issues/

