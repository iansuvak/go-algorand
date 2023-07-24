package network

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"

	"github.com/algorand/go-algorand/crypto"
)

// The following msgp objects are implemented in this file:
// disconnectReason
//         |-----> MarshalMsg
//         |-----> CanMarshalMsg
//         |-----> (*) UnmarshalMsg
//         |-----> (*) UnmarshalValidateMsg
//         |-----> (*) CanUnmarshalMsg
//         |-----> Msgsize
//         |-----> MsgIsZero
//         |-----> DisconnectReasonMaxSize()
//
// identityChallenge
//         |-----> (*) MarshalMsg
//         |-----> (*) CanMarshalMsg
//         |-----> (*) UnmarshalMsg
//         |-----> (*) UnmarshalValidateMsg
//         |-----> (*) CanUnmarshalMsg
//         |-----> (*) Msgsize
//         |-----> (*) MsgIsZero
//         |-----> IdentityChallengeMaxSize()
//
// identityChallengeResponse
//             |-----> (*) MarshalMsg
//             |-----> (*) CanMarshalMsg
//             |-----> (*) UnmarshalMsg
//             |-----> (*) UnmarshalValidateMsg
//             |-----> (*) CanUnmarshalMsg
//             |-----> (*) Msgsize
//             |-----> (*) MsgIsZero
//             |-----> IdentityChallengeResponseMaxSize()
//
// identityChallengeResponseSigned
//                |-----> (*) MarshalMsg
//                |-----> (*) CanMarshalMsg
//                |-----> (*) UnmarshalMsg
//                |-----> (*) UnmarshalValidateMsg
//                |-----> (*) CanUnmarshalMsg
//                |-----> (*) Msgsize
//                |-----> (*) MsgIsZero
//                |-----> IdentityChallengeResponseSignedMaxSize()
//
// identityChallengeSigned
//            |-----> (*) MarshalMsg
//            |-----> (*) CanMarshalMsg
//            |-----> (*) UnmarshalMsg
//            |-----> (*) UnmarshalValidateMsg
//            |-----> (*) CanUnmarshalMsg
//            |-----> (*) Msgsize
//            |-----> (*) MsgIsZero
//            |-----> IdentityChallengeSignedMaxSize()
//
// identityChallengeValue
//            |-----> (*) MarshalMsg
//            |-----> (*) CanMarshalMsg
//            |-----> (*) UnmarshalMsg
//            |-----> (*) UnmarshalValidateMsg
//            |-----> (*) CanUnmarshalMsg
//            |-----> (*) Msgsize
//            |-----> (*) MsgIsZero
//            |-----> IdentityChallengeValueMaxSize()
//
// identityVerificationMessage
//              |-----> (*) MarshalMsg
//              |-----> (*) CanMarshalMsg
//              |-----> (*) UnmarshalMsg
//              |-----> (*) UnmarshalValidateMsg
//              |-----> (*) CanUnmarshalMsg
//              |-----> (*) Msgsize
//              |-----> (*) MsgIsZero
//              |-----> IdentityVerificationMessageMaxSize()
//
// identityVerificationMessageSigned
//                 |-----> (*) MarshalMsg
//                 |-----> (*) CanMarshalMsg
//                 |-----> (*) UnmarshalMsg
//                 |-----> (*) UnmarshalValidateMsg
//                 |-----> (*) CanUnmarshalMsg
//                 |-----> (*) Msgsize
//                 |-----> (*) MsgIsZero
//                 |-----> IdentityVerificationMessageSignedMaxSize()
//

// MarshalMsg implements msgp.Marshaler
func (z disconnectReason) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

func (_ disconnectReason) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(disconnectReason)
	if !ok {
		_, ok = (z).(*disconnectReason)
	}
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *disconnectReason) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = disconnectReason(zb0001)
	}
	o = bts
	return
}

func (z *disconnectReason) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *disconnectReason) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *disconnectReason) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*disconnectReason)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z disconnectReason) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// MsgIsZero returns whether this is a zero value
func (z disconnectReason) MsgIsZero() bool {
	return z == ""
}

