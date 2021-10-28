package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"parsons.com/fds/goserver/redisclient"
	"parsons.com/fds/goserver/sessionmanager"
	"parsons.com/fds/goserver/structs"
)

func HandlePost(c *gin.Context) {
	session := sessions.Default(c)
	cookie := session.Get("HTTP_COOKIE")
	if cookie == nil {
		newCookie := sessionmanager.CreateNewSession()
		session.Set("HTTP_COOKIE", newCookie)
		session.Save()
		cookie = newCookie
	}
	var req structs.RPCRequest
	c.BindJSON(&req)
	redisclient.RPCRequest(cookie.(string), req.Id, req.Method, req.Params)
	response := map[string]int{"id": req.Id}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": response})
}

func HandleGet(c *gin.Context) {
	session := sessions.Default(c)
	cookie := session.Get("HTTP_COOKIE")
	if cookie == nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "result": "Client not recognized"})
	} else {
		var req structs.RPCResponse
		c.BindJSON(&req)
		resultString := redisclient.RPCResponse(cookie.(string), req.Id)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": resultString})
	}
}
