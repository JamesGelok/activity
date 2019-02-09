package streams

import (
	"context"
	"errors"
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"strings"
)

// JSONResolver resolves a JSON-deserialized map into its concrete ActivityStreams
// type
type JSONResolver struct {
	callbacks []interface{}
}

// NewJSONResolver creates a new Resolver that takes a JSON-deserialized generic
// map and determines the correct concrete Go type. The callback function is
// guaranteed to receive a value whose underlying ActivityStreams type matches
// the concrete interface name in its signature. The callback functions must
// be of the form:
//
//   func(context.Context, <TypeInterface>) error
//
// where TypeInterface is the code-generated interface for an ActivityStream
// type. An error is returned if a callback function does not match this
// signature.
func NewJSONResolver(callbacks ...interface{}) (*JSONResolver, error) {
	for _, cb := range callbacks {
		// Each callback function must satisfy one known function signature, or else we will generate a runtime error instead of silently fail.
		switch cb.(type) {
		case func(context.Context, vocab.ActivityStreamsAccept) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsActivity) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsAdd) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsAnnounce) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsApplication) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsArrive) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsArticle) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsAudio) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsBlock) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsCollection) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsCollectionPage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsCreate) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsDelete) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsDislike) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsDocument) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsEvent) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsFlag) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsFollow) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsGroup) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsIgnore) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsImage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsIntransitiveActivity) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsInvite) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsJoin) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsLeave) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsLike) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsLink) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsListen) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsMention) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsMove) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsNote) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsObject) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOffer) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOrderedCollection) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsOrganization) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPage) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPerson) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsPlace) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsProfile) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsQuestion) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsRead) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsReject) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsRelationship) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsRemove) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsService) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTentativeAccept) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTentativeReject) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTombstone) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsTravel) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsUndo) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsUpdate) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsVideo) error:
			// Do nothing, this callback has a correct signature.
		case func(context.Context, vocab.ActivityStreamsView) error:
			// Do nothing, this callback has a correct signature.
		default:
			return nil, errors.New("a callback function is of the wrong signature and would never be called")
		}
	}
	return &JSONResolver{callbacks: callbacks}, nil
}

// toAliasMap converts a JSONLD context into a map of vocabulary name to alias.
func toAliasMap(i interface{}) (m map[string]string) {
	m = make(map[string]string)
	toHttpHttpsFn := func(s string) (ok bool, http, https string) {
		if strings.HasPrefix(s, "http://") {
			ok = true
			http = s
			https = "https" + strings.TrimPrefix(s, "http")
		} else if strings.HasPrefix(s, "https://") {
			ok = true
			https = s
			http = "http" + strings.TrimPrefix(s, "https")
		}
		return
	}
	switch v := i.(type) {
	case string:
		// Single entry, no alias.

		if ok, http, https := toHttpHttpsFn(v); ok {
			m[http] = ""
			m[https] = ""
		} else {
			m[v] = ""
		}
	case []interface{}:
		// Recursively apply.

		for _, elem := range v {
			r := toAliasMap(elem)
			for k, val := range r {
				m[k] = val
			}
		}
	case map[string]interface{}:
		// Map any aliases.

		for k, val := range v {
			// Only handle string aliases.
			switch conc := val.(type) {
			case string:
				m[k] = conc
			}
		}
	}
	return
}