// MaxSize returns a maximum valid message size for this message type
func DisconnectReasonMaxSize() (s int) {
	panic("Unable to determine max size: String type string(z) is unbounded")
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityChallenge) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(3)
	var zb0002Mask uint8 /* 4 bits */
	if len((*z).PublicAddress) == 0 {
		zb0002Len--
		zb0002Mask |= 0x2
	}
	if (*z).Challenge == (identityChallengeValue{}) {
		zb0002Len--
		zb0002Mask |= 0x4
	}
	if (*z).Key.MsgIsZero() {
		zb0002Len--
		zb0002Mask |= 0x8
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x2) == 0 { // if not empty
			// string "a"
			o = append(o, 0xa1, 0x61)
			o = msgp.AppendBytes(o, (*z).PublicAddress)
		}
		if (zb0002Mask & 0x4) == 0 { // if not empty
			// string "c"
			o = append(o, 0xa1, 0x63)
			o = msgp.AppendBytes(o, ((*z).Challenge)[:])
		}
		if (zb0002Mask & 0x8) == 0 { // if not empty
			// string "pk"
			o = append(o, 0xa2, 0x70, 0x6b)
			o = (*z).Key.MarshalMsg(o)
		}
	}
	return
}

func (_ *identityChallenge) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallenge)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityChallenge) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			bts, err = (*z).Key.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Key")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			bts, err = msgp.ReadExactBytes(bts, ((*z).Challenge)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Challenge")
				return
			}
		}
		if zb0002 > 0 {
			zb0002--
			var zb0006 int
			zb0006, err = msgp.ReadBytesBytesHeader(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "PublicAddress")
				return
			}
			if zb0006 > maxAddressLen {
				err = msgp.ErrOverflow(uint64(zb0006), uint64(maxAddressLen))
				return
			}
			(*z).PublicAddress, bts, err = msgp.ReadBytesBytes(bts, (*z).PublicAddress)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "PublicAddress")
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
			(*z) = identityChallenge{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "pk":
				if validate && zb0005 && "pk" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Key.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Key")
					return
				}
				zb0004 = "pk"
			case "c":
				if validate && zb0005 && "c" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = msgp.ReadExactBytes(bts, ((*z).Challenge)[:])
				if err != nil {
					err = msgp.WrapError(err, "Challenge")
					return
				}
				zb0004 = "c"
			case "a":
				if validate && zb0005 && "a" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				var zb0007 int
				zb0007, err = msgp.ReadBytesBytesHeader(bts)
				if err != nil {
					err = msgp.WrapError(err, "PublicAddress")
					return
				}
				if zb0007 > maxAddressLen {
					err = msgp.ErrOverflow(uint64(zb0007), uint64(maxAddressLen))
					return
				}
				(*z).PublicAddress, bts, err = msgp.ReadBytesBytes(bts, (*z).PublicAddress)
				if err != nil {
					err = msgp.WrapError(err, "PublicAddress")
					return
				}
				zb0004 = "a"
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

func (z *identityChallenge) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityChallenge) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityChallenge) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallenge)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityChallenge) Msgsize() (s int) {
	s = 1 + 3 + (*z).Key.Msgsize() + 2 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 2 + msgp.BytesPrefixSize + len((*z).PublicAddress)
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityChallenge) MsgIsZero() bool {
	return ((*z).Key.MsgIsZero()) && ((*z).Challenge == (identityChallengeValue{})) && (len((*z).PublicAddress) == 0)
}

// MaxSize returns a maximum valid message size for this message type
func IdentityChallengeMaxSize() (s int) {
	s = 1 + 3 + crypto.PublicKeyMaxSize() + 2
	// Calculating size of array: z.Challenge
	s += msgp.ArrayHeaderSize + ((32) * (msgp.ByteSize))
	s += 2 + msgp.BytesPrefixSize + maxAddressLen
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityChallengeResponse) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0003Len := uint32(3)
	var zb0003Mask uint8 /* 4 bits */
	if (*z).Challenge == (identityChallengeValue{}) {
		zb0003Len--
		zb0003Mask |= 0x2
	}
	if (*z).Key.MsgIsZero() {
		zb0003Len--
		zb0003Mask |= 0x4
	}
	if (*z).ResponseChallenge == (identityChallengeValue{}) {
		zb0003Len--
		zb0003Mask |= 0x8
	}
	// variable map header, size zb0003Len
	o = append(o, 0x80|uint8(zb0003Len))
	if zb0003Len != 0 {
		if (zb0003Mask & 0x2) == 0 { // if not empty
			// string "c"
			o = append(o, 0xa1, 0x63)
			o = msgp.AppendBytes(o, ((*z).Challenge)[:])
		}
		if (zb0003Mask & 0x4) == 0 { // if not empty
			// string "pk"
			o = append(o, 0xa2, 0x70, 0x6b)
			o = (*z).Key.MarshalMsg(o)
		}
		if (zb0003Mask & 0x8) == 0 { // if not empty
			// string "rc"
			o = append(o, 0xa2, 0x72, 0x63)
			o = msgp.AppendBytes(o, ((*z).ResponseChallenge)[:])
		}
	}
	return
}

