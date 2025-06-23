using Go = import "/go.capnp";

@0x9e9ea8de4c951ce7;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Cancel {
    requestID @0 :Int64;
}
