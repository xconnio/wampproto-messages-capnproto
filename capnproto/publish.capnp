using Go = import "/go.capnp";

@0xf0723389dc3e4b37;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Publish {
    topic @0 :Text;
}
