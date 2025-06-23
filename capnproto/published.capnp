using Go = import "/go.capnp";

@0xf03ec1cc202d6900;
$Go.package("messages");
$Go.import("github.com/xconnio/wampproto-messages-capnproto/go/messages");

struct Published {
    requestID @0 :Int64;
    publicationID @1 :Int64;
}
