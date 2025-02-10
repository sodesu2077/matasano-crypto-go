package set1

import (
	"strings"
	"testing"

	"github.com/sodesu2077/matasano-crypto-go/utils"
)

func TestHexToBase64(t *testing.T) {
	result, _ := HexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if result != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Fatal("Can't convert hex into base64")
	}
}

func TestFixedXOR(t *testing.T) {
	result, _ := FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if result != "746865206b696420646f6e277420706c6179" {
		t.Fatal("Can't XOR hex strings")
	}
}

func TestSingleByteXORCypher(t *testing.T) {
	result, _, _ := SingleByteXORCypher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if string(result) != "Cooking MC's like a pound of bacon" {
		t.Fatal("Can't decrypt single byte XOR cypher")
	}
}

func TestDetectSingleCharXOR(t *testing.T) {
	result, _, _ := DetectSingleCharXOR("../../assets/data/set1.04.txt")
	if string(result) != "Now that the party is jumping\n" {
		t.Fatal("Can't detect single char XOR")
	}
}

func TestRepeatingKeyXOR(t *testing.T) {
	result, _ := RepeatingKeyXOR("ICE", "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal")
	hex, _ := utils.BytestoBase16(result)
	if hex != "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20690a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
		t.Fatal("Can't perform repeating key XOR")
	}
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
	result := BreakRepeatingKeyXOR("../../assets/data/set1.06.txt")
	if string(result) != "Terminator X: Bring the noise" {
		t.Fatal("Can't break repeating key XOR")
	}
}

func TestDecryptAES(t *testing.T) {
	result := DecryptAES("../../assets/data/set1.07.txt", []byte("YELLOW SUBMARINE"))
	if !strings.Contains(string(result), "I'm back and I'm ringin' the bell") {
		t.Fatal("Can't decrypt AES in ECB mode")
	}
}

func TestDetectAES(t *testing.T) {
	result := DetectAES("../../assets/data/set1.08.txt")
	if !strings.Contains(string(result), "d880619740a8a19b7840a8a31c810a3d") {
		t.Fatal("Can't detect AES in ECB mode")
	}
}