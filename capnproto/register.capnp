using Go = import "/go.capnp";

@0x9fb47c34036028b4;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Register {
    requestID @0 :Int64;
    procedure @1 :Text;
}
