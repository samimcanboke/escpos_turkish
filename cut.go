package escpos_turkish

func (c *Escpos) FullCut() {
	c.dev.Write([]byte{esc, 0x69})
}

func (c *Escpos) PartialCut() {
	c.dev.Write([]byte{esc, 0x6D})
}
