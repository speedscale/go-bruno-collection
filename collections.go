package collection

// EnvironmentVariablesSchema represents the schema for environment variables
type EnvironmentVariablesSchema struct {
	UID     UIDSchema
	Name    *string `validate:"omitempty"`
	Value   *string `validate:"omitempty"`
	Type    string  `validate:"required,oneof=text"`
	Enabled bool    `validate:"required"`
	Secret  *bool   `validate:"omitempty"`
}

// EnvironmentSchema represents the schema for environment
type EnvironmentSchema struct {
	UID       UIDSchema
	Name      string                       `validate:"required,min=1"`
	Variables []EnvironmentVariablesSchema `validate:"required,dive"`
}

// CollectionSchema represents the schema for collection
type CollectionSchema struct {
	Version              string `validate:"required,oneof=1"`
	UID                  UIDSchema
	Name                 string                 `validate:"required,min=1"`
	Items                []ItemSchema           `validate:"dive"`
	ActiveEnvironmentUID *string                `validate:"omitempty,len=21,alphanum"`
	Environments         []EnvironmentSchema    `validate:"dive"`
	Pathname             *string                `validate:"omitempty"`
	RunnerResult         map[string]interface{} `validate:"omitempty"`
	CollectionVariables  map[string]interface{} `validate:"omitempty"`
	BrunoConfig          map[string]interface{} `validate:"omitempty"`
}

// KeyValueSchema represents the schema for key-value pairs
type KeyValueSchema struct {
	UID         UIDSchema
	Name        *string `validate:"omitempty"`
	Value       *string `validate:"omitempty"`
	Description *string `validate:"omitempty"`
	Enabled     *bool   `validate:"omitempty"`
}

// VarsSchema represents the schema for vars
type VarsSchema struct {
	UID         UIDSchema
	Name        *string `validate:"omitempty"`
	Value       *string `validate:"omitempty"`
	Description *string `validate:"omitempty"`
	Enabled     *bool   `validate:"omitempty"`
	Local       *bool   `validate:"omitempty"`
}

// GraphqlBodySchema represents the schema for GraphQL body
type GraphqlBodySchema struct {
	Query     *string `validate:"omitempty"`
	Variables *string `validate:"omitempty"`
}

// MultipartFormSchema represents the schema for multipart form
type MultipartFormSchema struct {
	UID         UIDSchema
	Type        string  `validate:"required,oneof=file text"`
	Name        *string `validate:"omitempty"`
	Value       *string `validate:"omitempty"`
	Description *string `validate:"omitempty"`
	Enabled     *bool   `validate:"omitempty"`
}

// RequestBodySchema represents the schema for request body
type RequestBodySchema struct {
	Mode           string                `validate:"required,oneof=none json text xml formUrlEncoded multipartForm graphql sparql"`
	JSON           *string               `validate:"omitempty"`
	Text           *string               `validate:"omitempty"`
	XML            *string               `validate:"omitempty"`
	Sparql         *string               `validate:"omitempty"`
	FormUrlEncoded []KeyValueSchema      `validate:"omitempty,dive"`
	MultipartForm  []MultipartFormSchema `validate:"omitempty,dive"`
	Graphql        *GraphqlBodySchema    `validate:"omitempty"`
}

// AuthAwsV4Schema represents the schema for AWS V4 authentication
type AuthAwsV4Schema struct {
	AccessKeyId     *string `validate:"omitempty"`
	SecretAccessKey *string `validate:"omitempty"`
	SessionToken    *string `validate:"omitempty"`
	Service         *string `validate:"omitempty"`
	Region          *string `validate:"omitempty"`
	ProfileName     *string `validate:"omitempty"`
}

// AuthBasicSchema represents the schema for basic authentication
type AuthBasicSchema struct {
	Username *string `validate:"omitempty"`
	Password *string `validate:"omitempty"`
}

// AuthBearerSchema represents the schema for bearer authentication
type AuthBearerSchema struct {
	Token *string `validate:"omitempty"`
}

// AuthDigestSchema represents the schema for digest authentication
type AuthDigestSchema struct {
	Username *string `validate:"omitempty"`
	Password *string `validate:"omitempty"`
}

