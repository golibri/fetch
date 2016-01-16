# golibri/fetch
Get and normalize textual content from URLs. Does a few things that go's stdlib doesn't: 

- normalize URLs before executing a request (missing `http://`, beginning with `//`, ...)
- normalize response buffer to a valid UTF-8 string

# installation
`go get github.com/golibri/fetch`

# usage
````go
page, err := Get("http://blog.golang.org/go1.5")
normalized_url := page.Url
content_string := page.Body
````

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
