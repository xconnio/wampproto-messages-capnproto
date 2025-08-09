using Go = import "/go.capnp";

@0xb8e9720453f67577;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Subscribe {
    requestID @0 :UInt64;
    topic @1 :Text;
    match @2 :Match;

    enum Match {
        exact    @0;
        prefix   @1;
        wildcard @2;
    }
}
