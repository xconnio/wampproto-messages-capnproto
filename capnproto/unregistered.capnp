using Go = import "/go.capnp";

@0xf6f0840b4838627a;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Unregistered {
    requestID @0 :Int64;
}