// Oauth2Schema represents the schema for OAuth2
type Oauth2Schema struct {
	GrantType        string  `validate:"required,oneof=client_credentials password authorization_code"`
	Username         *string `validate:"omitempty"`
	Password         *string `validate:"omitempty"`
	CallbackUrl      *string `validate:"omitempty"`
	AuthorizationUrl *string `validate:"omitempty"`
	AccessTokenUrl   *string `validate:"omitempty"`
	ClientId         *string `validate:"omitempty"`
	ClientSecret     *string `validate:"omitempty"`
	Scope            *string `validate:"omitempty"`
	State            *string `validate:"omitempty"`
	Pkce             *bool   `validate:"omitempty"`
}

// AuthSchema represents the schema for authentication
type AuthSchema struct {
	Mode   string            `validate:"required,oneof=inherit none awsv4 basic bearer digest oauth2"`
	AwsV4  *AuthAwsV4Schema  `validate:"omitempty"`
	Basic  *AuthBasicSchema  `validate:"omitempty"`
	Bearer *AuthBearerSchema `validate:"omitempty"`
	Digest *AuthDigestSchema `validate:"omitempty"`
	Oauth2 *Oauth2Schema     `validate:"omitempty"`
}

// RequestParamsSchema represents the schema for request parameters
type RequestParamsSchema struct {
	UID         UIDSchema
	Name        *string `validate:"omitempty"`
	Value       *string `validate:"omitempty"`
	Description *string `validate:"omitempty"`
	Type        string  `validate:"required,oneof=query path"`
	Enabled     *bool   `validate:"omitempty"`
}

// ScriptSchema represents the schema for script
type ScriptSchema struct {
	Req *string `validate:"omitempty"`
	Res *string `validate:"omitempty"`
}

// VarsReqResSchema represents the schema for vars in request and response
type VarsReqResSchema struct {
	Req []VarsSchema `validate:"omitempty,dive"`
	Res []VarsSchema `validate:"omitempty,dive"`
}

// RequestSchema represents the schema for request
type RequestSchema struct {
	URL        string                `validate:"required"`
	Method     string                `validate:"required,oneof=GET POST PUT DELETE PATCH HEAD OPTIONS"`
	Headers    []KeyValueSchema      `validate:"required,dive"`
	Params     []RequestParamsSchema `validate:"dive"` // Bruno does not always output these
	Auth       AuthSchema            `validate:"required"`
	Body       RequestBodySchema     `validate:"required"`
	Script     ScriptSchema          `validate:"required"`
	Vars       *VarsReqResSchema     `validate:"omitempty"`
	Assertions []KeyValueSchema      `validate:"omitempty,dive"`
	Tests      *string               `validate:"omitempty"`
	Docs       *string               `validate:"omitempty"`
}

// ItemSchema represents the schema for item (aka a Request)
type ItemSchema struct {
	UID         UIDSchema
	Type        string         `validate:"required,oneof=http graphql http-request graphql-request folder js"`
	Seq         int            `validate:"omitempty,min=1"`
	Name        string         `validate:"required,min=1"`
	Request     *RequestSchema `validate:"required_if=Type http-request Type graphql-request'"`
	FileContent *string        `validate:"omitempty"`
	Items       []ItemSchema   `validate:"omitempty,dive"`
	Filename    *string        `validate:"omitempty"`
	Pathname    *string        `validate:"omitempty"`
}

// CreateCollection returns a new Collection.
func CreateCollection(name string, desc string) CollectionSchema {
	return CollectionSchema{
		Name:    name,
		Version: "1", // as of this commit, 1 was the only version allowed
	}
}

// AddItem appends an item (ItemSchema) to the existing Items slice. This is how you add a new
// request to the Collection.
//
// CollectionSchema is passed as a pointer so changes are preserved on the heap.
func AddItem(cs *CollectionSchema, item ItemSchema) {
	cs.Items = append(cs.Items, item)
}

// CreateRequest returns a minimally populated RequestSchema. The return value will pass data
// validation because the minimum set of fields is filled in with default values.
func CreateRequest(url string, method string) *RequestSchema {
	return &RequestSchema{
		URL:     url,
		Method:  method,
		Auth:    AuthSchema{Mode: "inherit"},
		Body:    RequestBodySchema{Mode: "none"},
		Headers: []KeyValueSchema{},
		Script:  ScriptSchema{},
	}
}
