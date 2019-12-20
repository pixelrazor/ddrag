package ddrago

import (
	"fmt"
	"net/http"
	"testing"
)

func TestVersions(t *testing.T) {
	versions, err := Client{http.DefaultClient}.Versions()
	fmt.Println("versions:", len(versions))
	if err != nil {
		t.Fatal(err)
	}
}
