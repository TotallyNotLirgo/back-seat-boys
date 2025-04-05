package parser

import "github.com/TotallyNotLirgo/back-seat-boys/models"

func (c *Parser) CheckAccessAcceptOnlySelf(id int) error {
    ctx := c.Request.Context()
    perm, ok := ctx.Value("permissions").(models.UserResponse)
    if !ok {
        return c.WriteError(401, "unauthorized")
    }
    if perm.UserId == id {
        return nil
    }
    return c.WriteError(403, "forbidden")
}

func (c *Parser) CheckAccessAcceptSelf(id int, role models.Role) error {
    ctx := c.Request.Context()
    perm, ok := ctx.Value("permissions").(models.UserResponse)
    if !ok {
        return c.WriteError(401, "unauthorized")
    }
    if perm.UserId == id || perm.Role.GreaterEqual(role) {
        return nil
    }
    return c.WriteError(403, "forbidden")
}

func (c *Parser) CheckAccess(role models.Role) error {
    ctx := c.Request.Context()
    perm, ok := ctx.Value("permissions").(models.UserResponse)
    if !ok {
        return c.WriteError(401, "unauthorized")
    }
    if perm.Role.GreaterEqual(role) {
        return nil
    }
    return c.WriteError(403, "forbidden")
}
