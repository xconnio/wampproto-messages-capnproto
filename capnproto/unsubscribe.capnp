using Go = import "/go.capnp";

@0xd6bce36744723ad9;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Unsubscribe {
    requestID @0 :Int64;
    subscriptionID @1 :Int64;
}
