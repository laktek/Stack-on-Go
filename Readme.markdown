## Stack on Go

This is a wrapper written in Go(lang) for [Stack Exchange API 2.0](https://api.stackexchange.com).

This project is still under development, not recommended for actual use. You can still explore while it's developed, use the tests as your guide :)

### Ideas

Build awesome analyzing tools based on Stack Exchange API.

Some examples:

* Language & Tools trends
* Potential Employee tracker

### Notes

* All timestamps takes the type int64. Use `time.SecondsToUTC(elem.Creation_date)` to parse it.
