package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/gofiber/fiber/v3"
	"github.com/markbates/goth/gothic"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type fiberResponseWriter struct {
	ctx fiber.Ctx
}

func (w *fiberResponseWriter) Header() http.Header {
	return http.Header{}
}

func (w *fiberResponseWriter) Write(b []byte) (int, error) {
	return w.ctx.Write(b)
}

func (w *fiberResponseWriter) WriteHeader(statusCode int) {
	w.ctx.Status(statusCode)
}

func createHTTPResponseWriterFromFiberCtx(c fiber.Ctx) http.ResponseWriter {
	return &fiberResponseWriter{ctx: c}
}

func AuthProviderCallback(c fiber.Ctx) error {

	provider := c.Params("provider")

	req := fasthttpToHttp(c)

	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))

	res := createHTTPResponseWriterFromFiberCtx(c)

	user, err := gothic.CompleteUserAuth(res, req)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	fmt.Println(user)

	util.LogHandler("Usuário autenticado com sucesso", nil, "AuthProviderCallback")

	c.Redirect().Status(fiber.StatusFound).To("/")

	return nil
}

func LogoutProvider(c fiber.Ctx) error {
	provider := c.Params("provider")

	req := fasthttpToHttp(c)

	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))

	res := createHTTPResponseWriterFromFiberCtx(c)

	gothic.Logout(res, req)

	c.Redirect().Status(fiber.StatusFound).To("/")

	return nil
}

func AuthProvider(c fiber.Ctx) error {

	provider := c.Params("provider")

	req := fasthttpToHttp(c)

	req = req.WithContext(context.WithValue(req.Context(), "provider", provider))

	res := createHTTPResponseWriterFromFiberCtx(c)

	gothic.BeginAuthHandler(res, req)
	c.SendStatus(fiber.StatusOK)

	return nil
}

func fasthttpToHttp(c fiber.Ctx) *http.Request {
	var req *http.Request
	fasthttpadaptor.ConvertRequest(c.Context(), req, true)
	return req
}
