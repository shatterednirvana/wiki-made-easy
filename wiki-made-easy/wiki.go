package wiki

import (
	"fmt"
	"http"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, mainPage)
}

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
      <h1>Krautscheid is a municipality in the district of Bitburg-Prüm, in Rhineland-Palatinate, western Germany.</h1>
      <h1><a href="http://en.wikipedia.org/wiki/Krautscheid">Learn more about this!</a></h1>
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

