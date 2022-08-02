package escpos_turkish

import (
	"fmt"
)

func onlyDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// Barkod Bölümü.

// HRI karakterlerinin konumunu ayarlar
// 0: Yazılmayacak
// 1: Barkodun üstünde
// 2: Barkodun Altında
// 3: İki Tarafta da yazılacak.
func (e *Escpos) HRIPosition(p uint8) {
	if p > 3 {
		p = 0
	}
	e.WriteByte([]byte{gs, 'H', p})
}

// HRI yazı tipini ikisinden birine ayarlar.
// false: Font A (12x24) veya
// true: Font B (9x24)
func (e *Escpos) HRIFont(p bool) {
	e.WriteByte([]byte{gs, 'f', boolToByte(p)})
}

// Bir barkodun yüksekliğini ayarlar. Varsayılan 162'dir.
func (e *Escpos) BarcodeHeight(p uint8) {
	e.WriteByte([]byte{gs, 'f', p})
}

// Barkod için yatay boyutu ayarlar. Varsayılan 3'tür. 2 ile 6 arasında olmalıdır.
func (e *Escpos) BarcodeWidth(p uint8) {
	if p < 2 {
		p = 2
	}
	if p > 6 {
		p = 6
	}
	e.WriteByte([]byte{gs, 'h', p})
}

// Bir UPC Barkodu yazdırır. Gelen veri yalnızca sayısal karakterlerden oluşabilir ve uzunluğu 11 veya 12 olmalıdır
func (e *Escpos) UPCA(code string) (int, error) {
	if len(code) != 11 && len(code) != 12 {
		return 0, fmt.Errorf("Gelen verinin uzunluğu 11 ile 12 arasında olmalıdır")
	}
	if !onlyDigits(code) {
		return 0, fmt.Errorf("Gelen veri yalnızca sayısal karakterler içerebilir")
	}
	byteCode := append([]byte(code), 0)
	e.WriteByte(append([]byte{gs, 'k', 0}, byteCode...))
	return 1, nil
}

// Bir UPCE Barkodu yazdırır. Gelen veri yalnızca sayısal karakterler olabilir ve 11 veya 12 uzunluğunda olmalıdır
func (e *Escpos) UPCE(code string) (int, error) {
	if len(code) != 11 && len(code) != 12 {
		return 0, fmt.Errorf("Gelen verinin uzunluğu 11 ile 12 arasında olmalıdır")
	}
	if !onlyDigits(code) {
		return 0, fmt.Errorf("Gelen veri yalnızca sayısal karakterler içerebilir")
	}
	byteCode := append([]byte(code), 0)
	e.WriteByte(append([]byte{gs, 'k', 1}, byteCode...))
	return 1, nil
}

// Bir EAN13 Barkodu yazdırır. Gelen veri yalnızca sayısal karakterlerden oluşabilir ve uzunluğu 12 veya 13 olmalıdır
func (e *Escpos) EAN13(code string) (int, error) {
	if len(code) != 12 && len(code) != 13 {
		return 0, fmt.Errorf("Gelen verinin uzunluğu 11 ile 12 arasında olmalıdır")
	}
	if !onlyDigits(code) {
		return 0, fmt.Errorf("Gelen veri yalnızca sayısal karakterler içerebilir")
	}
	byteCode := append([]byte(code), 0)
	e.WriteByte(append([]byte{gs, 'k', 2}, byteCode...))
	return 1, nil
}

// Bir EAN8 Barkodu yazdırır. Gelen veri  yalnızca sayısal karakterler olabilir ve 7 veya 8 uzunluğunda olmalıdır
func (e *Escpos) EAN8(code string) (int, error) {
	if len(code) != 7 && len(code) != 8 {
		return 0, fmt.Errorf("Gelen verinin uzunluğu 7 ile 8 arasında olmalıdır")
	}
	if !onlyDigits(code) {
		return 0, fmt.Errorf("Gelen veri yalnızca sayısal karakterler içerebilir")
	}
	byteCode := append([]byte(code), 0)
	e.WriteByte(append([]byte{gs, 'k', 3}, byteCode...))
	return 1, nil
}
