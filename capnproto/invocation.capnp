using Go = import "/go.capnp";

@0xddef303b1dcb4c1c;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Invocation {
    requestID @0 :Int64;
    registrationID @1 :Int64;
}
