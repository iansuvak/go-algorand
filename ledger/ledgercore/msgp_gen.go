package ledgercore

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"github.com/algorand/msgp/msgp"

	"github.com/algorand/go-algorand/crypto"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/protocol"
)

// The following msgp objects are implemented in this file:
// AccountTotals
//       |-----> (*) MarshalMsg
//       |-----> (*) CanMarshalMsg
//       |-----> (*) UnmarshalMsg
//       |-----> (*) UnmarshalValidateMsg
//       |-----> (*) CanUnmarshalMsg
//       |-----> (*) Msgsize
//       |-----> (*) MsgIsZero
//       |-----> AccountTotalsMaxSize()
//
// AlgoCount
//     |-----> (*) MarshalMsg
//     |-----> (*) CanMarshalMsg
//     |-----> (*) UnmarshalMsg
//     |-----> (*) UnmarshalValidateMsg
//     |-----> (*) CanUnmarshalMsg
//     |-----> (*) Msgsize
//     |-----> (*) MsgIsZero
//     |-----> AlgoCountMaxSize()
//
// OnlineRoundParamsData
//           |-----> (*) MarshalMsg
//           |-----> (*) CanMarshalMsg
//           |-----> (*) UnmarshalMsg
//           |-----> (*) UnmarshalValidateMsg
//           |-----> (*) CanUnmarshalMsg
//           |-----> (*) Msgsize
//           |-----> (*) MsgIsZero
//           |-----> OnlineRoundParamsDataMaxSize()
//
// StateProofVerificationContext
//               |-----> (*) MarshalMsg
//               |-----> (*) CanMarshalMsg
//               |-----> (*) UnmarshalMsg
//               |-----> (*) UnmarshalValidateMsg
//               |-----> (*) CanUnmarshalMsg
//               |-----> (*) Msgsize
//               |-----> (*) MsgIsZero
//               |-----> StateProofVerificationContextMaxSize()
//

// MarshalMsg implements msgp.Marshaler
func (z *AccountTotals) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(4)
	var zb0001Mask uint8 /* 5 bits */
	if ((*z).NotParticipating.Money.MsgIsZero()) && ((*z).NotParticipating.RewardUnits == 0) {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if ((*z).Offline.Money.MsgIsZero()) && ((*z).Offline.RewardUnits == 0) {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if ((*z).Online.Money.MsgIsZero()) && ((*z).Online.RewardUnits == 0) {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).RewardsLevel == 0 {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "notpart"
			o = append(o, 0xa7, 0x6e, 0x6f, 0x74, 0x70, 0x61, 0x72, 0x74)
			// omitempty: check for empty values
			zb0002Len := uint32(2)
			var zb0002Mask uint8 /* 3 bits */
			if (*z).NotParticipating.Money.MsgIsZero() {
				zb0002Len--
				zb0002Mask |= 0x2
			}
			if (*z).NotParticipating.RewardUnits == 0 {
				zb0002Len--
				zb0002Mask |= 0x4
			}
			// variable map header, size zb0002Len
			o = append(o, 0x80|uint8(zb0002Len))
			if (zb0002Mask & 0x2) == 0 { // if not empty
				// string "mon"
				o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
				o = (*z).NotParticipating.Money.MarshalMsg(o)
			}
			if (zb0002Mask & 0x4) == 0 { // if not empty
				// string "rwd"
				o = append(o, 0xa3, 0x72, 0x77, 0x64)
				o = msgp.AppendUint64(o, (*z).NotParticipating.RewardUnits)
			}
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "offline"
			o = append(o, 0xa7, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65)
			// omitempty: check for empty values
			zb0003Len := uint32(2)
			var zb0003Mask uint8 /* 3 bits */
			if (*z).Offline.Money.MsgIsZero() {
				zb0003Len--
				zb0003Mask |= 0x2
			}
			if (*z).Offline.RewardUnits == 0 {
				zb0003Len--
				zb0003Mask |= 0x4
			}
			// variable map header, size zb0003Len
			o = append(o, 0x80|uint8(zb0003Len))
			if (zb0003Mask & 0x2) == 0 { // if not empty
				// string "mon"
				o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
				o = (*z).Offline.Money.MarshalMsg(o)
			}
			if (zb0003Mask & 0x4) == 0 { // if not empty
				// string "rwd"
				o = append(o, 0xa3, 0x72, 0x77, 0x64)
				o = msgp.AppendUint64(o, (*z).Offline.RewardUnits)
			}
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "online"
			o = append(o, 0xa6, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65)
			// omitempty: check for empty values
			zb0004Len := uint32(2)
			var zb0004Mask uint8 /* 3 bits */
			if (*z).Online.Money.MsgIsZero() {
				zb0004Len--
				zb0004Mask |= 0x2
			}
			if (*z).Online.RewardUnits == 0 {
				zb0004Len--
				zb0004Mask |= 0x4
			}
			// variable map header, size zb0004Len
			o = append(o, 0x80|uint8(zb0004Len))
			if (zb0004Mask & 0x2) == 0 { // if not empty
				// string "mon"
				o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
				o = (*z).Online.Money.MarshalMsg(o)
			}
			if (zb0004Mask & 0x4) == 0 { // if not empty
				// string "rwd"
				o = append(o, 0xa3, 0x72, 0x77, 0x64)
				o = msgp.AppendUint64(o, (*z).Online.RewardUnits)
			}
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "rwdlvl"
			o = append(o, 0xa6, 0x72, 0x77, 0x64, 0x6c, 0x76, 0x6c)
			o = msgp.AppendUint64(o, (*z).RewardsLevel)
		}
	}
	return
}

