module studygo

go 1.14

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	golang.org/x/tools v0.0.0-20201023174141-c8cfbd0f21e6 // indirect
	rsc.io/quote v1.5.2
	studygo/pkg/go-home v0.0.0-00010101000000-000000000000
)

replace studygo/pkg/go-home => ./pkg/go-home
