using Go = import "/go.capnp";

@0xd6bce36744723ad9;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Unsubscribe {
    requestID @0 :Int64;
    subscriptionID @1 :Int64;
}
