package services

//go:generate go run gen_interface.go -svc=action -p=action -sf=service_generated.go
//go:generate go run gen_interface.go -svc=app-registry -p=appregistry -sf=service_generated.go
//go:generate go run gen_interface.go -svc=catalog -p=catalog -sf=service_generated.go
//go:generate go run gen_interface.go -svc=collect -p=collect -sf=service_generated.go
//go:generate go run gen_interface.go -svc=forwarders -p=forwarders -sf=service_generated.go
//go:generate go run gen_interface.go -svc=identity -p=identity -sf=service_generated.go
//go:generate go run gen_interface.go -svc=ingest -p=ingest -sf=service_generated.go
//go:generate go run gen_interface.go -svc=kvstore -p=kvstore -sf=service_generated.go
//go:generate go run gen_interface.go -svc=ml -p=ml -sf=service_generated.go
//go:generate go run gen_interface.go -svc=search -p=search -sf=service.go -sf=service_generated.go
//go:generate go run gen_interface.go -svc=streams -p=streams -sf=service_generated.go
//go:generate go run gen_interface.go -svc=provisioner -p=provisioner -sf=service_generated.go

//go:generate go run gen_interface.go -svc=anyType -p=anyType -sf=service_generated.go
//go:generate go run gen_interface.go -svc=enum -p=enum -sf=service_generated.go
//go:generate go run gen_interface.go -svc=of -p=of -sf=service_generated.go
//go:generate go run gen_interface.go -svc=ordering1 -p=ordering1 -sf=service_generated.go
//go:generate go run gen_interface.go -svc=ordering2 -p=ordering2 -sf=service_generated.go
//go:generate go run gen_interface.go -svc=allSupported -p=allSupported -sf=service_generated.go
