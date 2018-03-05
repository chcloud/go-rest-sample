package resource

import (
	"strconv"

	"github.com/emicklei/go-restful"

	"github.com/chcloud/go-rest-sample/pkg/mapper"
)

// UserResource user rest apid
type UserResource struct {
}

// WebService  route for resource
func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/").To(u.hello))
	ws.Route(ws.GET("/{user-id}").To(u.findUser).Doc("get a user").
		Param(ws.PathParameter("user-id", "id of user").DataType("number")).
		Writes(mapper.User{}))
	return ws
}
func (u UserResource) hello(request *restful.Request, res *restful.Response) {
	res.AddHeader("Content-Type", restful.MIME_JSON)
	res.Write([]byte("[\"hello\"]"))
}

func (u UserResource) findUser(request *restful.Request, res *restful.Response) {
	id, err := strconv.ParseInt(request.PathParameter("user-id"), 10, 64)
	if err != nil {
		//todo
	}
	um := &mapper.UserMapper{}
	user := um.FindOne(id)
	res.WriteEntity(user)

}
