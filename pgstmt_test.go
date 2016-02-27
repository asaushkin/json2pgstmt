// pgstmt_test
package json2pgstmt

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestMain(t *testing.T) {
	var m JsonRequest

	s := `{"Type":"Мой тип","Login":"Логин","Password":"Пароль","Schema":"Схема","Name":"Имя","Limit":10,"Offset":10,"Order":"Порядок","Tz":"MSK","Lang":"RU","Datasource":"main","Ctxn":"","Ctxa":"","Env":"","Params":[{"Name":"","Value":"","Type":"","Order":""}]}`

	b1 := []byte(s)
	json.Unmarshal(b1, &m)

	b2, _ := json.Marshal(m)

	if !bytes.Equal(b2, b1) {
		t.Error()
	}
}
