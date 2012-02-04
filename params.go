package stackongo

import (
	"fmt"
	"time"
	"strings"
)

type Params map[string]string

// Add adds a new parameter given by the key and value
func (p Params) Add(key string, value interface{}) {
	p[key] = fmt.Sprintf("%v", value)
}

// Del deletes the given parameter
func (p Params) Del(key string) {
	p[key] = "", false
}

// Page sets the page parameter
func (p Params) Page(page int) {
	p.Add("page", page)
}

// Pagesize sets the number of results for a page
func (p Params) Pagesize(num int) {
	p.Add("pagesize", num)
}

// Fromdate sets fromdate parameter from the given *Time reference 
func (p Params) Fromdate(date *time.Time) {
	p.Add("fromdate", date.Seconds())
}

// Todate sets todate parameter from the given *Time reference 
func (p Params) Todate(date *time.Time) {
	p.Add("todate", date.Seconds())
}

// Sort sets the field to sort results 
func (p Params) Sort(field string) {
	p.Add("sort", field)
}

// Order sets order to sort results 
func (p Params) Order(order string) {
	p.Add("order", order)
}

// Min sets the minimum value the field can take 
func (p Params) Min(num int) {
	p.Add("min", num)
}

// Max sets the maximum value the field can take 
func (p Params) Max(num int) {
	p.Add("max", num)
}

// AddVectorized adds a new parameter 
// using the given key and vectorized value (eg. tagged=ruby;java;go)
func (p Params) AddVectorized(key string, values []string) {
	p[key] = strings.Join(values, ";")
}
