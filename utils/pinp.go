package utils

import _ "unsafe"

//go:linkname ProcPin runtime.procPin
func ProcPin() int

//go:linkname ProcUnpin runtime.procUnpin
func ProcUnpin()
