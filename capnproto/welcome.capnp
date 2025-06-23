using Go = import "/go.capnp";

@0xe0ff291d03fb4fa7;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Welcome {
    sessionID @0 :Int64;
}
