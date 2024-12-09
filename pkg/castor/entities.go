//
// Copyright (c) 2022 - for information on the respective copyright owner
// see the NOTICE file and/or the repository https://github.com/carbynestack/castor.
//
// SPDX-License-Identifier: Apache-2.0
//

package castor

// TupleList is a collection of a specific type of tuples.
type TupleList struct {
	Tuples []Tuple `json:"tuples"`
}

// Tuple describes the actual tuple and its shares.
type Tuple struct {
	Shares []Share `json:"shares"`
}

// Share represents a single share of a tuple with its shared value and mac data.
type Share struct {
	Value string `json:"value"`
	Mac   string `json:"mac"`
}

// SPDZProtocol describes the protocol used for the MPC computation.
type SPDZProtocol struct {
	Descriptor string
	Shorthand  string
	Domain     SPDZDomain
}

// Domain
type SPDZDomain struct {
	Shorthand string
}

var (
	GFP  = SPDZDomain{"GFP"}
	GF2n = SPDZDomain{"GF2n"}
)

var (
	// Default
	SPDZGfp  = SPDZProtocol{"SPDZ gfp", "p", GFP}
	SPDZGf2n = SPDZProtocol{"SPDZ gf2n_", "2", GF2n}
	// Semi
	SPDZGfpD  = SPDZProtocol{"gfp", "Dp", GFP}
	SPDZGf2nD = SPDZProtocol{"gf2n_", "D2", GF2n}
	// Binary triple protocol for above
	//   default, Tinier (difference with Tiny?)
	SPDZGfpBT = SPDZProtocol{"evaluation secret", "TT", GFP} // GFP here is arguable
	//   SemiBin
	SPDZGfpDBT = SPDZProtocol{"binary secret", "DB", GFP} // GFP here is arguable
)

// SupportedSPDZProtocols is a list of all SPDZProtocol supported by castor and ephemeral.
var SupportedSPDZProtocols = []SPDZProtocol{
	SPDZGfp,
	SPDZGf2n,
	SPDZGfpD,
	SPDZGf2nD,
	SPDZGfpBT,
	SPDZGfpDBT,
}

// TupleType describes a type of Tuples provided by Castor.
type TupleType struct {
	Name              string
	PreprocessingName string
	SpdzProtocol      SPDZProtocol
	BitLength         uint
}

