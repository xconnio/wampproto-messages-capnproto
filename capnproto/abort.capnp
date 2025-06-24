using Go = import "/go.capnp";

@0xa54fcc87aefc9866;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Abort {
    reason @0 :Text;
}
