using Go = import "/go.capnp";

@0xc070d5b08d2eec09;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Event {
    subscriptionID @0 :UInt64;
    publicationID @1 :UInt64;
    payloadSerializerID @2 :UInt64;
    # options
    publisher @3 :UInt64;
    publisherAuthID @4 :Text;
    publisherAuthRole @5 :Text;
    topic @6 :Text;
}
