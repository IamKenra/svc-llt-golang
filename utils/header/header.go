package header

import (
	"svc-llt-golang/utils/response"

	"github.com/gofiber/fiber/v2"
)

const (
	XMemberHeader        = "X-Member"
	XMemberRequiredError = "X-Member header is required"
)

func ExtractXMember(ctx *fiber.Ctx) (string, error) {
	xMember := ctx.Get(XMemberHeader)
	if xMember == "" {
		return "", fiber.NewError(fiber.StatusBadRequest, XMemberRequiredError)
	}
	return xMember, nil
}

func ValidateAndExtractXMember(ctx *fiber.Ctx) (string, error) {
	xMember, err := ExtractXMember(ctx)
	if err != nil {
		response.BadRequest(ctx, XMemberRequiredError)
		return "", err
	}
	return xMember, nil
}