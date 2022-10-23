module github.com/rifat-simoom/payment-system/internal/common

go 1.18

require github.com/sirupsen/logrus v1.9.0

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect

require (
		cloud.google.com/go v0.75.0
    	firebase.google.com/go/v4 v4.7.1
    	github.com/deepmap/oapi-codegen v1.9.0
    	github.com/dgrijalva/jwt-go v3.2.0+incompatible
    	github.com/go-chi/chi/v5 v5.0.5
    	github.com/go-chi/cors v1.0.1
    	github.com/go-chi/render v1.0.1
    	github.com/golang/protobuf v1.5.2
    	github.com/google/uuid v1.1.2
    	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
    	github.com/mattn/go-colorable v0.1.8 // indirect
    	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b // indirect
    	github.com/onsi/ginkgo v1.12.0 // indirect
    	github.com/onsi/gomega v1.9.0 // indirect
    	github.com/pkg/errors v0.9.1
    	github.com/sirupsen/logrus v1.5.0
    	github.com/stretchr/testify v1.7.0
    	github.com/x-cray/logrus-prefixed-formatter v0.5.2
    	go.opencensus.io v0.22.5 // indirect
    	google.golang.org/api v0.40.0
    	google.golang.org/grpc v1.40.0
    	google.golang.org/protobuf v1.27.1
)