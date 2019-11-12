package utils

import (
	"fmt"
	"testing"
)

type HashDomainTestCase struct {
	domainToHash      string
	expectedByteArray []byte
}

// source: https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
func byteSlicesAreEqual(a, b []byte) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestHashDomain(t *testing.T) {
	// examples taken from official ENS documentation and EIP 137, then transformed to byte slices
	testCases := []HashDomainTestCase{
		{"eth", []byte{147, 205, 235, 112, 139, 117, 69, 220, 102, 142, 185, 40, 1, 118, 22, 157, 28, 51, 207, 216, 237, 111, 4, 105, 10, 11, 204, 136, 169, 63, 196, 174}},
		{"foo.eth", []byte{222, 155, 9, 253, 124, 95, 144, 30, 35, 163, 241, 159, 236, 197, 72, 40, 233, 200, 72, 83, 152, 1, 232, 101, 145, 189, 152, 1, 176, 25, 248, 79}},
		{"alice.eth", []byte{120, 113, 146, 252, 83, 120, 204, 50, 170, 149, 109, 223, 222, 219, 242, 107, 36, 232, 215, 142, 64, 16, 154, 221, 14, 234, 44, 26, 1, 44, 61, 236}},
	}
	hashDomainTestCases(t, testCases)
}

func hashDomainTestCases(t *testing.T, testCases []HashDomainTestCase) {
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.domainToHash), func(t *testing.T) {
			hashedAddress := DomainToHashedByteArray(testCase.domainToHash)
			if !(byteSlicesAreEqual(testCase.expectedByteArray, hashedAddress[:])) {
				t.Errorf("Expected byte array %v and got %v.", testCase.expectedByteArray, hashedAddress)
			}
		})
	}
}
