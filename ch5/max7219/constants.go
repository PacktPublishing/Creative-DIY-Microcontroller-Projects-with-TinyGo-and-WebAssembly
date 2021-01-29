package max7219

type Register byte

const (
	REG_NOOP         Register = 0x00
	REG_DIGIT0       Register = 0x01
	REG_DIGIT1       Register = 0x02
	REG_DIGIT2       Register = 0x03
	REG_DIGIT3       Register = 0x04
	REG_DIGIT4       Register = 0x05
	REG_DIGIT5       Register = 0x06
	REG_DIGIT6       Register = 0x07
	REG_DIGIT7       Register = 0x08
	REG_DECODE_MODE  Register = 0x09 // turn of for led matrix, turn on for digits
	REG_INTENSITY    Register = 0x0A
	REG_SCANLIMIT    Register = 0x0B
	REG_SHUTDOWN     Register = 0x0C // turn on for no shutdown mode
	REG_DISPLAY_TEST Register = 0x0F // turn off for no display test
)
