package main
import (
	"fmt"
	"os"
	"image/png"
	"log"
)

func main() {
	png_filename := os.Args[1]
	png_file,err := os.Open(png_filename)
	if err != nil { log.Fatal(err) }
	png, err := png.Decode(png_file)
	if err != nil { log.Fatal(err) }
	bounds := png.Bounds()
	palette := map[uint32]bool{}
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			_,_,_,a := png.At(x,y).RGBA()
			palette[a] = true
		}
	}
	for k,_ := range palette {
		fmt.Println(k)
	}
	for k,_ := range palette {
		if k != 65535 {
			os.Exit(1)
		}
	}
	os.Exit(0)
}
