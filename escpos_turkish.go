package escpos_turkish

import (
	"golang.org/x/text/encoding"
	"io"
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

func (e *Escpos) ConvertTR(str string) string {
	e.dev.Write([]byte{0x1B, 0x74, 0b1101})
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

func (e *Escpos) Writeln(text string) {
	e.Write(text + "\n")
}

// New create new Escpos struct and set default enconding
func New(dev io.ReadWriter) *Escpos {
	escpos := &Escpos{dev: dev}
	escpos.Charset(CharsetWindows1254)

	return escpos
}
