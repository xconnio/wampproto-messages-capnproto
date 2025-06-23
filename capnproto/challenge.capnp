using Go = import "/go.capnp";

@0x88995a52f153a0f3;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Challenge {
    authMethod @0 :Text;
}
