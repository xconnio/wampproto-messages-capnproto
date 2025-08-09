using Go = import "/go.capnp";

@0xe9675a9c2636038d;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Call {
    requestID @0 :UInt64;
    procedure @1 :Text;
    payloadSerializerID @2 :UInt64;
}
