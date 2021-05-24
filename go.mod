module test.com/gotemptest

go 1.15

require (
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/tkanos/gonfig v0.0.0-20210106201359-53e13348de2f
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace test.com/gotemptest/models => ./models
