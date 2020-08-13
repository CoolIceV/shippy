module cli-consignment

go 1.14

replace github.com/CoolIceV/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/CoolIceV/shippy/shippy-service-consignment v0.0.0-20200813095043-91a0a79c8487
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/grpc v1.31.0
)