func (_ *AccountTotals) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountTotals)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AccountTotals) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			var zb0005 int
			var zb0007 string
			var zb0008 bool
			var zb0006 bool
			_ = zb0007
			_ = zb0008
			zb0005, zb0006, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0005, zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Online")
					return
				}
				if validate {
					err = &msgp.ErrNonCanonical{}
					return
				}
				if zb0005 > 0 {
					zb0005--
					bts, err = (*z).Online.Money.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online", "struct-from-array", "Money")
						return
					}
				}
				if zb0005 > 0 {
					zb0005--
					(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online", "struct-from-array", "RewardUnits")
						return
					}
				}
				if zb0005 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0005)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Online")
					return
				}
				if zb0006 {
					(*z).Online = AlgoCount{}
				}
				for zb0005 > 0 {
					zb0005--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Online")
						return
					}
					switch string(field) {
					case "mon":
						if validate && zb0008 && "mon" < zb0007 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						bts, err = (*z).Online.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Online", "Money")
							return
						}
						zb0007 = "mon"
					case "rwd":
						if validate && zb0008 && "rwd" < zb0007 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Online", "RewardUnits")
							return
						}
						zb0007 = "rwd"
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Online")
							return
						}
					}
					zb0008 = true
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			var zb0009 int
			var zb0011 string
			var zb0012 bool
			var zb0010 bool
			_ = zb0011
			_ = zb0012
			zb0009, zb0010, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0009, zb0010, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Offline")
					return
				}
				if validate {
					err = &msgp.ErrNonCanonical{}
					return
				}
				if zb0009 > 0 {
					zb0009--
					bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline", "struct-from-array", "Money")
						return
					}
				}
				if zb0009 > 0 {
					zb0009--
					(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline", "struct-from-array", "RewardUnits")
						return
					}
				}
				if zb0009 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0009)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "Offline")
					return
				}
				if zb0010 {
					(*z).Offline = AlgoCount{}
				}
				for zb0009 > 0 {
					zb0009--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "Offline")
						return
					}
					switch string(field) {
					case "mon":
						if validate && zb0012 && "mon" < zb0011 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Offline", "Money")
							return
						}
						zb0011 = "mon"
					case "rwd":
						if validate && zb0012 && "rwd" < zb0011 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Offline", "RewardUnits")
							return
						}
						zb0011 = "rwd"
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "Offline")
							return
						}
					}
					zb0012 = true
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			var zb0013 int
			var zb0015 string
			var zb0016 bool
			var zb0014 bool
			_ = zb0015
			_ = zb0016
			zb0013, zb0014, bts, err = msgp.ReadMapHeaderBytes(bts)
			if _, ok := err.(msgp.TypeError); ok {
				zb0013, zb0014, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
					return
				}
				if validate {
					err = &msgp.ErrNonCanonical{}
					return
				}
				if zb0013 > 0 {
					zb0013--
					bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "struct-from-array", "Money")
						return
					}
				}
				if zb0013 > 0 {
					zb0013--
					(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "struct-from-array", "RewardUnits")
						return
					}
				}
				if zb0013 > 0 {
					err = msgp.ErrTooManyArrayFields(zb0013)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "struct-from-array")
						return
					}
				}
			} else {
				if err != nil {
					err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
					return
				}
				if zb0014 {
					(*z).NotParticipating = AlgoCount{}
				}
				for zb0013 > 0 {
					zb0013--
					field, bts, err = msgp.ReadMapKeyZC(bts)
					if err != nil {
						err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
						return
					}
					switch string(field) {
					case "mon":
						if validate && zb0016 && "mon" < zb0015 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "Money")
							return
						}
						zb0015 = "mon"
					case "rwd":
						if validate && zb0016 && "rwd" < zb0015 {
							err = &msgp.ErrNonCanonical{}
							return
						}
						(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "NotParticipating", "RewardUnits")
							return
						}
						zb0015 = "rwd"
					default:
						err = msgp.ErrNoField(string(field))
						if err != nil {
							err = msgp.WrapError(err, "struct-from-array", "NotParticipating")
							return
						}
					}
					zb0016 = true
				}
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "RewardsLevel")
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
			(*z) = AccountTotals{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "online":
				if validate && zb0004 && "online" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				var zb0017 int
				var zb0019 string
				var zb0020 bool
				var zb0018 bool
				_ = zb0019
				_ = zb0020
				zb0017, zb0018, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0017, zb0018, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Online")
						return
					}
					if validate {
						err = &msgp.ErrNonCanonical{}
						return
					}
					if zb0017 > 0 {
						zb0017--
						bts, err = (*z).Online.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "Online", "struct-from-array", "Money")
							return
						}
					}
					if zb0017 > 0 {
						zb0017--
						(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Online", "struct-from-array", "RewardUnits")
							return
						}
					}
					if zb0017 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0017)
						if err != nil {
							err = msgp.WrapError(err, "Online", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "Online")
						return
					}
					if zb0018 {
						(*z).Online = AlgoCount{}
					}
					for zb0017 > 0 {
						zb0017--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "Online")
							return
						}
						switch string(field) {
						case "mon":
							if validate && zb0020 && "mon" < zb0019 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							bts, err = (*z).Online.Money.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "Online", "Money")
								return
							}
							zb0019 = "mon"
						case "rwd":
							if validate && zb0020 && "rwd" < zb0019 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							(*z).Online.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Online", "RewardUnits")
								return
							}
							zb0019 = "rwd"
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "Online")
								return
							}
						}
						zb0020 = true
					}
				}
				zb0003 = "online"
			case "offline":
				if validate && zb0004 && "offline" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				var zb0021 int
				var zb0023 string
				var zb0024 bool
				var zb0022 bool
				_ = zb0023
				_ = zb0024
				zb0021, zb0022, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0021, zb0022, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "Offline")
						return
					}
					if validate {
						err = &msgp.ErrNonCanonical{}
						return
					}
					if zb0021 > 0 {
						zb0021--
						bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "Offline", "struct-from-array", "Money")
							return
						}
					}
					if zb0021 > 0 {
						zb0021--
						(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "Offline", "struct-from-array", "RewardUnits")
							return
						}
					}
					if zb0021 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0021)
						if err != nil {
							err = msgp.WrapError(err, "Offline", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "Offline")
						return
					}
					if zb0022 {
						(*z).Offline = AlgoCount{}
					}
					for zb0021 > 0 {
						zb0021--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "Offline")
							return
						}
						switch string(field) {
						case "mon":
							if validate && zb0024 && "mon" < zb0023 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							bts, err = (*z).Offline.Money.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "Offline", "Money")
								return
							}
							zb0023 = "mon"
						case "rwd":
							if validate && zb0024 && "rwd" < zb0023 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							(*z).Offline.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "Offline", "RewardUnits")
								return
							}
							zb0023 = "rwd"
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "Offline")
								return
							}
						}
						zb0024 = true
					}
				}
				zb0003 = "offline"
			case "notpart":
				if validate && zb0004 && "notpart" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				var zb0025 int
				var zb0027 string
				var zb0028 bool
				var zb0026 bool
				_ = zb0027
				_ = zb0028
				zb0025, zb0026, bts, err = msgp.ReadMapHeaderBytes(bts)
				if _, ok := err.(msgp.TypeError); ok {
					zb0025, zb0026, bts, err = msgp.ReadArrayHeaderBytes(bts)
					if err != nil {
						err = msgp.WrapError(err, "NotParticipating")
						return
					}
					if validate {
						err = &msgp.ErrNonCanonical{}
						return
					}
					if zb0025 > 0 {
						zb0025--
						bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating", "struct-from-array", "Money")
							return
						}
					}
					if zb0025 > 0 {
						zb0025--
						(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating", "struct-from-array", "RewardUnits")
							return
						}
					}
					if zb0025 > 0 {
						err = msgp.ErrTooManyArrayFields(zb0025)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating", "struct-from-array")
							return
						}
					}
				} else {
					if err != nil {
						err = msgp.WrapError(err, "NotParticipating")
						return
					}
					if zb0026 {
						(*z).NotParticipating = AlgoCount{}
					}
					for zb0025 > 0 {
						zb0025--
						field, bts, err = msgp.ReadMapKeyZC(bts)
						if err != nil {
							err = msgp.WrapError(err, "NotParticipating")
							return
						}
						switch string(field) {
						case "mon":
							if validate && zb0028 && "mon" < zb0027 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							bts, err = (*z).NotParticipating.Money.UnmarshalMsg(bts)
							if err != nil {
								err = msgp.WrapError(err, "NotParticipating", "Money")
								return
							}
							zb0027 = "mon"
						case "rwd":
							if validate && zb0028 && "rwd" < zb0027 {
								err = &msgp.ErrNonCanonical{}
								return
							}
							(*z).NotParticipating.RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
							if err != nil {
								err = msgp.WrapError(err, "NotParticipating", "RewardUnits")
								return
							}
							zb0027 = "rwd"
						default:
							err = msgp.ErrNoField(string(field))
							if err != nil {
								err = msgp.WrapError(err, "NotParticipating")
								return
							}
						}
						zb0028 = true
					}
				}
				zb0003 = "notpart"
			case "rwdlvl":
				if validate && zb0004 && "rwdlvl" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RewardsLevel")
					return
				}
				zb0003 = "rwdlvl"
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

