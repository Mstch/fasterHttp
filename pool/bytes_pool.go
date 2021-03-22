package pool

import (
	"math/bits"
	"runtime"
	"slowhttp/utils"
)

var bytesPool []map[int][][]byte

func init() {
	bytesPool = make([]map[int][][]byte, runtime.GOMAXPROCS(0))
	for i := range bytesPool {
		bytesPool[i] = make(map[int][][]byte)
	}
}
func GetBytes(need int) []byte {
	var localI int
	if need <= 64 {
		localI = 0
	} else {
		localI = bits.Len64(uint64(need)) - 6
	}
	blp := bytesPool[utils.ProcPin()]
	if _, ok := (blp)[localI]; !ok {
		(blp)[localI] = make([][]byte, 0)
	}
	if len((blp)[localI]) == 0 {
		(blp)[localI] = append((blp)[localI], make([]byte, 64<<localI))
	}
	buf := (blp)[localI][len((blp)[localI])-1]
	(blp)[localI] = (blp)[localI][:len((blp)[localI])-1]
	utils.ProcUnpin()
	return buf[:need]
}

func PutBytes(buf []byte) {
	if cap(buf) < 64 {
		return
	}
	localI := bits.Len64(uint64(cap(buf))) - 7
	blp := bytesPool[utils.ProcPin()]
	if _, ok := (blp)[localI]; !ok {
		(blp)[localI] = make([][]byte, 0)
	}
	(blp)[localI] = append((blp)[localI], buf[:cap(buf)])
	utils.ProcUnpin()
}
