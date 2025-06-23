using Go = import "/go.capnp";

@0xf0723389dc3e4b37;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Publish {
    requestID @0 :Int64;
    topic @1 :Text;
}
