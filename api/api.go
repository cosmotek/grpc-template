package api

//go:generate buf generate proto

import _ "embed"

//go:embed gen/proto/openapiv2/v1/api/service.swagger.json
var SwaggerSpecJSON []byte
