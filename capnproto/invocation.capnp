using Go = import "/go.capnp";

@0xddef303b1dcb4c1c;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Invocation {
    requestID @0 :UInt64;
    registrationID @1 :UInt64;
    payloadSerializerID @2 :UInt64;
    # details
    caller @3 :UInt64;
    callerAuthID @4 :Text;
    callerAuthRole @5 :Text;
    procedure @6 :Text;
}
