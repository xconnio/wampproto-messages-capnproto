using Go = import "/go.capnp";

@0x9a9bbbe2afa6a893;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Error {
    messageType @0 :UInt64;
    requestID @1 :UInt64;
    uri @2 :Text;
    payloadSerializerID @3 :UInt64;
}
