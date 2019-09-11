//go:generate protoc -I=. -I=../.. --go_out=paths=source_relative:. inner_device.proto inner_gateway.proto

package api
