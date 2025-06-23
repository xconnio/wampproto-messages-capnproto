using Go = import "/go.capnp";

@0xdf622ac58e2bb8a5;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Result {
    requestID @0 :Int64;
}
