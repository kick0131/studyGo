module studygo

go 1.14

require (
	cloud.google.com/go/firestore v1.3.0
	firebase.google.com/go v3.13.0+incompatible
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/lib/pq v1.8.0
	github.com/stretchr/testify v1.4.0
	golang.org/x/tools v0.0.0-20201023174141-c8cfbd0f21e6 // indirect
	google.golang.org/api v0.34.0
	studygo/pkg/go-home v0.0.0-00010101000000-000000000000 // indirect
)

replace studygo/pkg/go-home => ./pkg/go-home