func (z *AccountTotals) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *AccountTotals) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *AccountTotals) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*AccountTotals)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AccountTotals) Msgsize() (s int) {
	s = 1 + 7 + 1 + 4 + (*z).Online.Money.Msgsize() + 4 + msgp.Uint64Size + 8 + 1 + 4 + (*z).Offline.Money.Msgsize() + 4 + msgp.Uint64Size + 8 + 1 + 4 + (*z).NotParticipating.Money.Msgsize() + 4 + msgp.Uint64Size + 7 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *AccountTotals) MsgIsZero() bool {
	return (((*z).Online.Money.MsgIsZero()) && ((*z).Online.RewardUnits == 0)) && (((*z).Offline.Money.MsgIsZero()) && ((*z).Offline.RewardUnits == 0)) && (((*z).NotParticipating.Money.MsgIsZero()) && ((*z).NotParticipating.RewardUnits == 0)) && ((*z).RewardsLevel == 0)
}

// MaxSize returns a maximum valid message size for this message type
func AccountTotalsMaxSize() (s int) {
	s = 1 + 7 + 1 + 4 + basics.MicroAlgosMaxSize() + 4 + msgp.Uint64Size + 8 + 1 + 4 + basics.MicroAlgosMaxSize() + 4 + msgp.Uint64Size + 8 + 1 + 4 + basics.MicroAlgosMaxSize() + 4 + msgp.Uint64Size + 7 + msgp.Uint64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *AlgoCount) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(2)
	var zb0001Mask uint8 /* 3 bits */
	if (*z).Money.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).RewardUnits == 0 {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "mon"
			o = append(o, 0xa3, 0x6d, 0x6f, 0x6e)
			o = (*z).Money.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "rwd"
			o = append(o, 0xa3, 0x72, 0x77, 0x64)
			o = msgp.AppendUint64(o, (*z).RewardUnits)
		}
	}
	return
}

