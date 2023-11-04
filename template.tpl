{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

const (
    {{- range .MethodSets}}
        Operation{{$svrType}}{{.OriginalName}} = "/{{$svrName}}/{{.OriginalName}}"
    {{- end}}
)

type {{.ServiceType}}HTTPServer interface {
{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func Register{{.ServiceType}}HTTPServer(app *fiber.App, srv {{.ServiceType}}HTTPServer) {
	r := app.Group("/")
	{{- range .Methods}}
	r.{{.Method}}("{{.Path}}", _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv))
	{{- end}}
}

{{range .Methods}}
func _{{$svrType}}_{{.Name}}{{.Num}}_HTTP_Handler(srv {{$svrType}}HTTPServer) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var in {{.Request}}
		{{- if .HasBody}}
		if err := c.BodyParser(&in{{.Body}}); err != nil {
			return err
		}
		{{- end}}
		if err := c.QueryParser(&in); err != nil {
			return err
		}
		{{- if .HasVars}}
		if err := c.ParamsParser(&in); err != nil {
			return err
		}
		{{- end}}
		v, err := protovalidate.New(
			protovalidate.WithFailFast(true),
		)
		if err != nil {
			return err
		}
		if err = v.Validate(&in); err != nil {
			return err
		}
		out, err := srv.{{.Name}}(c.Context(), &in)
        if err != nil {
            return err
        }
        return c.JSON(fiber.Map{
            "code":    fiber.StatusOK,
            "message": "success",
            "data":    out,
        })
	}
}
{{end}}

// Unimplemented{{.ServiceType}}Server must be embedded to have forward compatible implementations.
type Unimplemented{{.ServiceType}}Server struct {
}

{{- range .MethodSets}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	func (Unimplemented{{$svrType}}Server) {{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error) {
    	return nil, fiber.NewError(fiber.StatusNotImplemented, "method {{.Name}} not implemented")
    }
{{- end}}