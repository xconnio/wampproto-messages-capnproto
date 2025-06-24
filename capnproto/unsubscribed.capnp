using Go = import "/go.capnp";

@0xafd1d51f3d0a000a;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Unsubscribed {
    requestID @0 :Int64;
}
