package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func StoreMedia(c fiber.Ctx) error {
	fmt.Println(c)
	return nil
}