package render

var InputsTemplate string = `
package {{ .PackageName }}

{{- range $obj := .EnumTypes}}
type {{ $obj.Name }} string
const (
	{{- range $value := $obj.Values }}
			{{ $obj.Name }}_{{ $value }} {{ $obj.Name }} = "{{ $value }}"
	{{- end }}
)

{{ end }}


{{- range $obj := .InputTypes }}
type {{$obj.Name}} struct {
{{- range $field := $obj.Fields }}
	{{- if $field.Type.IsList }}
		{{ pascal $field.Name }} {{if $field.Type.IsListNullable}}*{{end}}[]{{if $field.Type.IsTypeNullable}}*{{end}}{{$field.TypeName}} ~~json:"{{$field.Name}}"~~
	{{- else }}
		{{ pascal $field.Name }} {{if $field.Type.IsTypeNullable}}*{{end}}{{$field.TypeName}} ~~json:"{{$field.Name}}"~~
	{{- end }}
{{- end }}
}

{{ end }}
`