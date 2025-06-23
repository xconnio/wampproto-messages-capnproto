using Go = import "/go.capnp";

@0xc070d5b08d2eec09;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Event {
    subscriptionID @0 :Int64;
    publicationID @1 :Int64;
}
