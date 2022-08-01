# Escpos_Turkish (Türkçe Karakter Desteği)
ESC-POS Yazıcıları için türkçe karakter ve barkod destekli GoLang Kütüphanesidir. Kullanımı için önce kütüphaneyi indirin sonrasında ana paketinize aşağıdaki şekilde dahil edin. Yazıcınızın hangi lp portuna bağlı olduğunu mutlaka kontrol edin. Not: Root yetkisi gerektirmektedir. 

## Yükleme

```bash
go get -u github.com/samimcanboke/escpos_turkish
```

## Kullanım Örneği

```golang
package main

import (
	escpos "github.com/samimcanboke/escpos_turkish"
	"os"
)

func main() {
	f, err := os.OpenFile("/dev/usb/lp0", os.O_RDWR, 0) // terminalden kontrol edilmeli

	if err != nil {
		panic(err)
	}

	defer f.Close()

	p := escpos.New(f)
	p.Init()
    p.Charset(20) //0,2,3,4,5,19,20 olabilir.
	p.FontSize(2, 2)
	p.Font(escpos.FontB)
	p.FontAlign(escpos.AlignCenter)
	p.Writeln("Merhaba Dünya")
	p.Feed()

	p.FontSize(1, 1)
	p.Font(escpos.FontA)
	p.FontAlign(escpos.AlignLeft)
	p.Writeln("Türkçe Karakter Yazılabilir. Örnek ŞşÇçÖöÜüĞğİı... vb")

	p.FeedN(10)

	p.FullCut()
}
```
