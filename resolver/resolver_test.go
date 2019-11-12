package resolver

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type ResolveAddressTestCases struct {
	domainToResolve string
	expectedAddress common.Address
}

type ResolveContentTestCase struct {
	domainToResolve string
	expectedHash    common.Hash
}

func TestResolveDomainAddress(t *testing.T) {
	// if any of these fail, verify their correctness in RNS manager and update if necessary
	testCases := []ResolveAddressTestCases{
		{"alecavallero.rsk", common.HexToAddress("0xa78c937844b27bec024f042dcbe5b85d2b7344f6")},
		{"martin.rsk", common.HexToAddress("0xfb530616391cb526387bad651594bc21a77d3dfe")},
		{"pedro.rsk", common.HexToAddress("0x0164be16739135950c2fea0e75c98123f7ca06cf")},
		{"ny.consensus.rsk", common.HexToAddress("0xdbb8fd0a18fd84ba548a7e00e86465fe3de869f8")},
		{"marcelosdomain.rsk", common.HexToAddress("0xfF33bC3B7324C2A808A9D415935f8D991E6C406c")},
		{"unregistered.rsk", common.HexToAddress("0x0000000000000000000000000000000000000000")},
		{"invalidom.ain", common.HexToAddress("0x0000000000000000000000000000000000000000")},
	}
	resolveDomainAddressTestCases(t, testCases)
}

func resolveDomainAddressTestCases(t *testing.T, testCases []ResolveAddressTestCases) {
	var emptyAddress common.Address
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.domainToResolve), func(t *testing.T) {
			resolvedAddress, resolutionError := ResolveDomainAddress(testCase.domainToResolve)
			if testCase.expectedAddress != resolvedAddress {
				t.Errorf("Expected address %v and got %v.", testCase.expectedAddress, resolvedAddress)
			}
			if testCase.expectedAddress == emptyAddress {
				if resolutionError == nil {
					t.Errorf("Expected a non-nil address resolution error when resolving domain %v.", testCase.domainToResolve)
				}
			} else {
				if resolutionError != nil {
					t.Errorf("Expected a nil address resolution error when resolving domain %v.", testCase.domainToResolve)
				}
			}
		})
	}
}
func TestResolveDomainContent(t *testing.T) {
	// if any of these fail, verify their correctness in RNS manager and update if necessary
	testCases := []ResolveContentTestCase{
		{"vojtech.rsk", common.HexToHash("0x625f47dcda50ad052c620d2f63bd8ffc14f1184833b2f11876e21dc02df393f7")},
		{"marcelosdomain.rsk", common.HexToHash("0x88ced8ba8e9396672840b47e332b33d6679d9962d80cf340d3cf615db23d4e07")},
		{"unregistered.rsk", common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
		{"invalidom.ain", common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")},
	}
	resolveDomainContentTestCases(t, testCases)
}

func resolveDomainContentTestCases(t *testing.T, testCases []ResolveContentTestCase) {
	var emptyContent [32]byte
	for _, testCase := range testCases {
		t.Run(fmt.Sprint(testCase.domainToResolve), func(t *testing.T) {
			resolvedContent, resolutionError := ResolveDomainContent(testCase.domainToResolve)
			if testCase.expectedHash != resolvedContent {
				t.Errorf("Expected hash %v and got %v.", testCase.expectedHash, resolvedContent)
			}
			if testCase.expectedHash == emptyContent {
				if resolutionError == nil {
					t.Errorf("Expected a non-nil content resolution error when resolving domain %v.", testCase.domainToResolve)
				}
			} else {
				if resolutionError != nil {
					t.Errorf("Expected a nil content resolution error when resolving domain %v.", testCase.domainToResolve)
				}
			}
		})
	}
}
