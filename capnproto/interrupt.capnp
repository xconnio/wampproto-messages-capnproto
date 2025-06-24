using Go = import "/go.capnp";

@0xdff7acac49214441;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Interrupt {
    requestID @0 :Int64;
}
