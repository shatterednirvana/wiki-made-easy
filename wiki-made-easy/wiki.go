package wiki

import (
	"appengine"
	"appengine/urlfetch"

	"http"
	"regexp"
	"template"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://en.wikipedia.org/wiki/Special:Random")
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
	}

	body := resp.Body
	var allData []byte
	allData = make([]byte, resp.ContentLength)
	body.Read(allData)

	firstLine := getFirstLine(string(allData))
	url := getUrl(string(allData))

	// TODO - check if firstLine is empty, and if so, log it
	if firstLine == "" {
		c.Errorf("Saw empty first line for data: %s", allData)
	}

	data := map[string] string {
		"firstLine" : firstLine,
		"url" : url,
	}


	err2 := mainPageTemplate.Execute(w, data)
	if err2 != nil {
		http.Error(w, err2.String(), http.StatusInternalServerError)
	}
}

func getFirstLine(stringToRegex string) string {
	getFirstParagraphRegex := regexp.MustCompile("<p><b>([^.]*).")
	firstLine := getFirstParagraphRegex.FindString(stringToRegex)

	return removeString(firstLine, "<([^<>]*)>")
}

func getUrl(stringToRegex string) string {
	getUrlRegex := regexp.MustCompile("Retrieved from \"<a href=\".*\">.*</a>\"")
	urlInfo := getUrlRegex.FindString(stringToRegex)

	findLinkText := regexp.MustCompile(">http://.*<")
	url := findLinkText.FindString(urlInfo)
	url = removeString(url, "<")
	url = removeString(url, ">")
	return url
}

func removeString(stringToRemove, strRegex string) string {
	actualRegex := regexp.MustCompile(strRegex)
	return actualRegex.ReplaceAllString(stringToRemove, "")
}

const testData = `
not used right now
`

var mainPageTemplate = template.MustParse(mainPage, nil)

const mainPage = `
<html>
  <head>
    <title>Wiki Made Easy</title>
    <link rel="stylesheet" href="/static/css/bootstrap-1.1.0.min.css">
    <script src="static/js/jquery.js" type="text/javascript"></script>
    <script src="static/js/custom.js" type="text/javascript"></script>
  </head>
  <body>
    <div class="container">
      <br /><br /><br /><br /><br /><br /><br />
      <h1>Did you know...</h1>
      <h1>{firstLine|html}</h1>
      <h1><a href="{url|html}">Learn more about this!</a></h1>
      <h1><a href="/">Learn something else!</a></h1>
      <br /><br /><br /><br /><br /><br /><br />
      <div>
        <a href="http://code.google.com/appengine/">
          <img src="/static/img/appengine-silver-120x30.gif" alt="Powered by Google App Engine" />
        </a>
        <a href="http://golang.org">
          <img src="/static/img/Golang.png" alt="Powered by Go" />
        </a>
      </div>
    </div>
  </body>
</html>
`

