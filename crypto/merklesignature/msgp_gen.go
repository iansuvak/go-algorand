package merklesignature

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"

	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/crypto/merklearray"
)

// The following msgp objects are implemented in this file:
// Commitment
//      |-----> (*) MarshalMsg
//      |-----> (*) CanMarshalMsg
//      |-----> (*) UnmarshalMsg
//      |-----> (*) UnmarshalValidateMsg
//      |-----> (*) CanUnmarshalMsg
//      |-----> (*) Msgsize
//      |-----> (*) MsgIsZero
//      |-----> CommitmentMaxSize()
//
// KeyRoundPair
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) UnmarshalValidateMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//       |-----> KeyRoundPairMaxSize()
//
// Secrets
//    |-----> (*) MarshalMsg
//    |-----> (*) CanMarshalMsg
//    |-----> (*) UnmarshalMsg
//    |-----> (*) UnmarshalValidateMsg
//    |-----> (*) CanUnmarshalMsg
//    |-----> (*) Msgsize
//    |-----> (*) MsgIsZero
//    |-----> SecretsMaxSize()
//
// Signature
//     |-----> (*) MarshalMsg
//     |-----> (*) CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) UnmarshalValidateMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> (*) Msgsize
//     |-----> (*) MsgIsZero
//     |-----> SignatureMaxSize()
//
// SignerContext
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) UnmarshalValidateMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//       |-----> SignerContextMaxSize()
//
// Verifier
//     |-----> (*) MarshalMsg
//     |-----> (*) CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) UnmarshalValidateMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> (*) Msgsize
//     |-----> (*) MsgIsZero
//     |-----> VerifierMaxSize()
//

// MarshalMsg implements msgp.Marshaler
func (z *Commitment) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (*z)[:])
	return
}

func (_ *Commitment) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Commitment)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Commitment) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (*z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

func (z *Commitment) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *Commitment) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *Commitment) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Commitment)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Commitment) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (MerkleSignatureSchemeRootSize * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Commitment) MsgIsZero() bool {
	return (*z) == (Commitment{})
}

// MaxSize returns a maximum valid message size for this message type
func CommitmentMaxSize() (s int) {
	// Calculating size of array: z
	s = msgp.ArrayHeaderSize + ((MerkleSignatureSchemeRootSize) * (msgp.ByteSize))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *KeyRoundPair) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Key == nil {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Round == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "key"
			o = append(o, 0xa3, 0x6b, 0x65, 0x79)
			if (*z).Key == nil {
				o = msgp.AppendNil(o)
			} else {
				o = (*z).Key.MarshalMsg(o)
			}
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "rnd"
			o = append(o, 0xa3, 0x72, 0x6e, 0x64)
			o = msgp.AppendUint64(o, (*z).Round)
		}
	}
	return
}

func (_ *KeyRoundPair) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*KeyRoundPair)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *KeyRoundPair) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0003 string
	var zb0004 bool
	var zb0002 bool
	_ = zb0003
	_ = zb0004
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if validate {
			err = &msgp.ErrNonCanonical{}
			return
		}
		if zb0001 > 0 {
			zb0001--
			(*z).Round, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Round")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				(*z).Key = nil
			} else {
				if (*z).Key == nil {
					(*z).Key = new(crypto.FalconSigner)
				}
				bts, err = (*z).Key.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Key")
					return
				}
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = KeyRoundPair{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "rnd":
				if validate && zb0004 && "rnd" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).Round, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Round")
					return
				}
				zb0003 = "rnd"
			case "key":
				if validate && zb0004 && "key" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					(*z).Key = nil
				} else {
					if (*z).Key == nil {
						(*z).Key = new(crypto.FalconSigner)
					}
					bts, err = (*z).Key.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "Key")
						return
					}
				}
				zb0003 = "key"
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
			zb0004 = true
		}
	}
	o = bts
	return
}

func (z *KeyRoundPair) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *KeyRoundPair) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *KeyRoundPair) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*KeyRoundPair)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *KeyRoundPair) Msgsize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 4
	if (*z).Key == nil {
		s += msgp.NilSize
	} else {
		s += (*z).Key.Msgsize()
	}
	return
}

// MsgIsZero returns whether this is a zero value
func (z *KeyRoundPair) MsgIsZero() bool {
	return ((*z).Round == 0) && ((*z).Key == nil)
}

