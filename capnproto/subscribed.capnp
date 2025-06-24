using Go = import "/go.capnp";

@0xf249f7c43e04d03e;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Subscribed {
    requestID @0 :Int64;
    subscriptionID @1 :Int64;
}
