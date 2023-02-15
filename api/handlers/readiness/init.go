package readiness

const (
	// Name stores a human friendly, param-cased
	// unique identifier name for this endpoint
	Name = "kubernetes-readiness-probe"
	// APIGroup stores API group (prefix) for this URI
	// the full URI is /<prefix>/<path>
	APIGroup = ""
	// Path represents the URI path of this endpoint
	Path = "/readyz"
	// Path represents the URI path in this API group

)

var (
	// Router is accessible from outside in case
	// other packages need to access it's state
	Router = &handler{} //nolint:gochecknoglobals // safe as it has mutex guard and access internal state is through getter/setter functions
)

func init() { // nolint:gochecknoinits // this is to ensure router starts with OK state
	Router.SetStatus(OK)
}
