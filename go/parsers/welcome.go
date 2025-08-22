package parsers

import (
	"capnproto.org/go/capnp/v3"

	"github.com/xconnio/wampproto-go/messages"
	"github.com/xconnio/wampproto-serializer-capnproto/go/gen"
)

type Welcome struct {
	gen *gen.Welcome
}

func NewWelcomeFields(g *gen.Welcome) messages.WelcomeFields {
	return &Welcome{gen: g}
}

func (w *Welcome) SessionID() uint64 {
	return w.gen.SessionID()
}

func (w *Welcome) Details() map[string]any {
	authID, _ := w.gen.Authid()
	authRole, _ := w.gen.Authrole()
	authMethod, _ := w.gen.Authmethod()
	authProvider, _ := w.gen.Authprovider()

	details := map[string]any{
		"authid":       authID,
		"authrole":     authRole,
		"authmethod":   authMethod,
		"authprovider": authProvider,
	}

	genRoles, err := w.gen.Roles()
	if err != nil {
		return details
	}

	roles := map[string]any{}
	if r, err := genRoles.Dealer(); err == nil {
		dealer := map[string]any{}
		if r.CallerIdentification() {
			dealer["caller_identification"] = true
		}
		if r.CallTimeout() {
			dealer["call_timeout"] = true
		}
		if r.CallCanceling() {
			dealer["call_canceling"] = true
		}
		if r.ProgressiveCallResults() {
			dealer["progressive_call_results"] = true
		}
		if r.PatternBasedRegistration() {
			dealer["pattern_based_registration"] = true
		}
		if r.SharedRegistration() {
			dealer["shared_registration"] = true
		}
		roles["dealer"] = dealer
	}

	if r, err := genRoles.Broker(); err == nil {
		broker := map[string]any{}
		if r.PublisherIdentification() {
			broker["publisher_identification"] = true
		}
		if r.PublisherExclusion() {
			broker["publisher_exclusion"] = true
		}
		if r.AcknowledgeEventReceived() {
			broker["acknowledge_event_received"] = true
		}
		if r.PatternBasedSubscription() {
			broker["pattern_based_subscription"] = true
		}
		roles["broker"] = broker
	}

	details["roles"] = roles

	return details
}

func WelcomeToCapnproto(w *messages.Welcome) ([]byte, error) {
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		return nil, err
	}

	welcome, err := gen.NewRootWelcome(seg)
	if err != nil {
		return nil, err
	}

	welcome.SetSessionID(w.SessionID())

	authid, ok := w.Details()["authid"].(string)
	if ok {
		if err := welcome.SetAuthid(authid); err != nil {
			return nil, err
		}
	}

	authrole, ok := w.Details()["authrole"].(string)
	if ok {
		if err := welcome.SetAuthrole(authrole); err != nil {
			return nil, err
		}
	}

	authmethod, ok := w.Details()["authmethod"].(string)
	if ok {
		if err := welcome.SetAuthmethod(authmethod); err != nil {
			return nil, err
		}
	}

	authprovider, ok := w.Details()["authprovider"].(string)
	if ok {
		if err := welcome.SetAuthprovider(authprovider); err != nil {
			return nil, err
		}
	}

	roles, err := welcome.NewRoles()
	if err != nil {
		return nil, err
	}

	rolesMap, ok := w.Details()["roles"].(map[string]any)
	if ok {
		if dealerMap, ok := rolesMap["dealer"].(map[string]any); ok {
			dealer, err := roles.NewDealer()
			if err == nil {
				if v, ok := dealerMap["caller_identification"].(bool); ok {
					dealer.SetCallerIdentification(v)
				}
				if v, ok := dealerMap["call_timeout"].(bool); ok {
					dealer.SetCallTimeout(v)
				}
				if v, ok := dealerMap["call_canceling"].(bool); ok {
					dealer.SetCallCanceling(v)
				}
				if v, ok := dealerMap["progressive_call_results"].(bool); ok {
					dealer.SetProgressiveCallResults(v)
				}
				if v, ok := dealerMap["pattern_based_registration"].(bool); ok {
					dealer.SetPatternBasedRegistration(v)
				}
				_ = roles.SetDealer(dealer)
			}
		}

		if brokerMap, ok := rolesMap["broker"].(map[string]any); ok {
			broker, err := roles.NewBroker()
			if err == nil {
				if v, ok := brokerMap["publisher_identification"].(bool); ok {
					broker.SetPublisherIdentification(v)
				}
				if v, ok := brokerMap["publisher_exclusion"].(bool); ok {
					broker.SetPublisherExclusion(v)
				}
				if v, ok := brokerMap["acknowledge_event_received"].(bool); ok {
					broker.SetAcknowledgeEventReceived(v)
				}
				if v, ok := brokerMap["pattern_based_subscription"].(bool); ok {
					broker.SetPatternBasedSubscription(v)
				}
				_ = roles.SetBroker(broker)
			}
		}
	}

	if err := welcome.SetRoles(roles); err != nil {
		return nil, err
	}

	data, err := msg.MarshalPacked()
	if err != nil {
		return nil, err
	}

	return PrependHeader(messages.MessageTypeWelcome, data, nil), nil
}

func CapnprotoToWelcome(data []byte) (*messages.Welcome, error) {
	msg, err := capnp.UnmarshalPacked(data)
	if err != nil {
		return nil, err
	}

	welcome, err := gen.ReadRootWelcome(msg)
	if err != nil {
		return nil, err
	}

	return messages.NewWelcomeWithFields(NewWelcomeFields(&welcome)), nil
}
