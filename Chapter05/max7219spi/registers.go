package max7219spi

const (
	REG_NOOP         byte = 0x00
	REG_DIGIT0       byte = 0x01
	REG_DIGIT1       byte = 0x02
	REG_DIGIT2       byte = 0x03
	REG_DIGIT3       byte = 0x04
	REG_DIGIT4       byte = 0x05
	REG_DIGIT5       byte = 0x06
	REG_DIGIT6       byte = 0x07
	REG_DIGIT7       byte = 0x08
	REG_DECODE_MODE  byte = 0x09
	REG_INTENSITY    byte = 0x0A
	REG_SCANLIMIT    byte = 0x0B
	REG_SHUTDOWN     byte = 0x0C
	REG_DISPLAY_TEST byte = 0x0F
)
