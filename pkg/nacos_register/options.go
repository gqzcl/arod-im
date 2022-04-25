package nacos

type options struct {
	prefix  string
	weight  float64
	cluster string
	group   string
	kind    string
}

// Option is nacos option.
type Option func(o *options)

// WithPrefix with prefix path.
func WithPrefix(prefix string) Option {
	return func(o *options) { o.prefix = prefix }
}

// WithWeight with weight option.
func WithWeight(weight float64) Option {
	return func(o *options) { o.weight = weight }
}

// WithCluster with cluster option.
func WithCluster(cluster string) Option {
	return func(o *options) { o.cluster = cluster }
}

// WithGroup with group option.
func WithGroup(group string) Option {
	return func(o *options) { o.group = group }
}

// WithDefaultKind with default kind option.
func WithDefaultKind(kind string) Option {
	return func(o *options) { o.kind = kind }
}