func (_ *AlgoCount) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*AlgoCount)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AlgoCount) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			bts, err = (*z).Money.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Money")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "RewardUnits")
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
			(*z) = AlgoCount{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "mon":
				if validate && zb0004 && "mon" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Money.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Money")
					return
				}
				zb0003 = "mon"
			case "rwd":
				if validate && zb0004 && "rwd" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).RewardUnits, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RewardUnits")
					return
				}
				zb0003 = "rwd"
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

func (z *AlgoCount) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *AlgoCount) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *AlgoCount) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*AlgoCount)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *AlgoCount) Msgsize() (s int) {
	s = 1 + 4 + (*z).Money.Msgsize() + 4 + msgp.Uint64Size
	return
}

// MsgIsZero returns whether this is a zero value
func (z *AlgoCount) MsgIsZero() bool {
	return ((*z).Money.MsgIsZero()) && ((*z).RewardUnits == 0)
}

// MaxSize returns a maximum valid message size for this message type
func AlgoCountMaxSize() (s int) {
	s = 1 + 4 + basics.MicroAlgosMaxSize() + 4 + msgp.Uint64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *OnlineRoundParamsData) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(3)
	var zb0001Mask uint8 /* 4 bits */
	if (*z).OnlineSupply == 0 {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).CurrentProtocol.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).RewardsLevel == 0 {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "online"
			o = append(o, 0xa6, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65)
			o = msgp.AppendUint64(o, (*z).OnlineSupply)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "proto"
			o = append(o, 0xa5, 0x70, 0x72, 0x6f, 0x74, 0x6f)
			o = (*z).CurrentProtocol.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "rwdlvl"
			o = append(o, 0xa6, 0x72, 0x77, 0x64, 0x6c, 0x76, 0x6c)
			o = msgp.AppendUint64(o, (*z).RewardsLevel)
		}
	}
	return
}

