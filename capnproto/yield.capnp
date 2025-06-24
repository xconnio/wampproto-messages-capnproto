using Go = import "/go.capnp";

@0xc4ebc781b16b19ff;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Yield {
    requestID @0 :Int64;
}