// MaxSize returns a maximum valid message size for this message type
func KeyRoundPairMaxSize() (s int) {
	s = 1 + 4 + msgp.Uint64Size + 4
	s += crypto.FalconSignerMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Secrets) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(3)
	var zb0002Mask uint8 /* 6 bits */
	if (*z).SignerContext.FirstValid == 0 {
		zb0002Len--
		zb0002Mask |= 0x8
	}
	if (*z).SignerContext.KeyLifetime == 0 {
		zb0002Len--
		zb0002Mask |= 0x10
	}
	if (*z).SignerContext.Tree.MsgIsZero() {
		zb0002Len--
		zb0002Mask |= 0x20
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x8) == 0 { // if not empty
			// string "fv"
			o = append(o, 0xa2, 0x66, 0x76)
			o = msgp.AppendUint64(o, (*z).SignerContext.FirstValid)
		}
		if (zb0002Mask & 0x10) == 0 { // if not empty
			// string "iv"
			o = append(o, 0xa2, 0x69, 0x76)
			o = msgp.AppendUint64(o, (*z).SignerContext.KeyLifetime)
		}
		if (zb0002Mask & 0x20) == 0 { // if not empty
			// string "tree"
			o = append(o, 0xa4, 0x74, 0x72, 0x65, 0x65)
			o = (*z).SignerContext.Tree.MarshalMsg(o)
		}
	}
	return
}

func (_ *Secrets) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Secrets)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Secrets) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0002 int
	var zb0004 string
	var zb0005 bool
	var zb0003 bool
	_ = zb0004
	_ = zb0005
	zb0002, zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0002, zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if validate {
			err = &msgp.ErrNonCanonical{}
			return
		}
		if zb0002 > 0 {
			zb0002--
			(*z).SignerContext.FirstValid, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "FirstValid")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			(*z).SignerContext.KeyLifetime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "KeyLifetime")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			bts, err = (*z).SignerContext.Tree.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Tree")
				return
			}
		}
		if zb0002 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0002)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0003 {
			(*z) = Secrets{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "fv":
				if validate && zb0005 && "fv" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).SignerContext.FirstValid, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "FirstValid")
					return
				}
				zb0004 = "fv"
			case "iv":
				if validate && zb0005 && "iv" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).SignerContext.KeyLifetime, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "KeyLifetime")
					return
				}
				zb0004 = "iv"
			case "tree":
				if validate && zb0005 && "tree" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).SignerContext.Tree.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tree")
					return
				}
				zb0004 = "tree"
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
			zb0005 = true
		}
	}
	o = bts
	return
}

func (z *Secrets) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *Secrets) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *Secrets) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Secrets)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Secrets) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 3 + msgp.Uint64Size + 5 + (*z).SignerContext.Tree.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Secrets) MsgIsZero() bool {
	return ((*z).SignerContext.FirstValid == 0) && ((*z).SignerContext.KeyLifetime == 0) && ((*z).SignerContext.Tree.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func SecretsMaxSize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 3 + msgp.Uint64Size + 5 + merklearray.TreeMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Signature) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(4)
	var zb0001Mask uint8 /* 5 bits */
	if (*z).VectorCommitmentIndex == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Proof.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Signature.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).VerifyingKey.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "idx"
			o = append(o, 0xa3, 0x69, 0x64, 0x78)
			o = msgp.AppendUint64(o, (*z).VectorCommitmentIndex)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "prf"
			o = append(o, 0xa3, 0x70, 0x72, 0x66)
			o = (*z).Proof.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "sig"
			o = append(o, 0xa3, 0x73, 0x69, 0x67)
			o = (*z).Signature.MarshalMsg(o)
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "vkey"
			o = append(o, 0xa4, 0x76, 0x6b, 0x65, 0x79)
			o = (*z).VerifyingKey.MarshalMsg(o)
		}
	}
	return
}

func (_ *Signature) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Signature)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Signature) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0003 string
	var zb0004 bool
	var zb0002 bool
	_ = zb0003
	_ = zb0004
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if validate {
			err = &msgp.ErrNonCanonical{}
			return
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Signature.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Signature")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).VectorCommitmentIndex, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VectorCommitmentIndex")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Proof.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Proof")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).VerifyingKey.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VerifyingKey")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = Signature{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "sig":
				if validate && zb0004 && "sig" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Signature.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Signature")
					return
				}
				zb0003 = "sig"
			case "idx":
				if validate && zb0004 && "idx" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).VectorCommitmentIndex, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "VectorCommitmentIndex")
					return
				}
				zb0003 = "idx"
			case "prf":
				if validate && zb0004 && "prf" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Proof.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Proof")
					return
				}
				zb0003 = "prf"
			case "vkey":
				if validate && zb0004 && "vkey" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).VerifyingKey.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VerifyingKey")
					return
				}
				zb0003 = "vkey"
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
			zb0004 = true
		}
	}
	o = bts
	return
}