func (_ *OnlineRoundParamsData) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*OnlineRoundParamsData)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *OnlineRoundParamsData) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			(*z).OnlineSupply, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "OnlineSupply")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "RewardsLevel")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).CurrentProtocol.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "CurrentProtocol")
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
			(*z) = OnlineRoundParamsData{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "online":
				if validate && zb0004 && "online" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).OnlineSupply, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "OnlineSupply")
					return
				}
				zb0003 = "online"
			case "rwdlvl":
				if validate && zb0004 && "rwdlvl" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				(*z).RewardsLevel, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "RewardsLevel")
					return
				}
				zb0003 = "rwdlvl"
			case "proto":
				if validate && zb0004 && "proto" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).CurrentProtocol.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "CurrentProtocol")
					return
				}
				zb0003 = "proto"
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

func (z *OnlineRoundParamsData) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *OnlineRoundParamsData) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *OnlineRoundParamsData) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*OnlineRoundParamsData)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *OnlineRoundParamsData) Msgsize() (s int) {
	s = 1 + 7 + msgp.Uint64Size + 7 + msgp.Uint64Size + 6 + (*z).CurrentProtocol.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *OnlineRoundParamsData) MsgIsZero() bool {
	return ((*z).OnlineSupply == 0) && ((*z).RewardsLevel == 0) && ((*z).CurrentProtocol.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func OnlineRoundParamsDataMaxSize() (s int) {
	s = 1 + 7 + msgp.Uint64Size + 7 + msgp.Uint64Size + 6 + protocol.ConsensusVersionMaxSize()
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StateProofVerificationContext) MarshalMsg(b []byte) (o []byte) {
	o = msgp.Require(b, z.Msgsize())
	// omitempty: check for empty values
	zb0001Len := uint32(4)
	var zb0001Mask uint8 /* 5 bits */
	if (*z).OnlineTotalWeight.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x2
	}
	if (*z).LastAttestedRound.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x4
	}
	if (*z).Version.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x8
	}
	if (*z).VotersCommitment.MsgIsZero() {
		zb0001Len--
		zb0001Mask |= 0x10
	}
	// variable map header, size zb0001Len
	o = append(o, 0x80|uint8(zb0001Len))
	if zb0001Len != 0 {
		if (zb0001Mask & 0x2) == 0 { // if not empty
			// string "pw"
			o = append(o, 0xa2, 0x70, 0x77)
			o = (*z).OnlineTotalWeight.MarshalMsg(o)
		}
		if (zb0001Mask & 0x4) == 0 { // if not empty
			// string "spround"
			o = append(o, 0xa7, 0x73, 0x70, 0x72, 0x6f, 0x75, 0x6e, 0x64)
			o = (*z).LastAttestedRound.MarshalMsg(o)
		}
		if (zb0001Mask & 0x8) == 0 { // if not empty
			// string "v"
			o = append(o, 0xa1, 0x76)
			o = (*z).Version.MarshalMsg(o)
		}
		if (zb0001Mask & 0x10) == 0 { // if not empty
			// string "vc"
			o = append(o, 0xa2, 0x76, 0x63)
			o = (*z).VotersCommitment.MarshalMsg(o)
		}
	}
	return
}

