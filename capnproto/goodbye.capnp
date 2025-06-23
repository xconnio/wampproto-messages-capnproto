using Go = import "/go.capnp";

@0xf24976474795665c;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Goodbye {
    reason @0 :Text;
}
