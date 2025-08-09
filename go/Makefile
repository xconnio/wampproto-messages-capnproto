gen-schema:
	mkdir -p gen
	rm -rf gen/*
	capnp compile -I $(shell go list -m -f '{{.Dir}}' capnproto.org/go/capnp/v3)/std -o go:gen --src-prefix ../capnproto ../capnproto/* 
