using Go = import "/go.capnp";

@0xe9675a9c2636038d;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/go/messages");

struct Call {
    requestID @0 :Int64;
    procedure @1 :Text;
}