func (z *Signature) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *Signature) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *Signature) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Signature)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Signature) Msgsize() (s int) {
	s = 1 + 4 + (*z).Signature.Msgsize() + 4 + msgp.Uint64Size + 4 + (*z).Proof.Msgsize() + 5 + (*z).VerifyingKey.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Signature) MsgIsZero() bool {
	return ((*z).Signature.MsgIsZero()) && ((*z).VectorCommitmentIndex == 0) && ((*z).Proof.MsgIsZero()) && ((*z).VerifyingKey.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func SignatureMaxSize() (s int) {
	s = 1 + 4 + crypto.FalconSignatureMaxSize() + 4 + msgp.Uint64Size + 4 + merklearray.SingleLeafProofMaxSize() + 5 + crypto.FalconVerifierMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SignerContext) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 4 bits */
	if (*z).FirstValid == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).KeyLifetime == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Tree.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "fv"
			o = append(o, 0xa2, 0x66, 0x76)
			o = msgp.AppendUint64(o, (*z).FirstValid)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "iv"
			o = append(o, 0xa2, 0x69, 0x76)
			o = msgp.AppendUint64(o, (*z).KeyLifetime)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "tree"
			o = append(o, 0xa4, 0x74, 0x72, 0x65, 0x65)
			o = (*z).Tree.MarshalMsg(o)
		}
	}
	return
}

func (_ *SignerContext) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*SignerContext)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SignerContext) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 int
	var zb0003 string
	var zb0004 bool
	var zb0002 bool
	_ = zb0003
	_ = zb0004
	zb0001, zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0001, zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if validate {
			err = &msgp.ErrNonCanonical{}
			return
		}
		if zb0001 > 0 {
			zb0001--
			(*z).FirstValid, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "FirstValid")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).KeyLifetime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "KeyLifetime")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Tree.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Tree")
				return
			}
		}
		if zb0001 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0001)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0002 {
			(*z) = SignerContext{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "fv":
				if validate && zb0004 && "fv" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).FirstValid, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "FirstValid")
					return
				}
				zb0003 = "fv"
			case "iv":
				if validate && zb0004 && "iv" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).KeyLifetime, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "KeyLifetime")
					return
				}
				zb0003 = "iv"
			case "tree":
				if validate && zb0004 && "tree" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Tree.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Tree")
					return
				}
				zb0003 = "tree"
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
			zb0004 = true
		}
	}
	o = bts
	return
}

func (z *SignerContext) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *SignerContext) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *SignerContext) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*SignerContext)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SignerContext) Msgsize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 3 + msgp.Uint64Size + 5 + (*z).Tree.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *SignerContext) MsgIsZero() bool {
	return ((*z).FirstValid == 0) && ((*z).KeyLifetime == 0) && ((*z).Tree.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func SignerContextMaxSize() (s int) {
	s = 1 + 3 + msgp.Uint64Size + 3 + msgp.Uint64Size + 5 + merklearray.TreeMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Verifier) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(2)
	var zb0002Mask uint8 /* 3 bits */
	if (*z).Commitment == (Commitment{}) {
		zb0002Len--
		zb0002Mask |= 0x2
	}
	if (*z).KeyLifetime == 0 {
		zb0002Len--
		zb0002Mask |= 0x4
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x2) == 0 { // if not empty
			// string "cmt"
			o = append(o, 0xa3, 0x63, 0x6d, 0x74)
			o = msgp.AppendBytes(o, ((*z).Commitment)[:])
		}
		if (zb0002Mask & 0x4) == 0 { // if not empty
			// string "lf"
			o = append(o, 0xa2, 0x6c, 0x66)
			o = msgp.AppendUint64(o, (*z).KeyLifetime)
		}
	}
	return
}

func (_ *Verifier) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*Verifier)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Verifier) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0002 int
	var zb0004 string
	var zb0005 bool
	var zb0003 bool
	_ = zb0004
	_ = zb0005
	zb0002, zb0003, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0002, zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if validate {
			err = &msgp.ErrNonCanonical{}
			return
		}
		if zb0002 > 0 {
			zb0002--
			bts, err = msgp.ReadExactBytes(bts, ((*z).Commitment)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Commitment")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			(*z).KeyLifetime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "KeyLifetime")
				return
			}
		}
		if zb0002 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0002)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array")
				return
			}
		}
	} else {
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if zb0003 {
			(*z) = Verifier{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "cmt":
				if validate && zb0005 && "cmt" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = msgp.ReadExactBytes(bts, ((*z).Commitment)[:])
				if err != nil {
					err = msgp.WrapError(err, "Commitment")
					return
				}
				zb0004 = "cmt"
			case "lf":
				if validate && zb0005 && "lf" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).KeyLifetime, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "KeyLifetime")
					return
				}
				zb0004 = "lf"
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
			zb0005 = true
		}
	}
	o = bts
	return
}

func (z *Verifier) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *Verifier) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *Verifier) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*Verifier)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Verifier) Msgsize() (s int) {
	s = 1 + 4 + msgp.ArrayHeaderSize + (MerkleSignatureSchemeRootSize * (msgp.ByteSize)) + 3 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *Verifier) MsgIsZero() bool {
	return ((*z).Commitment == (Commitment{})) && ((*z).KeyLifetime == 0)
}

// MaxSize returns a maximum valid message size for this message type
func VerifierMaxSize() (s int) {
	s = 1 + 4
	// Calculating size of array: z.Commitment
	s += msgp.ArrayHeaderSize + ((MerkleSignatureSchemeRootSize) * (msgp.ByteSize))
	s += 3 + msgp.Uint64Size
	return
}