var (
	// NOTE: There are secure and unsecure edaBits but they end up in the same filenames
	//		 So should we even try to make the distinction?

	// 40 is the default security parameter, they seem to be used in loops
	//    this is potentially configurable but not addressing that yet
	// 41 looks like the default security parameter + 1, used in single comparisons?
	// 64 default limb size or default bit_length or half the prime bit_length
	//    hardcoded in some other places in Carbyne Stack
	// 32 just to have an alternative to 64
	EdaBitGfp32      = TupleType{"EDABIT_GFP_32", "edaBits", SPDZGfp, 32}
	EdaBitGfp40      = TupleType{"EDABIT_GFP_40", "edaBits", SPDZGfp, 40}
	EdaBitGfp41      = TupleType{"EDABIT_GFP_41", "edaBits", SPDZGfp, 41}
	EdaBitGfp64      = TupleType{"EDABIT_GFP_64", "edaBits", SPDZGfp, 64}
	EdaBitGfp32D     = TupleType{"EDABIT_GFP_32", "edaBits", SPDZGfpD, 32}
	EdaBitGfp40D     = TupleType{"EDABIT_GFP_40", "edaBits", SPDZGfpD, 40}
	EdaBitGfp41D     = TupleType{"EDABIT_GFP_41", "edaBits", SPDZGfpD, 41}
	EdaBitGfp64D     = TupleType{"EDABIT_GFP_64", "edaBits", SPDZGfpD, 64}
	BinaryTripleGfp  = TupleType{"BINARY_TRIPLE_GFP", "Triples", SPDZGfpBT, 0}
	BinaryTripleGfpD = TupleType{"BINARY_TRIPLE_GFP", "Triples", SPDZGfpDBT, 0}
	DaBitGfp         = TupleType{"DABIT_GFP", "daBits", SPDZGfp, 0}
	DaBitGf2n        = TupleType{"DABIT_GF2N", "daBits", SPDZGf2n, 0}

	// BitGfp describes the Bits tuple type in the Mudulo a Prime domain.
	BitGfp = TupleType{"BIT_GFP", "Bits", SPDZGfp, 0}
	// BitGf2n describes the Bits tuple type in the GF(2^n) domain.
	BitGf2n = TupleType{"BIT_GF2N", "Bits", SPDZGf2n, 0}
	// InputMaskGfp describes the Inputs tuple type in the Mudulo a Prime domain.
	InputMaskGfp = TupleType{"INPUT_MASK_GFP", "Inputs", SPDZGfp, 0}
	// InputMaskGf2n describes the Inputs tuple type in the GF(2^n) domain.
	InputMaskGf2n = TupleType{"INPUT_MASK_GF2N", "Inputs", SPDZGf2n, 0}
	// InverseTupleGfp describes the Inverses tuple type in the Mudulo a Prime domain.
	InverseTupleGfp = TupleType{"INVERSE_TUPLE_GFP", "Inverses", SPDZGfp, 0}
	// InverseTupleGf2n describes the Inverses tuple type in the GF(2^n) domain.
	InverseTupleGf2n = TupleType{"INVERSE_TUPLE_GF2N", "Inverses", SPDZGf2n, 0}
	// SquareTupleGfp describes the Squares tuple type in the Mudulo a Prime domain.
	SquareTupleGfp = TupleType{"SQUARE_TUPLE_GFP", "Squares", SPDZGfp, 0}
	// SquareTupleGf2n describes the Squares tuple type in the GF(2^n) domain.
	SquareTupleGf2n = TupleType{"SQUARE_TUPLE_GF2N", "Squares", SPDZGf2n, 0}
	// MultiplicationTripleGfp describes the Triples tuple type in the Mudulo a Prime domain.
	MultiplicationTripleGfp = TupleType{"MULTIPLICATION_TRIPLE_GFP", "Triples", SPDZGfp, 0}
	// MultiplicationTripleGf2n describes the Triples tuple type in the GF(2^n) domain.
	MultiplicationTripleGf2n = TupleType{"MULTIPLICATION_TRIPLE_GF2N", "Triples", SPDZGf2n, 0}
	// BitGfp describes the DaBits tuple type in the Mudulo a Prime domain.
	DaBitGfpD = TupleType{"DABIT_GFP", "daBits", SPDZGfpD, 0}
	// BitGf2n describes the DaBits tuple type in the GF(2^n) domain.
	DaBitGf2nD = TupleType{"DABIT_GF2N", "daBits", SPDZGf2nD, 0}
	// BitGfp describes the Bits tuple type in the Mudulo a Prime domain.
	BitGfpD = TupleType{"BIT_GFP", "Bits", SPDZGfpD, 0}
	// BitGf2n describes the Bits tuple type in the GF(2^n) domain.
	BitGf2nD = TupleType{"BIT_GF2N", "Bits", SPDZGf2nD, 0}
	// InputMaskGfp describes the Inputs tuple type in the Mudulo a Prime domain.
	InputMaskGfpD = TupleType{"INPUT_MASK_GFP", "Inputs", SPDZGfpD, 0}
	// InputMaskGf2n describes the Inputs tuple type in the GF(2^n) domain.
	InputMaskGf2nD = TupleType{"INPUT_MASK_GF2N", "Inputs", SPDZGf2nD, 0}
	// InverseTupleGfp describes the Inverses tuple type in the Mudulo a Prime domain.
	InverseTupleGfpD = TupleType{"INVERSE_TUPLE_GFP", "Inverses", SPDZGfpD, 0}
	// InverseTupleGf2n describes the Inverses tuple type in the GF(2^n) domain.
	InverseTupleGf2nD = TupleType{"INVERSE_TUPLE_GF2N", "Inverses", SPDZGf2nD, 0}
	// SquareTupleGfp describes the Squares tuple type in the Mudulo a Prime domain.
	SquareTupleGfpD = TupleType{"SQUARE_TUPLE_GFP", "Squares", SPDZGfpD, 0}
	// SquareTupleGf2n describes the Squares tuple type in the GF(2^n) domain.
	SquareTupleGf2nD = TupleType{"SQUARE_TUPLE_GF2N", "Squares", SPDZGf2nD, 0}
	// MultiplicationTripleGfp describes the Triples tuple type in the Mudulo a Prime domain.
	MultiplicationTripleGfpD = TupleType{"MULTIPLICATION_TRIPLE_GFP", "Triples", SPDZGfpD, 0}
	// MultiplicationTripleGf2n describes the Triples tuple type in the GF(2^n) domain.
	MultiplicationTripleGf2nD = TupleType{"MULTIPLICATION_TRIPLE_GF2N", "Triples", SPDZGf2nD, 0}
)

// SupportedTupleTypes is a list of all tuple types supported by the castor client.
func SupportedTupleTypes(protocol SPDZProtocol) []TupleType {
	switch protocol {
	case SPDZGfp, SPDZGf2n:
		return []TupleType{
			EdaBitGfp32,
			EdaBitGfp40,
			EdaBitGfp41,
			EdaBitGfp64,
			BinaryTripleGfp,
			DaBitGfp,
			DaBitGf2n,
			BitGfp,
			BitGf2n,
			InputMaskGfp,
			InputMaskGf2n,
			InverseTupleGfp,
			InverseTupleGf2n,
			SquareTupleGfp,
			SquareTupleGf2n,
			MultiplicationTripleGfp,
			MultiplicationTripleGf2n,
		}
	case SPDZGfpD, SPDZGf2nD:
		return []TupleType{
			EdaBitGfp32D,
			EdaBitGfp40D,
			EdaBitGfp41D,
			EdaBitGfp64D,
			BinaryTripleGfpD,
			DaBitGfpD,
			DaBitGf2nD,
			BitGfpD,
			BitGf2nD,
			InputMaskGfpD,
			InputMaskGf2nD,
			InverseTupleGfpD,
			InverseTupleGf2nD,
			SquareTupleGfpD,
			SquareTupleGf2nD,
			MultiplicationTripleGfpD,
			MultiplicationTripleGf2nD,
		}
	default:
		return []TupleType{}
	}
}
