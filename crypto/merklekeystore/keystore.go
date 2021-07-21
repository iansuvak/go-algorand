// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package merklekeystore

import (
	"fmt"
	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/crypto/merklearray"
	"github.com/algorand/go-algorand/protocol"
)

type (
	//EphemeralKeys represent the possible keys inside the keystore.
	EphemeralKeys struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		SignatureAlgorithms []crypto.SignatureAlgorithm `codec:"sks,allocbound=-"`
		firstRound          uint64                      `codec:"rnd"`
	}

	// EphemeralPublicKey is a key tied to a specific round and is committed by the merklekeystore.Signer.
	EphemeralPublicKey struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		VerifyingKey crypto.VerifyingKey `codec:"pk"`
		Round        uint64              `codec:"rnd"`
		Pos          uint64              `codec:"pos"`
	}

	//Proof represent the merkle proof in each signature.
	//msgp:allocbound Proof -
	Proof []crypto.Digest

	// Signature is a byte signature on a crypto.Hashable object,
	// crypto.VerifyingKey and includes a merkle proof for the key.
	Signature struct {
		_struct              struct{} `codec:",omitempty,omitemptyarray"`
		crypto.ByteSignature `codec:"bsig"`
		Proof                `codec:"prf"`
		VKey                 EphemeralPublicKey `codec:"vkey"`
	}

	// Signer is a merkleKeyStore, contain multiple keys which can be used per round.
	Signer struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`
		// these keys are the keys used to sign in a round.
		// should be disposed of once possible.
		EphemeralKeys    EphemeralKeys `codec:"keys"`
		FirstRound       uint64        `codec:"srnd"`
		merklearray.Tree `codec:"tree"`
	}

	// Verifier Is a way to verify a Signature produced by merklekeystore.Signer.
	// it also serves as a commit over all keys contained in the merklekeystore.Signer.
	Verifier struct {
		_struct struct{} `codec:",omitempty,omitemptyarray"`

		root       crypto.Digest `codec:"r"`
		FirstValid uint64        `codec:"fv"`
	}
)

// ToBeHashed implementation means EphemeralPublicKey is crypto.Hashable.
func (e *EphemeralPublicKey) ToBeHashed() (protocol.HashID, []byte) {
	return protocol.EphemeralPK, protocol.Encode(e)
}

//Length returns the amount of disposable keys
func (d *EphemeralKeys) Length() uint64 {
	return uint64(len(d.SignatureAlgorithms))
}

// GetHash Gets the hash of the VerifyingKey tied to the signatureAlgorithm in pos.
func (d *EphemeralKeys) GetHash(pos uint64) (crypto.Digest, error) {
	ephPK := EphemeralPublicKey{
		VerifyingKey: d.SignatureAlgorithms[pos].GetSigner().GetVerifyingKey(),
		Round:        d.firstRound + pos,
		Pos:          pos,
	}
	return crypto.HashObj(&ephPK), nil
}

var errStartBiggerThanEndRound = fmt.Errorf("cannot create merkleKeyStore because end round is smaller then start round")

// New Generates a merklekeystore.Signer
// Note that the signer will have keys for the rounds  [firstValid, lastValid]
func New(firstValid, lastValid uint64, sigAlgoType crypto.AlgorithmType) (*Signer, error) {
	if firstValid > lastValid {
		return nil, errStartBiggerThanEndRound
	}

	keys := make([]crypto.SignatureAlgorithm, lastValid-firstValid+1)
	for i := range keys {
		keys[i] = *crypto.NewSigner(sigAlgoType)
	}
	ephKeys := EphemeralKeys{
		SignatureAlgorithms: keys,
		firstRound:          firstValid,
	}
	tree, err := merklearray.Build(&ephKeys)
	if err != nil {
		return nil, err
	}

	return &Signer{
		EphemeralKeys: ephKeys,
		FirstRound:    firstValid,
		Tree:          *tree,
	}, nil
}

// GetVerifier can be used to store the commitment and verifier for this signer.
func (m *Signer) GetVerifier() *Verifier {
	return &Verifier{
		root: m.Root(),
	}
}

// Sign outputs a signature + proof for the signing key.
func (m *Signer) Sign(hashable crypto.Hashable, round uint64) (Signature, error) {
	pos, err := m.getKeyPosition(round)
	if err != nil {
		return Signature{}, err
	}

	proof, err := m.Prove([]uint64{pos})
	if err != nil {
		return Signature{}, err
	}

	signer := m.EphemeralKeys.SignatureAlgorithms[pos].GetSigner()
	return Signature{
		ByteSignature: signer.Sign(hashable),
		Proof:         proof,
		VKey: EphemeralPublicKey{
			VerifyingKey: signer.GetVerifyingKey(),
			Round:        round,
			Pos:          pos,
		},
	}, nil
}

var errOutOfBounds = fmt.Errorf("cannot find signing key for given round")

func (m *Signer) getKeyPosition(round uint64) (uint64, error) {
	if round < m.FirstRound {
		return 0, errOutOfBounds
	}

	pos := round - m.FirstRound
	if pos >= uint64(len(m.EphemeralKeys.SignatureAlgorithms)) {
		return 0, errOutOfBounds
	}
	return pos, nil
}

// Verify receives a signature over a specific crypto.Hashable object, and makes certain the signature is correct.
func (v *Verifier) Verify(obj crypto.Hashable, sig Signature) error {
	isInTree := merklearray.Verify(v.root, map[uint64]crypto.Digest{sig.VKey.Pos: crypto.HashObj(&sig.VKey)}, sig.Proof)
	if isInTree != nil {
		return isInTree
	}
	return sig.VKey.VerifyingKey.GetVerifier().Verify(obj, sig.ByteSignature)
}