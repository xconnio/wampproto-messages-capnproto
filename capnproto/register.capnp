using Go = import "/go.capnp";

@0x9fb47c34036028b4;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Register {
    requestID @0 :UInt64;
    procedure @1 :Text;
    # options
    invoke @2 :Invoke;
    match @3 :Match;

    enum Invoke {
        single     @0;
        roundrobin @1;
        random     @2;
        first      @3;
        last       @4;
    }

    enum Match {
        exact    @0;
        prefix   @1;
        wildcard @2;
    }
}
