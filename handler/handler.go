package handler

import (
	"slowhttp/req"
	"slowhttp/resp"
)

type Handler func(req *req.Request, res *resp.Response)