func (_ *identityChallengeResponse) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeResponse)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityChallengeResponse) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0003 int
	var zb0005 string
	var zb0006 bool
	var zb0004 bool
	_ = zb0005
	_ = zb0006
	zb0003, zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
	if _, ok := err.(msgp.TypeError); ok {
		zb0003, zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		if validate {
			err = &msgp.ErrNonCanonical{}
			return
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = (*z).Key.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Key")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = msgp.ReadExactBytes(bts, ((*z).Challenge)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Challenge")
				return
			}
		}
		if zb0003 > 0 {
			zb0003--
			bts, err = msgp.ReadExactBytes(bts, ((*z).ResponseChallenge)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "ResponseChallenge")
				return
			}
		}
		if zb0003 > 0 {
			err = msgp.ErrTooManyArrayFields(zb0003)
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
		if zb0004 {
			(*z) = identityChallengeResponse{}
		}
		for zb0003 > 0 {
			zb0003--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "pk":
				if validate && zb0006 && "pk" < zb0005 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Key.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Key")
					return
				}
				zb0005 = "pk"
			case "c":
				if validate && zb0006 && "c" < zb0005 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = msgp.ReadExactBytes(bts, ((*z).Challenge)[:])
				if err != nil {
					err = msgp.WrapError(err, "Challenge")
					return
				}
				zb0005 = "c"
			case "rc":
				if validate && zb0006 && "rc" < zb0005 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = msgp.ReadExactBytes(bts, ((*z).ResponseChallenge)[:])
				if err != nil {
					err = msgp.WrapError(err, "ResponseChallenge")
					return
				}
				zb0005 = "rc"
			default:
				err = msgp.ErrNoField(string(field))
				if err != nil {
					err = msgp.WrapError(err)
					return
				}
			}
			zb0006 = true
		}
	}
	o = bts
	return
}

func (z *identityChallengeResponse) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityChallengeResponse) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityChallengeResponse) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeResponse)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityChallengeResponse) Msgsize() (s int) {
	s = 1 + 3 + (*z).Key.Msgsize() + 2 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 3 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityChallengeResponse) MsgIsZero() bool {
	return ((*z).Key.MsgIsZero()) && ((*z).Challenge == (identityChallengeValue{})) && ((*z).ResponseChallenge == (identityChallengeValue{}))
}

// MaxSize returns a maximum valid message size for this message type
func IdentityChallengeResponseMaxSize() (s int) {
	s = 1 + 3 + crypto.PublicKeyMaxSize() + 2
	// Calculating size of array: z.Challenge
	s += msgp.ArrayHeaderSize + ((32) * (msgp.ByteSize))
	s += 3
	// Calculating size of array: z.ResponseChallenge
	s += msgp.ArrayHeaderSize + ((32) * (msgp.ByteSize))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityChallengeResponseSigned) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Msg.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Signature.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "icr"
			o = append(o, 0xa3, 0x69, 0x63, 0x72)
			o = (*z).Msg.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "sig"
			o = append(o, 0xa3, 0x73, 0x69, 0x67)
			o = (*z).Signature.MarshalMsg(o)
		}
	}
	return
}

func (_ *identityChallengeResponseSigned) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeResponseSigned)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityChallengeResponseSigned) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			bts, err = (*z).Msg.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Msg")
				return
			}
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
			(*z) = identityChallengeResponseSigned{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "icr":
				if validate && zb0004 && "icr" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Msg.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Msg")
					return
				}
				zb0003 = "icr"
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

func (z *identityChallengeResponseSigned) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityChallengeResponseSigned) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityChallengeResponseSigned) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeResponseSigned)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityChallengeResponseSigned) Msgsize() (s int) {
	s = 1 + 4 + (*z).Msg.Msgsize() + 4 + (*z).Signature.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityChallengeResponseSigned) MsgIsZero() bool {
	return ((*z).Msg.MsgIsZero()) && ((*z).Signature.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func IdentityChallengeResponseSignedMaxSize() (s int) {
	s = 1 + 4 + IdentityChallengeResponseMaxSize() + 4 + crypto.SignatureMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityChallengeSigned) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Msg.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).Signature.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "ic"
			o = append(o, 0xa2, 0x69, 0x63)
			o = (*z).Msg.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "sig"
			o = append(o, 0xa3, 0x73, 0x69, 0x67)
			o = (*z).Signature.MarshalMsg(o)
		}
	}
	return
}

