// SPDX-License-Identifier: BlueOak-1.0.0
// Copyright Russell Hernandez Ruiz <qrpnxz@hyperlife.xyz>

// Command wkdhash takes a mail user as its first argument and
// prints the base32-encoded sha1 hash in a line as defined at
// https://tools.ietf.org/html/draft-koch-openpgp-webkey-service-11
//
// Example:
// 	 $ wkdhash test
// 	iffe93qcsgp4c8ncbb378rxjo6cn9q6u
// 	 $ wkdhash TEST
// 	iffe93qcsgp4c8ncbb378rxjo6cn9q6u
// 	 $ wkdhash lest
// 	4ykxa7wnwasr7goc634k6w1ej1ib8zge
package main

import (
	"crypto/sha1"
	"encoding/base32"
	"os"
)

// "Z-Base-32 method as described in [RFC6189], section 5.1.6."
// https://tools.ietf.org/html/draft-koch-openpgp-webkey-service-11#section-3.1
// https://tools.ietf.org/html/rfc6189#section-5.1.6
var zBase32 = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769")

func main() {
	localPart := []byte(os.Args[1])

	// "...all upper-case ASCII characters in a User ID are mapped
	// to lowercase.  Non-ASCII characters are not changed."
	for i, v := range localPart {
		if 'A' <= v && v <= 'Z' {
			localPart[i] = v + ('a' - 'A')
		}
	}

	// "The so mapped local-part is hashed using the SHA-1
	// algorithm."
	hash := sha1.Sum(localPart)

	// "The resulting 160 bit digest is encoded using the Z-Base-32
	// method as described in [RFC6189], section 5.1.6."
	encLen := zBase32.EncodedLen(len(hash))
	encHash := make([]byte, encLen+1)
	zBase32.Encode(encHash, hash[:])

	encHash[encLen] = '\n'
	_, err := os.Stdout.Write(encHash)
	if err != nil {
		panic(err)
	}
}
