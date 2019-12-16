package chunk

type (
	Constraint func(string) bool

	Option func(*Chunk)

	Chunk struct {
		name        string
		constraints []Constraint
	}
)

func (ch *Chunk) Match(rch string) (name, string, value interface{}, ok bool) {
	name = ch.name
	ok = false
	for _, check := range ch.constraints {
		if ok = check(rch); !ok {
			return
		}
	}
}

func New(name string, opts ...Option) *Chunk {
	chunk := &Chunk{
		name: name,
	}
	for _, opt := range opts {
		opt(chunk)
	}
	return chunk
}

func WithCustomMatch(mfn func(string) bool) Option {
	return func(ch *Chunk) {
		ch.constraints = append(ch.constraints, mfn)
	}
}

func WithExactMatch(pch string) Option {
	return func(ch *Chunk) {
		ch.constraints = append(ch.constraints, func(rch string) bool {
			return pch == rch
		})
	}
}

func WithDictMatch(dict []string) Option {
	return func(ch *Chunk) {
		ch.constraints = append(ch.constraints, func(rch string) bool {
			for _, pch := range dict {
				if rch == pch {
					return true
				}
			}
			return false
		})
	}
}
