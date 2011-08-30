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
  </body>
</html>
`

