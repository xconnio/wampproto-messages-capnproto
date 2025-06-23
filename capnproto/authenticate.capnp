using Go = import "/go.capnp";

@0xe31d7a0f6a73f54d;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Authenticate {
    signature @0 :Text;
}