func (_ *identityChallengeSigned) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeSigned)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityChallengeSigned) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			bts, err = (*z).Msg.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Msg")
				return
			}
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
			(*z) = identityChallengeSigned{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "ic":
				if validate && zb0004 && "ic" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Msg.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Msg")
					return
				}
				zb0003 = "ic"
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

func (z *identityChallengeSigned) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityChallengeSigned) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityChallengeSigned) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeSigned)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityChallengeSigned) Msgsize() (s int) {
	s = 1 + 3 + (*z).Msg.Msgsize() + 4 + (*z).Signature.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityChallengeSigned) MsgIsZero() bool {
	return ((*z).Msg.MsgIsZero()) && ((*z).Signature.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func IdentityChallengeSignedMaxSize() (s int) {
	s = 1 + 3 + IdentityChallengeMaxSize() + 4 + crypto.SignatureMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityChallengeValue) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, (*z)[:])
	return
}

func (_ *identityChallengeValue) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeValue)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityChallengeValue) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
	bts, err = msgp.ReadExactBytes(bts, (*z)[:])
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	o = bts
	return
}

func (z *identityChallengeValue) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityChallengeValue) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityChallengeValue) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityChallengeValue)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityChallengeValue) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (32 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityChallengeValue) MsgIsZero() bool {
	return (*z) == (identityChallengeValue{})
}

// MaxSize returns a maximum valid message size for this message type
func IdentityChallengeValueMaxSize() (s int) {
	// Calculating size of array: z
	s = msgp.ArrayHeaderSize + ((32) * (msgp.ByteSize))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityVerificationMessage) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(1)
	var zb0002Mask uint8 /* 2 bits */
	if (*z).ResponseChallenge == (identityChallengeValue{}) {
		zb0002Len--
		zb0002Mask |= 0x2
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x2) == 0 { // if not empty
			// string "rc"
			o = append(o, 0xa2, 0x72, 0x63)
			o = msgp.AppendBytes(o, ((*z).ResponseChallenge)[:])
		}
	}
	return
}

func (_ *identityVerificationMessage) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityVerificationMessage)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityVerificationMessage) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			bts, err = msgp.ReadExactBytes(bts, ((*z).ResponseChallenge)[:])
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "ResponseChallenge")
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
			(*z) = identityVerificationMessage{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "rc":
				if validate && zb0005 && "rc" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = msgp.ReadExactBytes(bts, ((*z).ResponseChallenge)[:])
				if err != nil {
					err = msgp.WrapError(err, "ResponseChallenge")
					return
				}
				zb0004 = "rc"
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

func (z *identityVerificationMessage) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityVerificationMessage) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityVerificationMessage) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityVerificationMessage)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityVerificationMessage) Msgsize() (s int) {
	s = 1 + 3 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize))
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityVerificationMessage) MsgIsZero() bool {
	return ((*z).ResponseChallenge == (identityChallengeValue{}))
}

// MaxSize returns a maximum valid message size for this message type
func IdentityVerificationMessageMaxSize() (s int) {
	s = 1 + 3
	// Calculating size of array: z.ResponseChallenge
	s += msgp.ArrayHeaderSize + ((32) * (msgp.ByteSize))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *identityVerificationMessageSigned) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0002Len := uint32(2)
	var zb0002Mask uint8 /* 3 bits */
	if (*z).Msg.ResponseChallenge == (identityChallengeValue{}) {
		zb0002Len--
		zb0002Mask |= 0x2
	}
	if (*z).Signature.MsgIsZero() {
		zb0002Len--
		zb0002Mask |= 0x4
	}
	// variable map header, size zb0002Len
	o = append(o, 0x80|uint8(zb0002Len))
	if zb0002Len != 0 {
		if (zb0002Mask & 0x2) == 0 { // if not empty
			// string "ivm"
			o = append(o, 0xa3, 0x69, 0x76, 0x6d)
			// omitempty: check for empty values
			zb0003Len := uint32(1)
			var zb0003Mask uint8 /* 2 bits */
			if (*z).Msg.ResponseChallenge == (identityChallengeValue{}) {
				zb0003Len--
				zb0003Mask |= 0x2
			}
			// variable map header, size zb0003Len
			o = append(o, 0x80|uint8(zb0003Len))
			if (zb0003Mask & 0x2) == 0 { // if not empty
				// string "rc"
				o = append(o, 0xa2, 0x72, 0x63)
				o = msgp.AppendBytes(o, ((*z).Msg.ResponseChallenge)[:])
			}
		}
		if (zb0002Mask & 0x4) == 0 { // if not empty
			// string "sig"
			o = append(o, 0xa3, 0x73, 0x69, 0x67)
			o = (*z).Signature.MarshalMsg(o)
		}
	}
	return
}

