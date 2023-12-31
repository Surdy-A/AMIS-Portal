package controllers

import (
	"github.com/Surdy-A/amis_portal/internal/modules/user/requests/auth"
	UserService "github.com/Surdy-A/amis_portal/internal/modules/user/services"
	"github.com/Surdy-A/amis_portal/pkg/converters"
	"github.com/Surdy-A/amis_portal/pkg/errors"
	"github.com/Surdy-A/amis_portal/pkg/html"
	"github.com/Surdy-A/amis_portal/pkg/old"
	"github.com/Surdy-A/amis_portal/pkg/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate the request
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		errors.Init()
		errors.Add("Email", "Email address already exists")
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/register")
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// after creating the user > redirect to home page
	log.Printf("The user created successfully with a name %s \n", user.Username)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// validate the request
	var request auth.LoginRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		errors.Init()
		errors.SetFromErrors(err)
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		errors.Init()
		errors.Add("email", err.Error())
		sessions.Set(c, "errors", converters.MapToString(errors.Get()))

		old.Init()
		old.Set(c)
		sessions.Set(c, "old", converters.UrlValuesToString(old.Get()))

		c.Redirect(http.StatusFound, "/login")
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	// after creating the user > redirect to home page
	log.Printf("The user logged in successfully with a name %s \n", user.Username)
	c.Redirect(http.StatusFound, "/")
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")

	c.Redirect(http.StatusFound, "/")
}
