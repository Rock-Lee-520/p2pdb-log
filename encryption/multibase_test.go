package encryption

import (
	"bytes"
	"testing"

	multi "github.com/multiformats/go-multibase"
)

var sampleBytes = []byte("Decentralize everything!!!")
var encodedSamples = map[multi.Encoding]string{
	multi.Identity:          string(rune(0x00)) + "Decentralize everything!!!",
	multi.Base2:             "00100010001100101011000110110010101101110011101000111001001100001011011000110100101111010011001010010000001100101011101100110010101110010011110010111010001101000011010010110111001100111001000010010000100100001",
	multi.Base16:            "f446563656e7472616c697a652065766572797468696e67212121",
	multi.Base16Upper:       "F446563656E7472616C697A652065766572797468696E67212121",
	multi.Base32:            "birswgzloorzgc3djpjssazlwmvzhs5dinfxgoijbee",
	multi.Base32Upper:       "BIRSWGZLOORZGC3DJPJSSAZLWMVZHS5DINFXGOIJBEE",
	multi.Base32pad:         "cirswgzloorzgc3djpjssazlwmvzhs5dinfxgoijbee======",
	multi.Base32padUpper:    "CIRSWGZLOORZGC3DJPJSSAZLWMVZHS5DINFXGOIJBEE======",
	multi.Base32hex:         "v8him6pbeehp62r39f9ii0pbmclp7it38d5n6e89144",
	multi.Base32hexUpper:    "V8HIM6PBEEHP62R39F9II0PBMCLP7IT38D5N6E89144",
	multi.Base32hexPad:      "t8him6pbeehp62r39f9ii0pbmclp7it38d5n6e89144======",
	multi.Base32hexPadUpper: "T8HIM6PBEEHP62R39F9II0PBMCLP7IT38D5N6E89144======",
	multi.Base36:            "km552ng4dabi4neu1oo8l4i5mndwmpc3mkukwtxy9",
	multi.Base36Upper:       "KM552NG4DABI4NEU1OO8L4I5MNDWMPC3MKUKWTXY9",
	multi.Base58BTC:         "z36UQrhJq9fNDS7DiAHM9YXqDHMPfr4EMArvt",
	multi.Base58Flickr:      "Z36tpRGiQ9Endr7dHahm9xwQdhmoER4emaRVT",
	multi.Base64:            "mRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE",
	multi.Base64url:         "uRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE",
	multi.Base64pad:         "MRGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE=",
	multi.Base64urlPad:      "URGVjZW50cmFsaXplIGV2ZXJ5dGhpbmchISE=",
	//	multi.Base256Emoji:      "🚀💛✋💃✋😻😈🥺🤤🍀🌟💐✋😅✋💦✋🥺🏃😈😴🌟😻😝👏👏👏",
}

func testEncode(t *testing.T, encoding multi.Encoding, bytes []byte, expected string) {
	actual, err := BaseEncode(encoding, bytes)
	if err != nil {
		t.Error(err)
		return
	}
	if actual != expected {
		t.Errorf("encoding failed for %c (%d / %s), expected: %s, got: %s", encoding, encoding, multi.EncodingToStr[encoding], expected, actual)
	}
}

func testDecode(t *testing.T, expectedEncoding multi.Encoding, expectedBytes []byte, data string) {
	actualEncoding, actualBytes, err := BaseDecode(data)
	if err != nil {
		t.Error(err)
		return
	}
	if actualEncoding != expectedEncoding {
		t.Errorf("wrong encoding code, expected: %c (%d), got %c (%d)", expectedEncoding, expectedEncoding, actualEncoding, actualEncoding)
	}
	if !bytes.Equal(actualBytes, expectedBytes) {
		t.Errorf("decoding failed for %c (%d), expected: %v, got %v", actualEncoding, actualEncoding, expectedBytes, actualBytes)
	}
}

func TestBaseEncode(t *testing.T) {
	for encoding := range multi.EncodingToStr {
		testEncode(t, encoding, sampleBytes, encodedSamples[encoding])
	}
}

func TestBaseDecode(t *testing.T) {
	for encoding := range multi.EncodingToStr {
		testDecode(t, encoding, sampleBytes, encodedSamples[encoding])
	}
}

func TestBaseEnCodeAndBaseDecode(t *testing.T) {
	var str = "Decentralize everything!!!"
	var Bytes = []byte(str)
	expected := "birswgzloorzgc3djpjssazlwmvzhs5dinfxgoijbee"
	actual, err := BaseEncode(98, Bytes)
	if err != nil {
		t.Error(err)
		return
	}
	//debug.Dump(actual)

	if actual != expected {
		t.Errorf("encoding failed for %c (%d / %s), expected: %s, got: %s", 98, 98, multi.EncodingToStr[98], expected, actual)
	}
	var result []byte
	encoding, result, err := BaseDecode(expected)
	if err != nil {
		t.Error(err)
		return
	}

	// debug.Dump(result)
	// debug.Dump(encoding)
	//byte to string
	resultToString := string(result[:])
	//debug.Dump(resultToString)
	if resultToString != str {
		t.Errorf("decoding failed for %c (%d / %s), expected: %s, got: %s", 98, 98, multi.EncodingToStr[encoding], str, resultToString)
	}
}
