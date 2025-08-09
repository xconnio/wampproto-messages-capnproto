using Go = import "/go.capnp";

@0xd57bb24e47819e45;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Registered {
    requestID @0 :UInt64;
    registrationID @1 :UInt64;
}
