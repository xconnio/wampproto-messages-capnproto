using Go = import "/go.capnp";

@0xb8e9720453f67577;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Subscribe {
    requestID @0 :Int64;
    topic @1 :Text;
}
