package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

const (
	Head = `
<html>
<head>
<script type="text/JavaScript">
<!--
function timedRefresh(timeoutPeriod) {
	setTimeout("location.reload(true);",timeoutPeriod);
}
//   -->
</script>
</head>
<body onload="JavaScript:timedRefresh(4000);">`
	Tail = `
</body>
</html>`
)



func Handler(w http.ResponseWriter, req *http.Request) {
	out, _ := exec.Command("./sc.sh").Output()
	fmt.Println(req)
	fmt.Fprintf(w, "%s %s %s", Head, out, Tail)
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":26666", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
