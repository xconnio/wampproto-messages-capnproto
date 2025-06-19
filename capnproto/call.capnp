using Go = import "/go.capnp";

@0xe9675a9c2636038d;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Call {
    procedure @0 :Text;
}
