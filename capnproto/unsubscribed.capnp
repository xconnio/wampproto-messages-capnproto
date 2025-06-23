using Go = import "/go.capnp";

@0xafd1d51f3d0a000a;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/go/messages");

struct Unsubscribed {
    requestID @0 :Int64;
}