// Resolve determines the ActivityStreams type of the payload, then applies the
// first callback function whose signature accepts the ActivityStreams value's
// type. This strictly assures that the callback function will only be passed
// ActivityStream objects whose type matches its interface. Returns an error
// if the ActivityStreams type does not match callbackers or is not a type
// handled by the generated code. If multiple types are present, it will check
// each one in order and apply only the first one. It returns an unhandled
// error for a multi-typed object if none of the types were able to be handled.
func (this JSONResolver) Resolve(ctx context.Context, m map[string]interface{}) error {
	typeValue, ok := m["type"]
	if !ok {
		return fmt.Errorf("cannot determine ActivityStreams type: 'type' property is missing")
	}
	rawContext, ok := m["@context"]
	if !ok {
		return fmt.Errorf("cannot determine ActivityStreams type: '@context' is missing")
	}
	aliasMap := toAliasMap(rawContext)
	// Begin: Private lambda to handle a single string "type" value. Makes code generation easier.
	handleFn := func(typeString string) error {
		ActivityStreamsAlias, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]
		if !ok {
			ActivityStreamsAlias, _ = aliasMap["http://www.w3.org/TR/activitystreams-vocabulary"]
		}
		if len(ActivityStreamsAlias) > 0 {
			ActivityStreamsAlias += ":"
		}

		if typeString == ActivityStreamsAlias+"Accept" {
			v, err := mgr.DeserializeAcceptActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAccept) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Activity" {
			v, err := mgr.DeserializeActivityActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsActivity) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Add" {
			v, err := mgr.DeserializeAddActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAdd) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Announce" {
			v, err := mgr.DeserializeAnnounceActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAnnounce) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Application" {
			v, err := mgr.DeserializeApplicationActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsApplication) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Arrive" {
			v, err := mgr.DeserializeArriveActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsArrive) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Article" {
			v, err := mgr.DeserializeArticleActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsArticle) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Audio" {
			v, err := mgr.DeserializeAudioActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsAudio) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Block" {
			v, err := mgr.DeserializeBlockActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsBlock) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Collection" {
			v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsCollection) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"CollectionPage" {
			v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsCollectionPage) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Create" {
			v, err := mgr.DeserializeCreateActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsCreate) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Delete" {
			v, err := mgr.DeserializeDeleteActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsDelete) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Dislike" {
			v, err := mgr.DeserializeDislikeActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsDislike) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Document" {
			v, err := mgr.DeserializeDocumentActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsDocument) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Event" {
			v, err := mgr.DeserializeEventActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsEvent) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Flag" {
			v, err := mgr.DeserializeFlagActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsFlag) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Follow" {
			v, err := mgr.DeserializeFollowActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsFollow) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Group" {
			v, err := mgr.DeserializeGroupActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsGroup) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Ignore" {
			v, err := mgr.DeserializeIgnoreActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsIgnore) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Image" {
			v, err := mgr.DeserializeImageActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsImage) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"IntransitiveActivity" {
			v, err := mgr.DeserializeIntransitiveActivityActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsIntransitiveActivity) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Invite" {
			v, err := mgr.DeserializeInviteActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsInvite) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Join" {
			v, err := mgr.DeserializeJoinActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsJoin) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Leave" {
			v, err := mgr.DeserializeLeaveActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsLeave) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Like" {
			v, err := mgr.DeserializeLikeActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsLike) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Link" {
			v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsLink) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Listen" {
			v, err := mgr.DeserializeListenActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsListen) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Mention" {
			v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsMention) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Move" {
			v, err := mgr.DeserializeMoveActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsMove) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Note" {
			v, err := mgr.DeserializeNoteActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsNote) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Object" {
			v, err := mgr.DeserializeObjectActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsObject) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Offer" {
			v, err := mgr.DeserializeOfferActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOffer) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"OrderedCollection" {
			v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOrderedCollection) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"OrderedCollectionPage" {
			v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOrderedCollectionPage) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Organization" {
			v, err := mgr.DeserializeOrganizationActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsOrganization) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Page" {
			v, err := mgr.DeserializePageActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPage) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Person" {
			v, err := mgr.DeserializePersonActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPerson) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Place" {
			v, err := mgr.DeserializePlaceActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsPlace) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Profile" {
			v, err := mgr.DeserializeProfileActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsProfile) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Question" {
			v, err := mgr.DeserializeQuestionActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsQuestion) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Read" {
			v, err := mgr.DeserializeReadActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsRead) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Reject" {
			v, err := mgr.DeserializeRejectActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsReject) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Relationship" {
			v, err := mgr.DeserializeRelationshipActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsRelationship) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Remove" {
			v, err := mgr.DeserializeRemoveActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsRemove) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Service" {
			v, err := mgr.DeserializeServiceActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsService) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"TentativeAccept" {
			v, err := mgr.DeserializeTentativeAcceptActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTentativeAccept) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"TentativeReject" {
			v, err := mgr.DeserializeTentativeRejectActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTentativeReject) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Tombstone" {
			v, err := mgr.DeserializeTombstoneActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTombstone) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Travel" {
			v, err := mgr.DeserializeTravelActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsTravel) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Undo" {
			v, err := mgr.DeserializeUndoActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsUndo) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Update" {
			v, err := mgr.DeserializeUpdateActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsUpdate) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"Video" {
			v, err := mgr.DeserializeVideoActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsVideo) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else if typeString == ActivityStreamsAlias+"View" {
			v, err := mgr.DeserializeViewActivityStreams()(m, aliasMap)
			if err != nil {
				return err
			}
			for _, i := range this.callbacks {
				if fn, ok := i.(func(context.Context, vocab.ActivityStreamsView) error); ok {
					return fn(ctx, v)
				}
			}
			return ErrNoCallbackMatch
		} else {
			return ErrUnhandledType
		}
	}
	// End: Private lambda
	if typeStr, ok := typeValue.(string); ok {
		return handleFn(typeStr)
	} else if typeIArr, ok := typeValue.([]interface{}); ok {
		for _, typeI := range typeIArr {
			if typeStr, ok := typeI.(string); ok {
				if err := handleFn(typeStr); err == nil {
					return nil
				} else if err == ErrUnhandledType {
					// Keep trying other types: only if all fail do we return this error.
					continue
				} else {
					return err
				}
			}
		}
		return ErrUnhandledType
	} else {
		return ErrUnhandledType
	}
}
