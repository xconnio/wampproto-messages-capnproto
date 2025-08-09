using Go = import "/go.capnp";

@0x9e9ea8de4c951ce7;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Cancel {
    requestID @0 :UInt64;
}
