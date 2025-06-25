package parsers

import (
	"bytes"

	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-messages-capnproto/wampmsgscapnp-go/gen"
)

type Hello struct {
	gen *gen.Hello
}

func NewHelloFields(gen *gen.Hello) messages.HelloFields {
	return &Hello{gen: gen}
}

func (h *Hello) Realm() string {
	realm, _ := h.gen.Realm()
	return realm
}

func (h *Hello) AuthID() string {
	authid, _ := h.gen.Authid()
	return authid
}

func (h *Hello) AuthMethods() []string {
	list, err := h.gen.Authmethods()
	if err != nil {
		return nil
	}
	var authmethods []string
	for i := 0; i < list.Len(); i++ {
		authmethod, _ := list.At(i)
		authmethods = append(authmethods, authmethod)
	}
	return authmethods
}

func (h *Hello) AuthExtra() map[string]any {
	return map[string]any{}
}

func (h *Hello) Roles() map[string]any {
	roles := map[string]any{}
	genRoles, err := h.gen.Roles()
	if err != nil {
		return roles
	}

	if r, err := genRoles.Caller(); err == nil {
		roles["caller"] = map[string]any{
			"caller_identification":    r.CallerIdentification(),
			"call_timeout":             r.CallTimeout(),
			"call_canceling":           r.CallCanceling(),
			"progressive_call_results": r.ProgressiveCallResults(),
		}
	}
	if r, err := genRoles.Callee(); err == nil {
		roles["callee"] = map[string]any{
			"caller_identification":      r.CallerIdentification(),
			"call_timeout":               r.CallTimeout(),
			"call_canceling":             r.CallCanceling(),
			"progressive_call_results":   r.ProgressiveCallResults(),
			"pattern_based_registration": r.PatternBasedRegistration(),
			"shared_registration":        r.SharedRegistration(),
		}
	}
	if r, err := genRoles.Publisher(); err == nil {
		roles["publisher"] = map[string]any{
			"publisher_identification":   r.PublisherIdentification(),
			"publisher_exclusion":        r.PublisherExclusion(),
			"subscriber_blackwhite_list": r.SubscriberBlackWhiteListing(),
			"acknowledge_event_received": r.AcknowledgeEventReceived(),
		}
	}
	if r, err := genRoles.Subscriber(); err == nil {
		roles["subscriber"] = map[string]any{
			"publisher_identification":   r.PublisherIdentification(),
			"pattern_based_subscription": r.PatternBasedSubscription(),
			"event_history":              r.EventHistory(),
		}
	}

	return roles
}

func HelloToCapnproto(h *messages.Hello) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	hello, err := gen.NewHello(seg)
	if err != nil {
		return nil, err
	}

	if err := hello.SetRealm(h.Realm()); err != nil {
		return nil, err
	}
	if err := hello.SetAuthid(h.AuthID()); err != nil {
		return nil, err
	}

	methods, err := hello.NewAuthmethods(int32(len(h.AuthMethods())))
	if err != nil {
		return nil, err
	}
	for i, m := range h.AuthMethods() {
		_ = methods.Set(i, m)
	}
	if err := hello.SetAuthmethods(methods); err != nil {
		return nil, err
	}

	roles, err := hello.NewRoles()
	if err != nil {
		return nil, err
	}

	if callerMap, ok := h.Roles()["caller"].(map[string]any); ok {
		caller, err := roles.NewCaller()
		if err == nil {
			if v, ok := callerMap["caller_identification"].(bool); ok {
				caller.SetCallerIdentification(v)
			}
			if v, ok := callerMap["call_timeout"].(bool); ok {
				caller.SetCallTimeout(v)
			}
			if v, ok := callerMap["call_canceling"].(bool); ok {
				caller.SetCallCanceling(v)
			}
			if v, ok := callerMap["progressive_call_results"].(bool); ok {
				caller.SetProgressiveCallResults(v)
			}
			_ = roles.SetCaller(caller)
		}
	}

	if calleeMap, ok := h.Roles()["callee"].(map[string]any); ok {
		callee, err := roles.NewCallee()
		if err == nil {
			if v, ok := calleeMap["caller_identification"].(bool); ok {
				callee.SetCallerIdentification(v)
			}
			if v, ok := calleeMap["call_timeout"].(bool); ok {
				callee.SetCallTimeout(v)
			}
			if v, ok := calleeMap["call_canceling"].(bool); ok {
				callee.SetCallCanceling(v)
			}
			if v, ok := calleeMap["progressive_call_results"].(bool); ok {
				callee.SetProgressiveCallResults(v)
			}
			if v, ok := calleeMap["pattern_based_registration"].(bool); ok {
				callee.SetPatternBasedRegistration(v)
			}
			if v, ok := calleeMap["shared_registration"].(bool); ok {
				callee.SetSharedRegistration(v)
			}
			_ = roles.SetCallee(callee)
		}
	}

	if publisherMap, ok := h.Roles()["publisher"].(map[string]any); ok {
		publisher, err := roles.NewPublisher()
		if err == nil {
			if v, ok := publisherMap["publisher_identification"].(bool); ok {
				publisher.SetPublisherIdentification(v)
			}
			if v, ok := publisherMap["publisher_exclusion"].(bool); ok {
				publisher.SetPublisherExclusion(v)
			}
			if v, ok := publisherMap["subscriber_blackwhite_list"].(bool); ok {
				publisher.SetSubscriberBlackWhiteListing(v)
			}
			if v, ok := publisherMap["acknowledge_event_received"].(bool); ok {
				publisher.SetAcknowledgeEventReceived(v)
			}
			_ = roles.SetPublisher(publisher)
		}
	}

	if subscriberMap, ok := h.Roles()["subscriber"].(map[string]any); ok {
		subscriber, err := roles.NewSubscriber()
		if err == nil {
			if v, ok := subscriberMap["publisher_identification"].(bool); ok {
				subscriber.SetPublisherIdentification(v)
			}
			if v, ok := subscriberMap["pattern_based_subscription"].(bool); ok {
				subscriber.SetPatternBasedSubscription(v)
			}
			if v, ok := subscriberMap["event_history"].(bool); ok {
				subscriber.SetEventHistory(v)
			}
			_ = roles.SetSubscriber(subscriber)
		}
	}

	if err := hello.SetRoles(roles); err != nil {
		return nil, err
	}

	var data bytes.Buffer
	if err := capnp.NewEncoder(&data).Encode(msg); err != nil {
		return nil, err
	}

	byteValue := byte(messages.MessageTypeHello & 0xFF)
	return append([]byte{byteValue}, data.Bytes()...), nil
}

func CapnprotoToHello(data []byte) (*messages.Hello, error) {
	msg, err := capnp.NewDecoder(bytes.NewReader(data)).Decode()
	if err != nil {
		return nil, err
	}

	hello, err := gen.ReadRootHello(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewHelloWithFields(NewHelloFields(&hello)), nil
}
