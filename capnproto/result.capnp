using Go = import "/go.capnp";

@0xdf622ac58e2bb8a5;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Result {
    requestID @0 :UInt64;
    payloadSerializerID @1 :UInt64;
}
