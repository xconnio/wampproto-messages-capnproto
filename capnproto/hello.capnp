using Go = import "/go.capnp";

@0xf20692942fafaf1e;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Hello {
    realm @0 :Text;
    authid @1 :Text;
    authmethods @2 :List(Text);
    roles @3 :Roles;

    struct Roles {
        caller @0 :Caller;
        callee @1 :Callee;
        publisher @2 :Publisher;
        subscriber @3 :Subscriber;

        struct Caller {
            callerIdentification @0 :Bool;
            callTimeout @1 :Bool;
            callCanceling @2 :Bool;
            progressiveCallResults @3 :Bool;
        }

        struct Callee {
            callerIdentification @0 :Bool;
            callTimeout @1 :Bool;
            callCanceling @2 :Bool;
            progressiveCallResults @3 :Bool;
            patternBasedRegistration @4 :Bool;
            sharedRegistration @5 :Bool;
        }

        struct Publisher {
            publisherIdentification @0 :Bool;
            publisherExclusion @1 :Bool;
            subscriberBlackWhiteListing @2 :Bool;
            acknowledgeEventReceived @3: Bool;
        }

        struct Subscriber {
            publisherIdentification @0 :Bool;
            patternBasedSubscription @1 :Bool;
            eventHistory @2 :Bool;
        }
    }
}
