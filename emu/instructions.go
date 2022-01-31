package emu

var (
	CYCLES = []int{
		1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1,
		0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1,
		2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1,
		2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4,
		2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4,
		3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4,
		3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4,
	}

	INSTRUCTIONS = map[uint8]func(*CPU){
		0x09: addHLBC,
		0x19: addHLDE,
		0x29: addHLHL,
		0x2F: cpl,
		0x39: addHLSP,
		0x80: addAB,
		0x81: addAC,
		0x82: addAD,
		0x83: addAE,
		0x84: addAH,
		0x85: addAL,
		0x86: addAHL,
		0x87: addAA,
		0x88: adcAB,
		0x89: adcAC,
		0x8A: adcAD,
		0x8B: adcAE,
		0x8C: adcAH,
		0x8D: adcAL,
		0x8E: adcAHL,
		0x8F: adcAA,
		0x90: subAB,
		0x91: subAC,
		0x92: subAD,
		0x93: subAE,
		0x94: subAH,
		0x95: subAL,
		0x96: subAHL,
		0x97: subAA,
		0x98: sbcAB,
		0x99: sbcAC,
		0x9A: sbcAD,
		0x9B: sbcAE,
		0x9C: sbcAH,
		0x9D: sbcAL,
		0x9E: sbcAHL,
		0x9F: sbcAA,
		0xA0: andAB,
		0xA1: andAC,
		0xA2: andAD,
		0xA3: andAE,
		0xA4: andAH,
		0xA5: andAL,
		0xA6: andAHL,
		0xA7: andAA,
		0xA8: xorAB,
		0xA9: xorAC,
		0xAA: xorAD,
		0xAB: xorAE,
		0xAC: xorAH,
		0xAD: xorAL,
		0xAE: xorAHL,
		0xAF: xorAA,
		0xB0: orAB,
		0xB1: orAC,
		0xB2: orAD,
		0xB3: orAE,
		0xB4: orAH,
		0xB5: orAL,
		0xB6: orAHL,
		0xB7: orAA,
		0xB8: cpAB,
		0xB9: cpAC,
		0xBA: cpAD,
		0xBB: cpAE,
		0xBC: cpAH,
		0xBD: cpAL,
		0xBE: cpAHL,
		0xBF: cpAA,
		0xC6: adi,
		0xCE: aci,
		0xD6: sui,
		0xDE: sbi,
		0xE6: ani,
		0xEE: xri,
		0xF6: ori,
		0xFE: cpi,
	}
)
