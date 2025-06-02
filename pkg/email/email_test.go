package email

import (
	"fmt"
	"testing"
)

func TestSendCodeEmail(t *testing.T) {
	code, err := SendCodeEmail("1442578851@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(code)
}
