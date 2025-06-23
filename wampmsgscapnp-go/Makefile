gen-schema:
	mkdir -p gen
	rm -rf gen/*
	capnp compile -I $(shell go list -m -f '{{.Dir}}' github.com/capnproto/go-capnp)/std -o go:gen --src-prefix ../capnproto ../capnproto/* 
