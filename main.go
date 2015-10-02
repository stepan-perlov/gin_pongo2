package gin_pongo2

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type (
	PongoProduction struct {
		Templates map[string]*pongo2.Template
		Path      string
	}

	PongoDebug struct {
		Path string
	}

	Pongo struct {
		Template *pongo2.Template
		Name     string
		Data     interface{}
	}
)

func NewProduction(path string) *PongoProduction {
	return &PongoProduction{map[string]*pongo2.Template{}, path}
}

func NewDebug(path string) *PongoDebug {
	return &PongoDebug{path}
}

func (p PongoProduction) Instance(name string, data interface{}) render.Render {
	var t *pongo2.Template
	if tmpl, ok := p.Templates[name]; ok {
		t = tmpl
	} else {
		tmpl := pongo2.Must(pongo2.FromFile(path.Join(p.Path, name)))
		p.Templates[name] = tmpl
		t = tmpl
	}
	return Pongo{
		Template: t,
		Name:     name,
		Data:     data,
	}
}

func (p PongoDebug) Instance(name string, data interface{}) render.Render {
	t := pongo2.Must(pongo2.FromFile(path.Join(p.Path, name)))
	return Pongo{
		Template: t,
		Name:     name,
		Data:     data,
	}
}

func (p Pongo) Render(w http.ResponseWriter) error {
	ctx := pongo2.Context(p.Data.(pongo2.Context))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return p.Template.ExecuteWriter(ctx, w)
}

func MakeContext(c *gin.Context, data map[string]interface{}) pongo2.Context {
	if c.Keys != nil {
		for key, value := range c.Keys {
			if _, exists := data[key]; !exists {
				data[key] = value
			}

		}
	}
	return data
}
