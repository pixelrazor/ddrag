package ddrago

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestVersions(t *testing.T) {
	versions, err := Versions()
	fmt.Println("versions:", len(versions))
	if err != nil {
		t.Fatal(err)
	}
	languages, err := Languages()
	fmt.Println("languages:", languages)
	if err != nil {
		t.Fatal(err)
	}
	champs, err := Champions(English_UnitedStates, versions[0])
	fmt.Println("champs:", len(champs))
	if err != nil {
		t.Fatal(err)
	}
	for id, c := range champs {
		fmt.Println(id)
		//fmt.Println(c.Passive.Image.Full)
		img, err := c.Spells[0].Image.Full.Fetch(versions[0])
		if err != nil {
			t.Fatal(err)
		}
		f, _ := os.Create(string(id) + ".png")
		png.Encode(f, img)
		f.Close()
	}
}
