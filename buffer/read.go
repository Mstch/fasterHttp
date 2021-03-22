package buffer

import (
	"bytes"
	"io"
	"net/url"
	"slowhttp/constant"
	"slowhttp/errors"
	"slowhttp/pool"
	"slowhttp/req"
	"slowhttp/utils"
	"strconv"
)

const MaxHeaderLineLen = 8192

type ReadBuffer struct {
	Buffer
}

func (b *ReadBuffer) ReadOnceFrom(r io.Reader) (n int64, err error) {
	i := b.grow(MinRead)
	b.buf = b.buf[:i]
	m, e := r.Read(b.buf[i:cap(b.buf)])
	if m < 0 {
		panic(errNegativeRead)
	}
	b.buf = b.buf[:i+m]
	n += int64(m)
	return n, e
}

func (b *ReadBuffer) ReadAReqFrom(req *req.Request, r io.Reader) error {
	//read request line
	for {
		crlfI := bytes.Index(b.buf[b.off:], constant.CRLF)
		if crlfI < 0 {
			if b.Len() > MaxHeaderLineLen {
				return errors.ErrInvalidRequestLine
			}
			_, err := b.ReadOnceFrom(r)
			if err != nil {
				return err
			}
		} else {
			line := b.buf[b.off : b.off+crlfI]
			b.off += crlfI + 2
			err := b.resolveRequestLine(req, line)
			if err != nil {
				return err
			}
			break
		}
	}

	//read headers
	for {
		crlfI := bytes.Index(b.buf[b.off:], constant.CRLF)
		if crlfI < 0 {
			if b.Len() > MaxHeaderLineLen {
				return errors.ErrInvalidHeader
			}
			_, err := b.ReadOnceFrom(r)
			if err != nil {
				return errors.NewHttpError(err.Error(), -1)
			}
		} else {
			line := b.buf[b.off : b.off+crlfI]
			b.off += crlfI + 2
			if len(line) == 0 { //header end ,start read body
				if contentLength, ok := req.Headers[constant.ContentLength]; ok {
					var err error
					req.ContentLength, err = strconv.Atoi(contentLength)
					if err != nil {
						return errors.ErrInvalidContentLength
					}
				} else {
					req.ContentLength = 0
				}
				break
			}
			err := b.resolveHeaderLine(req, line)
			if err != nil {
				return err
			}
		}
	}

	//read body
	needRead := req.ContentLength - b.Len()
	req.Body = pool.GetBytes(req.ContentLength)
	if needRead <= 0 { //has bean read  enough bytes into buf
		_, _ = b.Read(req.Body)
		return nil
	} else {
		n := 0
		if b.Len() > 0 {
			n, _ = b.Read(req.Body)
		}
		rdLen, err := io.ReadFull(r, req.Body[n:])
		if err != nil {
			return err
		}
		//j4check
		if rdLen != needRead {
			panic("invalid read")
		}
	}
	return nil
}

func (b *ReadBuffer) resolveRequestLine(req *req.Request, lineBuf []byte) error {

	//read method
	start := 0
	spaceIndex := bytes.IndexByte(lineBuf, ' ')
	if spaceIndex <= 0 {
		return errors.ErrInvalidRequestLine
	}
	method := utils.ConvBytesToString(lineBuf[start : start+spaceIndex])
	if constStrMethod, ok := constant.ValidMethod[method]; ok {
		req.Method = constStrMethod
	} else {
		return errors.ErrInvalidMethod
	}
	start += spaceIndex + 1
	if start >= len(lineBuf) {
		return errors.ErrInvalidRequestLine
	}

	//read path
	spaceIndex = bytes.IndexByte(lineBuf[start:], ' ')
	if spaceIndex <= 0 {
		return errors.ErrInvalidRequestLine
	}
	pathBuf := pool.GetBytes(spaceIndex)
	copy(pathBuf, lineBuf[start:start+spaceIndex])
	req.Path = utils.ConvBytesToString(pathBuf)
	start += spaceIndex + 1
	if start >= len(lineBuf) {
		return errors.ErrInvalidRequestLine
	}
	u, err := url.Parse(req.Path)
	if err != nil {
		return errors.ErrInvalidPath
	}
	req.URL = u

	//read http version
	if utils.ConvBytesToString(lineBuf[start:]) == constant.HttpVersion11 {
		req.Version = constant.HttpVersion11
	} else {
		return errors.ErrInvalidVersion
	}
	return nil
}

func (b *ReadBuffer) resolveHeaderLine(req *req.Request, line []byte) error {
	colonIndex := bytes.Index(line, constant.ColonSpace)
	if colonIndex <= 0 || colonIndex >= len(line)-2 {
		return errors.ErrInvalidHeader
	}
	var k, v string
	srck := utils.ConvBytesToString(line[:colonIndex])
	if ck, ok := constant.ConstHeaderk[srck]; ok {
		k = ck
	} else {
		k = string(line[:colonIndex])
	}
	vbuf := pool.GetBytes(len(line) - (colonIndex + 2))
	copy(vbuf, line[colonIndex+2:])
	v = utils.ConvBytesToString(vbuf)
	req.Headers[k] = v
	return nil
}