func (_ *identityVerificationMessageSigned) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityVerificationMessageSigned)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *identityVerificationMessageSigned) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			var zb0006 int
			var zb0008 string
			var zb0009 bool
			var zb0007 bool
			_ = zb0008
			_ = zb0009
			zb0006, zb0007, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0006, zb0007, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Msg")
					return
				}
				if validate {
					err = &msgp.ErrNonCanonical{}
					return
				}
				if zb0006 > 0 {
					zb0006--
					bts, err = msgp.ReadExactBytes(bts, ((*z).Msg.ResponseChallenge)[:])
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Msg", "struct-from-array", "ResponseChallenge")
						return
					}
				}
				if zb0006 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0006)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Msg", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Msg")
					return
				}
				if zb0007 {
					(*z).Msg = identityVerificationMessage{}
				}
				for zb0006 > 0 {
					zb0006--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Msg")
						return
					}
					switch string(field) {
					case "rc":
						if validate && zb0009 && "rc" < zb0008 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						bts, err = msgp.ReadExactBytes(bts, ((*z).Msg.ResponseChallenge)[:])
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Msg", "ResponseChallenge")
							return
						}
						zb0008 = "rc"
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Msg")
							return
						}
					}
					zb0009 = true
				}
			}
		}
		if zb0002 > 0 {
			zb0002--
			bts, err = (*z).Signature.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Signature")
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
			(*z) = identityVerificationMessageSigned{}
		}
		for zb0002 > 0 {
			zb0002--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "ivm":
				if validate && zb0005 && "ivm" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				var zb0010 int
				var zb0012 string
				var zb0013 bool
				var zb0011 bool
				_ = zb0012
				_ = zb0013
				zb0010, zb0011, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0010, zb0011, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Msg")
						return
					}
					if validate {
						err = &msgp.ErrNonCanonical{}
						return
					}
					if zb0010 > 0 {
						zb0010--
						bts, err = msgp.ReadExactBytes(bts, ((*z).Msg.ResponseChallenge)[:])
						if err != nil {
							err = msgp.WrapError(err, "Msg", "struct-from-array", "ResponseChallenge")
							return
						}
					}
					if zb0010 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0010)
						if err != nil {
							err = msgp.WrapError(err, "Msg", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "Msg")
						return
					}
					if zb0011 {
						(*z).Msg = identityVerificationMessage{}
					}
					for zb0010 > 0 {
						zb0010--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "Msg")
							return
						}
						switch string(field) {
						case "rc":
							if validate && zb0013 && "rc" < zb0012 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							bts, err = msgp.ReadExactBytes(bts, ((*z).Msg.ResponseChallenge)[:])
							if err != nil {
								err = msgp.WrapError(err, "Msg", "ResponseChallenge")
								return
							}
							zb0012 = "rc"
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "Msg")
								return
							}
						}
						zb0013 = true
					}
				}
				zb0004 = "ivm"
			case "sig":
				if validate && zb0005 && "sig" < zb0004 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Signature.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Signature")
					return
				}
				zb0004 = "sig"
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

func (z *identityVerificationMessageSigned) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *identityVerificationMessageSigned) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *identityVerificationMessageSigned) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*identityVerificationMessageSigned)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *identityVerificationMessageSigned) Msgsize() (s int) {
	s = 1 + 4 + 1 + 3 + msgp.ArrayHeaderSize + (32 * (msgp.ByteSize)) + 4 + (*z).Signature.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *identityVerificationMessageSigned) MsgIsZero() bool {
	return ((*z).Msg.ResponseChallenge == (identityChallengeValue{})) && ((*z).Signature.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func IdentityVerificationMessageSignedMaxSize() (s int) {
	s = 1 + 4 + 1 + 3
	// Calculating size of array: z.Msg.ResponseChallenge
	s += msgp.ArrayHeaderSize + ((32) * (msgp.ByteSize))
	s += 4 + crypto.SignatureMaxSize()
	return
}
