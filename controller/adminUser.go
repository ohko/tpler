package controller

import (
	"net/http"
	"strconv"
	"strings"

	"tpler/common"
	"tpler/model"

	"github.com/ohko/hst"
)

// AdminUserController 用户管理控制器
type AdminUserController struct {
	controller
}

// List 用户列表
func (o *AdminUserController) List(ctx *hst.Context) {
	if ctx.IsAjax() {
		startRow, _ := strconv.Atoi(ctx.R.FormValue("startRow"))
		endRow, _ := strconv.Atoi(ctx.R.FormValue("endRow"))
		total, us, err := model.DBUser.List(startRow, endRow-startRow)
		if err != nil {
			o.renderAdminError(ctx, err.Error())
		}

		ctx.JSON2(200, 0, map[string]interface{}{
			"success": true,
			"total":   total,
			"rows":    us,
		})
	}

	o.renderAdmin(ctx, nil, "admin/user/list.html")
}

// Add 增加用户
func (o *AdminUserController) Add(ctx *hst.Context) {
	if ctx.R.Method == "GET" {
		o.renderAdmin(ctx, nil, "admin/user/add.html")
	}

	u := &model.User{
		User:  ctx.R.FormValue("User"),
		Pass:  string(common.Hash([]byte(ctx.R.FormValue("Pass")))),
		Email: ctx.R.FormValue("Email"),
	}
	if err := model.DBUser.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	if ctx.IsAjax() {
		ctx.JSON2(200, 0, "ok")
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}

// Detail 查看用户
func (o *AdminUserController) Detail(ctx *hst.Context) {
	id, _ := strconv.Atoi(ctx.R.FormValue("ID"))
	u, err := model.DBUser.Get(id)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.IsAjax() {
		ctx.JSON2(200, 0, u)
	}
	o.renderAdmin(ctx, u, "admin/user/edit.html")
}

// Edit 编辑用户
func (o *AdminUserController) Edit(ctx *hst.Context) {
	id, _ := strconv.Atoi(ctx.R.FormValue("ID"))
	u, err := model.DBUser.Get(id)
	if err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.R.Method == "GET" {
		if ctx.IsAjax() {
			ctx.JSON2(200, 0, u)
		}
		o.renderAdmin(ctx, u, "admin/user/edit.html")
	}

	pass := ctx.R.FormValue("Pass")
	if pass != "" {
		u.Pass = string(common.Hash([]byte(pass)))
	}
	u.Email = ctx.R.FormValue("Email")
	if err := u.Save(u); err != nil {
		o.renderAdminError(ctx, err.Error())
	}
	if ctx.IsAjax() {
		ctx.JSON2(200, 0, "ok")
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}

// Delete 删除用户
func (o *AdminUserController) Delete(ctx *hst.Context) {
	_ids := strings.Split(ctx.R.FormValue("IDs"), ",")
	ids := []int{}
	for _, v := range _ids {
		tmp, _ := strconv.Atoi(v)
		if tmp == 0 {
			continue
		}
		ids = append(ids, tmp)
	}
	if err := model.DBUser.Delete(ids); err != nil {
		o.renderAdminError(ctx, err.Error())
	}

	if ctx.IsAjax() {
		ctx.JSON2(200, 0, "ok")
	}
	http.Redirect(ctx.W, ctx.R, "/admin_user/list", http.StatusFound)
}
