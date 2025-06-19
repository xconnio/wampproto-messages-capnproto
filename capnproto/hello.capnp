using Go = import "/go.capnp";

@0xf20692942fafaf1e;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-capnproto/go/messages");

struct Hello {
    realm @0 :Text;
    authid @1 :Text;
    authmethods @2 :List(Text);
    roles @3 :Roles;

    struct Roles {
        caller @0 :Caller;
        callee @1 :Caller;
        publisher @2 :Caller;
        subscriber @3 :Caller;

        struct Caller {
            callerIdentification @0 :Bool;
        }
    }
}
