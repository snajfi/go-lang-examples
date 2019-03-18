package main

import (
	"math"
)

// Variables declaration can be wrapped into 'var' block
var (
	logicalValue = true // Boolean value (true,false)
)

// Variable can be also declared without var block

var sentence string // Variable holds strings.String can be empty, but not nil. Strings are imutable.

// Signed integers
var eightBitInt int8 = math.MaxInt8       // -128, 127
var sixteenBitInt int16 = math.MaxInt16   // -32768, 32767
var thirtyTwoBitint int32 = math.MaxInt32 // -2147483648, 2147483647
var sixtyFourBitInt int64 = math.MaxInt64 // -9223372036854775808, 9223372036854775807

// Unsigned integers / 0 is min value
var uEightBitInt uint8 = math.MaxUint8       // max 255
var uSixteenBitInt uint16 = math.MaxUint16   // max 65535
var uThirtyTwoBitInt uint32 = math.MaxUint32 // max 4294967295
var uSixtyFourBitInt uint64 = math.MaxUint64 // max 18446744073709551615

var pointerAddressHolder uintptr // Integer type, which is large enough to hold address of any pointer.

var integer int   // Lengh depends on system. Is 32bit wide on 32bit system and 64bit wide on 64bit system.
var uInteger uint // Lengh depends on system. Is 32bit wide on 32bit system and 64bit wide on 64bit system.

// Floats
var thirtyTwoBitFloat float32
var sixtyFourBitFloat float64

// Complex numbers
var sixtyFourBitComplex complex64              // Float32 real and imaginary parts
var oneHundredTwentyEightBitComplex complex128 // Float64 real and imaginary parts

// Aliases
var aliasForUnit8 byte // Alias for

func main() {

}

func conditionWithLogicalValue() {
	if logicalValue {
		println("Our bool variable was set to true.")
	}
}
