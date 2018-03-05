package main

import (
	// "html/template"
	// "io"
	"net/http"
	"path"
	"runtime"

	_ "github.com/chcloud/go-rest-sample/pkg/config"
	"github.com/chcloud/go-rest-sample/pkg/mapper"
	"github.com/chcloud/go-rest-sample/pkg/migrate"
	rs "github.com/chcloud/go-rest-sample/pkg/resource"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/labstack/echo"
	"github.com/chcloud/go-rest-sample/pkg/assets"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	_ "github.com/go-xorm/xorm"
	_ "github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/mysql"
	_ "github.com/spf13/viper"
	"gopkg.in/logger.v1"
)

// func (p PackageResource)

const (
//HomeURL for qtrader.io home page URL
// HomeURL = "https://www.xgopkg.com"
)

func init() {
	log.SetOutputLevel(0)
	mapper.Connect()
	migrate.Migrate()
}
func main() {
	urs := rs.UserResource{}
	restful.DefaultContainer.Add(urs.WebService())
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(),
		APIPath:     "/swagger.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	//todo env get swaagerr ui dis
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("No caller information")
	}
	log.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(assets.FS("swagger/dist"))))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Go REST Sample",
			Description: "RESTful framework integration for golang",
			Contact: &spec.ContactInfo{
				Name:  "go-rest-sample",
				Email: "dev_support@<yourcompany>.com",
				URL:   "http://<yourservice>.com",
			},
			License: &spec.License{
				Name: "MIT",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
	swo.Tags = []spec.Tag{{TagProps: spec.TagProps{
		Name:        "users",
		Description: "Managing users"}}}
}
