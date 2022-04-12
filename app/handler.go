package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	rep "quand/app/repository"
)

func (h *Handler) notFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "notFound.html", nil)
}

func (h *Handler) redirectToHome(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, "/home")
}

func (h *Handler) redirectToProfile(c *gin.Context) {
	name, err := c.Cookie("Username")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/qs/profile/login")
	}
	c.Redirect(http.StatusSeeOther, "/qs/profile/"+name)
}

func (h *Handler) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{"Url": c.Request.Host + c.Request.RequestURI})
}

func (h *Handler) profile(c *gin.Context) {
	nameParam := c.Param("name")
	_, err := c.Cookie("Username")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/qs/profile/login")
	} else {
		user, err := rep.GetUser(nameParam)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.HTML(http.StatusOK, "profile.html", user)
	}
}

func (h *Handler) login(c *gin.Context) {
	name, err := c.Cookie("Username")
	if err != nil {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		c.Redirect(http.StatusSeeOther, "/qs/profile/"+name)
	}
}

func (h *Handler) question(c *gin.Context) {
	q := rep.GetQuestion()
	c.HTML(http.StatusOK, "question.html", gin.H{"Text": q.Text, "Kind": q.Kind, "Author": q.Author})
}

func (h *Handler) logout(c *gin.Context) {
	c.SetCookie("Username", "", -1, "", "/", false, false)
}

func (h *Handler) verifyEmail(c *gin.Context) {

}
