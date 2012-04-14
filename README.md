## Stack on Go

This is a wrapper written in Golang for [Stack Exchange API 2.0](https://api.stackexchange.com).

### Installation

Let's have a look how to get started with *Stack on Go*.

*Stack on Go* fully supports Go1.

To install the package, run:

```bash
  go get github.com/laktek/Stack-on-Go 
```
 
### Basic Usage

Once installed, you can use *Stack on Go* by importing it in your source.

```go
  import "github.com/laktek/Stack-on-Go/stackongo"
```

By default, package will be named as `stackongo`. If you want, you can give an alternate name at the import.

Stack Exchange API contains global and site specific methods. Global methods can be directly called like this:

```go
  sites, err := stackongo.AllSites(params)
```

Before calling site specific methods, you need to create a new session. A site identifier should be passed as a string (usually, it's the domain of the site).

```go
  session := stackongo.NewSession("stackoverflow")
```

Then call the methods in scope of the created session.

```go
  info, err := session.Info()
```

Most methods accept a map of parameters. There's a special `Params` type that you can use to create a parameter map. 

```go
  //set the params
  params := make(stackongo.Params)
  params.Add("filter", "total")
  params.AddVectorized("tagged", []string("go", "ruby", "java"))

  questions, err := session.AllQuestions(params)
```

If you prefer, you can pass your parameters directly in a `map[string]string` literal:

```go
  questions, err := session.AllQuestions(map[string]string{"filter": "total", "tagged": "go;ruby;java"})
```

Most methods returns a `struct` containing a collection of items and meta information (more details available in [StackExchange docs](https://api.stackexchange.com/docs/wrapper) ). You can traverse through the results to create an output:

```go
  for _, question := range questions.Items {
		fmt.Printf("%v\n", question.Title)
		fmt.Printf("Asked By: %v on %v\n", question.Owner.Display_name, time.SecondsToUTC(question.Creation_date))
		fmt.Printf("Link: %v\n\n", question.Link)
	}
```

You can use the returned meta information to make run-time decisions. For example, you can check whether there are more results and load them progressively.

```go
  if questions.Has_more {
    params.Page(page + 1)
    questions, err = session.AllQuestions(params)
	}
```

### Authentication 

Stack Exchange follows the OAuth 2.0 workflow for user authentication. *Stack on Go* includes two helper functions tailored for authentication offered by the Stack Exchange API.

`AuthURL` returns you a URL to redirect the user for authentication and `ObtainAcessToken` should be called from the handler of redirected URI to obtain the access token.

Check the following code sample, which explains the authentication flow:

```go
  func init() {
    http.HandleFunc("/", authorize)
    http.HandleFunc("/profile", profile)
  }

  func authorize(w http.ResponseWriter, r *http.Request) {
    auth_url := stackongo.AuthURL(client_id, "http://myapp.com/profile", map[string]string{"scope": "read_inbox"})

    header := w.Header()
    header.Add("Location", auth_url)
    w.WriteHeader(302)
  }

  func profile(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    access_token, err := stackongo.ObtainAccessToken(client_id, client_secret, code, "http://myapp.com/profile")

    if err != nil {
      fmt.Fprintf(w, "%v", err.String())
    } else {
      //get authenticated user
      session := stackongo.NewSession("stackoverflow")
      user, err := session.AuthenticatedUser(map[string]string{}, map[string]string{"key": client_key, "access_token": access_token["access_token"]})
      
      // do more with the authenticated user
    }

  }
```

### Using with AppEngine

If you plan to deploy your app on [Google AppEngine](http://code.google.com/appengine/docs/go/), remember to do a one slight modification in your code. Since AppEngine has a special package to fetch external URLs you have to set it as the transport method for *Stack on Go*. 

Here's how to do it:

```go

  import (
    "github.com/laktek/Stack-on-Go/stackongo"
    "appengine/urlfetch"
  )

  func main(){
    c := appengine.NewContext(r)
    ut := &urlfetch.Transport{Context: c}

    stackongo.SetTransport(ut) //set urlfetch as the transport

		session := stackongo.NewSession("stackoverflow")
    info, err := session.Info()
  }
```

### Tests and Documentation

Methods are organized into files by their return types. You can see how to use a method by checking the relavant test file. There's 100% test coverage for all methods.

*Stack on Go* implements all methods defined in the Stack Exchange API 2.0. You can use its [documentation](https://api.stackexchange.com/docs) to learn more about parameters available parameters to a method, available filters and fields you can expect in a response.

### Examples

Please check the `/examples` directory to learn various ways you can use *Stack on Go*.

### Issues & Suggestions

Please report any bugs or feature requests here:
http://github.com/laktek/Stack-on-Go/issues/

