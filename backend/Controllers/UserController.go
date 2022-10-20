package controllers

import "github.com/gofiber/fiber/v2"

func UserIndex(c *fiber.Ctx) error {
	return c.SendString("index1")
}
func UserStore(c *fiber.Ctx) error {
	return c.SendString("index2")
}
func UserUpdate(c *fiber.Ctx) error {
	return c.SendString("index3")
}
func UserGetId(c *fiber.Ctx) error {
	return c.SendString("index4")
}
func UserDestroy(c *fiber.Ctx) error {
	return c.SendString("index5")
}
