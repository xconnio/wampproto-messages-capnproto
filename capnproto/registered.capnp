using Go = import "/go.capnp";

@0xd57bb24e47819e45;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Registered {
    requestID @0 :Int64;
    registrationID @1 :Int64;
}
