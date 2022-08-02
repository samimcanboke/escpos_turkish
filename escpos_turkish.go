package escpos_turkish

import (
	"fmt"
	"golang.org/x/text/encoding"
	"image"
	"io"
	"log"
	"os"
	"strings"
)

const (
	esc = 0x1B
	gs  = 0x1D
	lf  = 0x0A
)

type Escpos struct {
	dev io.ReadWriter
	enc *encoding.Encoder
}

// Init printer cleaning your old configs
func (e *Escpos) Init() {
	e.dev.Write([]byte{esc, 0x40})
}

// Feed skip one line of paper
func (e *Escpos) Feed() {
	e.dev.Write([]byte{lf})
}

// Feed skip n lines of paper
func (e *Escpos) FeedN(n byte) {
	e.dev.Write([]byte{esc, 0x64, n})
}

// SelfTest start self test of printer
func (e *Escpos) SelfTest() {
	e.dev.Write([]byte{gs, 0x28, 0x41, 0x02, 0x00, 0x00, 0x02})
}

// Write print text
func (e *Escpos) Write(text string) {
	str, err := e.enc.String(text)

	if err != nil {
		panic(err)
	}

	e.dev.Write([]byte(str))
}

func (e *Escpos) WriteByte(text []byte) {
	e.dev.Write(text)
}

func convertTR(str string) string {
	if strings.ContainsAny(str, "ŞşÖöÇçİıĞğÜü") {
		newStr := strings.ReplaceAll(str, "Ş", string([]byte{0x9E}))
		newStr = strings.ReplaceAll(newStr, "ş", string([]byte{0x9F}))
		newStr = strings.ReplaceAll(newStr, "ğ", string([]byte{0xA7}))
		newStr = strings.ReplaceAll(newStr, "Ğ", string([]byte{0xA6}))
		newStr = strings.ReplaceAll(newStr, "ı", string([]byte{0x8D}))
		newStr = strings.ReplaceAll(newStr, "İ", string([]byte{0x98}))
		newStr = strings.ReplaceAll(newStr, "ö", string([]byte{0x94}))
		newStr = strings.ReplaceAll(newStr, "Ö", string([]byte{0x99}))
		newStr = strings.ReplaceAll(newStr, "ü", string([]byte{0x81}))
		newStr = strings.ReplaceAll(newStr, "Ü", string([]byte{0x9A}))
		newStr = strings.ReplaceAll(newStr, "ç", string([]byte{0x87}))
		newStr = strings.ReplaceAll(newStr, "Ç", string([]byte{0x80}))
		return newStr
	} else {
		return str
	}
}
func (e *Escpos) WriteTr(str string) {
	newStr := convertTR(str)
	e.dev.Write([]byte(newStr))
}

func (e *Escpos) Writeln(text string) {
	e.Write(text + "\n")
}

// New create new Escpos struct and set default enconding
func New(dev io.ReadWriter) *Escpos {
	escpos := &Escpos{dev: dev}
	escpos.Charset(CharsetWindows1254)

	return escpos
}

func (e *Escpos) PrintImageFromFile(imgName string) {
	path, err := os.Getwd()
	if strings.Contains(imgName, path) {
		path = imgName
	} else {
		path = path + imgName
	}
	imgFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(imgFile)
	imgFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	_, _, _, _, imgData := PrintImage(img)
	e.dev.Write(imgData)
}

func (e *Escpos) PrintBarcode(typ string, textPosition uint8, textFont bool, barcodeHeight uint8, barcodeWidth uint8, data string) {
	e.HRIPosition(textPosition)
	e.HRIFont(textFont)
	e.BarcodeWidth(barcodeWidth)
	e.BarcodeHeight(barcodeHeight)
	switch typ {
	case "EAN13":
		ean13, err := e.EAN13(data)
		if err != nil {
			fmt.Println("Barkod Yazma Hatası", err)
		}
		fmt.Println(ean13)
	case "EAN8":
		ean8, err := e.EAN8(data)
		if err != nil {
			fmt.Println("Barkod Yazma Hatası", err)
		}
		fmt.Println(ean8)
	case "UPCE":
		upce, err := e.UPCE(data)
		if err != nil {
			fmt.Println("Barkod Yazma Hatası", err)
		}
		fmt.Println(upce)
	case "UPCA":
		upca, err := e.UPCA(data)
		if err != nil {
			fmt.Println("Barkod Yazma Hatası", err)
		}
		fmt.Println(upca)
	}
}