func (_ *StateProofVerificationContext) CanMarshalMsg(z interface{}) bool {
	_, ok := (z).(*StateProofVerificationContext)
	return ok
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StateProofVerificationContext) unmarshalMsg(bts []byte, validate bool) (o []byte, err error) {
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
			bts, err = (*z).LastAttestedRound.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "LastAttestedRound")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).VotersCommitment.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "VotersCommitment")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).OnlineTotalWeight.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "OnlineTotalWeight")
				return
			}
		}
		if zb0001 > 0 {
			zb0001--
			bts, err = (*z).Version.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "struct-from-array", "Version")
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
			(*z) = StateProofVerificationContext{}
		}
		for zb0001 > 0 {
			zb0001--
			field, bts, err = msgp.ReadMapKeyZC(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
			switch string(field) {
			case "spround":
				if validate && zb0004 && "spround" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).LastAttestedRound.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "LastAttestedRound")
					return
				}
				zb0003 = "spround"
			case "vc":
				if validate && zb0004 && "vc" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).VotersCommitment.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "VotersCommitment")
					return
				}
				zb0003 = "vc"
			case "pw":
				if validate && zb0004 && "pw" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).OnlineTotalWeight.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "OnlineTotalWeight")
					return
				}
				zb0003 = "pw"
			case "v":
				if validate && zb0004 && "v" < zb0003 {
					err = &msgp.ErrNonCanonical{}
					return
				}
				bts, err = (*z).Version.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Version")
					return
				}
				zb0003 = "v"
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

func (z *StateProofVerificationContext) UnmarshalMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, false)
}
func (z *StateProofVerificationContext) UnmarshalValidateMsg(bts []byte) (o []byte, err error) {
	return z.unmarshalMsg(bts, true)
}
func (_ *StateProofVerificationContext) CanUnmarshalMsg(z interface{}) bool {
	_, ok := (z).(*StateProofVerificationContext)
	return ok
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StateProofVerificationContext) Msgsize() (s int) {
	s = 1 + 8 + (*z).LastAttestedRound.Msgsize() + 3 + (*z).VotersCommitment.Msgsize() + 3 + (*z).OnlineTotalWeight.Msgsize() + 2 + (*z).Version.Msgsize()
	return
}

// MsgIsZero returns whether this is a zero value
func (z *StateProofVerificationContext) MsgIsZero() bool {
	return ((*z).LastAttestedRound.MsgIsZero()) && ((*z).VotersCommitment.MsgIsZero()) && ((*z).OnlineTotalWeight.MsgIsZero()) && ((*z).Version.MsgIsZero())
}

// MaxSize returns a maximum valid message size for this message type
func StateProofVerificationContextMaxSize() (s int) {
	s = 1 + 8 + basics.RoundMaxSize() + 3 + crypto.GenericDigestMaxSize() + 3 + basics.MicroAlgosMaxSize() + 2 + protocol.ConsensusVersionMaxSize()
	return
}
