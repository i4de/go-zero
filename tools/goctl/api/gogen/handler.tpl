package {{.PkgName}}

import (
	"net/http"
	{{if .HasRequest}}"github.com/i-Things/things/shared/errors"{{end}}
    "github.com/i-Things/things/shared/result"
	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			result.Http(w, r, nil, errors.Parameter.WithMsg("入参不正确:"+err.Error()))
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		{{if .HasResp}}result.Http(w, r, resp, err){{else}}result.Http(w, r, nil, err){{end}}
	}
}
