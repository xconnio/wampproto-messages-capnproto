using Go = import "/go.capnp";

@0xa54fcc87aefc9866;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Abort {
    reason @0 :Text;
}
