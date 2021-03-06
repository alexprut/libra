// Copyright (c) The Libra Core Contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	stdlib "testing/librastdlib"
	libra "testing/libratypes"
)

func main() {
	token := &libra.TypeTag__Struct{
		Value: libra.StructTag{
			Address: libra.AccountAddress(
				[16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			),
			Module:     libra.Identifier("LBR"),
			Name:       libra.Identifier("LBR"),
			TypeParams: []libra.TypeTag{},
		},
	}
	payee := libra.AccountAddress(
		[16]uint8{0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22},
	)
	amount := uint64(1_234_567)
	script := stdlib.EncodePeerToPeerWithMetadataScript(token, payee, amount, []uint8{}, []uint8{})

	call, err := stdlib.DecodeScript(&script)
	if err != nil {
		panic(fmt.Sprintf("failed to decode script: %v", err))
	}
	payment := call.(*stdlib.ScriptCall__PeerToPeerWithMetadata)
	if payment.Amount != amount || payment.Payee != payee {
		panic("wrong script content")
	}

	bytes, err := script.LcsSerialize()
	if err != nil {
		panic("failed to serialize")
	}
	for _, b := range bytes {
		fmt.Printf("%d ", b)
	}
	fmt.Printf("\n")
}
