package util

func UnDecodeLength(srclen int) int {
	return (srclen * 3) / 4
}

func UnEncodeLength(srclen int) int {
	destlen := (srclen * 4) / 3

	if (srclen*4)%3 != 0 {
		destlen = destlen + 1
	}
	return destlen
}

func UuEnocde(src []byte) []byte {
	//	log.Printf("%X", src)
	srclen := len(src)
	destlen := UnEncodeLength(srclen)
	dest := make([]byte, destlen)
	var destIndex int = 0
	var srcIndex int = 0
	for srcIndex < srclen {
		//Encode A
		dest[destIndex] = byte(((uint32(src[srcIndex]) >> 2) & 0x3F) + 0x20)
		destIndex = destIndex + 1
		dest[destIndex] = byte(((uint32(src[srcIndex]) << 4) & 0x3F))
		srcIndex = srcIndex + 1

		if srcIndex >= srclen {
			dest[destIndex] = dest[destIndex] + 0x20
			destIndex = destIndex + 1
			break
		}

		// Encode B
		dest[destIndex] = byte(uint32(dest[destIndex]) | ((uint32(src[srcIndex]) >> 4) & 0x3F))
		dest[destIndex] = dest[destIndex] + 0x20
		destIndex = destIndex + 1
		dest[destIndex] = byte(((uint32(src[srcIndex]) << 2) & 0x3F))
		srcIndex = srcIndex + 1
		if srcIndex >= srclen {
			dest[destIndex] = dest[destIndex] + 0x20
			destIndex = destIndex + 1
			break
		}

		// Encode C
		dest[destIndex] = byte(uint32(dest[destIndex]) | ((uint32(src[srcIndex]) >> 6) & 0x3F))
		dest[destIndex] = dest[destIndex] + 0x20
		destIndex = destIndex + 1
		dest[destIndex] = byte((uint32(src[srcIndex]) & 0x3F) + 0x20)
		srcIndex = srcIndex + 1
		destIndex = destIndex + 1

	}

	return dest

}

func UuDecode(src []byte) []byte {
	srclen := len(src)
	destlen := UnDecodeLength(srclen)
	dest := make([]byte, destlen)

	var destIndex int = 0
	var srcIndex int = 0

	for srcIndex < srclen {
		// Decode A
		dest[destIndex] = byte((((uint32(src[srcIndex]) - 0x20) << 2) & 0xFC))
		srcIndex = srcIndex + 1
		dest[destIndex] = byte(uint32(dest[destIndex]) | (((uint32(src[srcIndex]) - 0x20) >> 4) & 0x03))
		destIndex = destIndex + 1
		if (srcIndex + 1) >= srclen {
			break
		}

		// Decode B
		dest[destIndex] = byte((((uint32(src[srcIndex]) - 0x20) << 4) & 0xF0))
		srcIndex = srcIndex + 1
		dest[destIndex] = byte(uint32(dest[destIndex]) | (((uint32(src[srcIndex]) - 0x20) >> 2) & 0x0F))
		destIndex = destIndex + 1
		if (srcIndex + 1) >= srclen {
			break
		}

		// Decode C
		dest[destIndex] = byte((((uint32(src[srcIndex]) - 0x20) << 6) & 0xC0))
		srcIndex = srcIndex + 1
		dest[destIndex] = byte(uint32(dest[destIndex]) | ((uint32(src[srcIndex]) - 0x20) & 0x3F))
		destIndex = destIndex + 1
		srcIndex = srcIndex + 1

	}

	return dest

}
