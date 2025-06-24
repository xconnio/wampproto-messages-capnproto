using Go = import "/go.capnp";

@0xa98624701878b311;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Unregister {
    requestID @0 :Int64;
    registrationID @1 :Int64;
}
