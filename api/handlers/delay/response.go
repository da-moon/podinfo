package delay

// Response struct represents this endpoints JSON
// Response
//
//go:generate gomodifytags -override -file $GOFILE -struct Response -add-tags json,yaml,mapstructure -w -transform snakecase
type Response struct {
	Delay uint64 `json:"delay" mapstructure:"delay" yaml:"delay"`
}
