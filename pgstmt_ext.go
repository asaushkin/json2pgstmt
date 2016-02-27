// pgstmt_ext
package json2pgstmt

import (
	"fmt"
)

type AuthParam struct {
	Login    string
	Password string
	Ctxn     string
	Ctxa     string
}

type ConnetcionParam struct {
	Tz   string
	Lang string

	Datasource string
}

func main() {
	fmt.Println("Hello World!")
}
