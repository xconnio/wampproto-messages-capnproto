using Go = import "/go.capnp";

@0xe0ff291d03fb4fa7;
$Go.package("gen");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go");

struct Welcome {
    sessionID @0 :UInt64;
    authid @1 :Text;
    authrole @2 :Text;
    authmethod @3 :Text;
    authprovider @4 :Text;
    roles @5 :Roles;

    struct Roles {
        dealer @0 :Dealer;
        broker @1 :Broker;

        struct Dealer {
            callerIdentification @0 :Bool;
            callTimeout @1 :Bool;
            callCanceling @2 :Bool;
            progressiveCallResults @3 :Bool;
            patternBasedRegistration @4 :Bool;
            sharedRegistration @5 :Bool;
        }

        struct Broker {
            publisherIdentification @0 :Bool;
            publisherExclusion @1 :Bool;
            acknowledgeEventReceived @2: Bool;
            patternBasedSubscription @3 :Bool;
        }
    }
}
