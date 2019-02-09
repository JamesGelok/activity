package typeblock

import vocab "github.com/go-fed/activity/streams/vocab"

var mgr privateManager

var typePropertyConstructor func() vocab.ActivityStreamsTypeProperty

// privateManager abstracts the code-generated manager that provides access to
// concrete implementations.
type privateManager interface {
	// DeserializeActorPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsActorProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeActorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsActorProperty, error)
	// DeserializeAltitudePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsAltitudeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeAltitudePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAltitudeProperty, error)
	// DeserializeAttachmentPropertyActivityStreams returns the
	// deserialization method for the "ActivityStreamsAttachmentProperty"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeAttachmentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAttachmentProperty, error)
	// DeserializeAttributedToPropertyActivityStreams returns the
	// deserialization method for the
	// "ActivityStreamsAttributedToProperty" non-functional property in
	// the vocabulary "ActivityStreams"
	DeserializeAttributedToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAttributedToProperty, error)
	// DeserializeAudiencePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsAudienceProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeAudiencePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsAudienceProperty, error)
	// DeserializeBccPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsBccProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeBccPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBccProperty, error)
	// DeserializeBtoPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsBtoProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeBtoPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsBtoProperty, error)
	// DeserializeCcPropertyActivityStreams returns the deserialization method
	// for the "ActivityStreamsCcProperty" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeCcPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsCcProperty, error)
	// DeserializeContentPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsContentProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeContentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsContentProperty, error)
	// DeserializeContextPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsContextProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeContextPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsContextProperty, error)
	// DeserializeDurationPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsDurationProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeDurationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsDurationProperty, error)
	// DeserializeEndTimePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsEndTimeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeEndTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsEndTimeProperty, error)
	// DeserializeGeneratorPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsGeneratorProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeGeneratorPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsGeneratorProperty, error)
	// DeserializeIconPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsIconProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeIconPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIconProperty, error)
	// DeserializeIdPropertyActivityStreams returns the deserialization method
	// for the "ActivityStreamsIdProperty" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeIdPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsIdProperty, error)
	// DeserializeImagePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsImageProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeImagePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsImageProperty, error)
	// DeserializeInReplyToPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsInReplyToProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeInReplyToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInReplyToProperty, error)
	// DeserializeInstrumentPropertyActivityStreams returns the
	// deserialization method for the "ActivityStreamsInstrumentProperty"
	// non-functional property in the vocabulary "ActivityStreams"
	DeserializeInstrumentPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsInstrumentProperty, error)
	// DeserializeLikesPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsLikesProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeLikesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLikesProperty, error)
	// DeserializeLocationPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsLocationProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeLocationPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsLocationProperty, error)
	// DeserializeMediaTypePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsMediaTypeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeMediaTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsMediaTypeProperty, error)
	// DeserializeNamePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsNameProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeNamePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsNameProperty, error)
	// DeserializeObjectPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsObjectProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeObjectPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsObjectProperty, error)
	// DeserializeOriginPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsOriginProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeOriginPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsOriginProperty, error)
	// DeserializePreviewPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsPreviewProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializePreviewPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPreviewProperty, error)
	// DeserializePublishedPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsPublishedProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializePublishedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsPublishedProperty, error)
	// DeserializeRepliesPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsRepliesProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeRepliesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsRepliesProperty, error)
	// DeserializeResultPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsResultProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeResultPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsResultProperty, error)
	// DeserializeSharesPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsSharesProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeSharesPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSharesProperty, error)
	// DeserializeStartTimePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsStartTimeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeStartTimePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsStartTimeProperty, error)
	// DeserializeSummaryPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsSummaryProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeSummaryPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsSummaryProperty, error)
	// DeserializeTagPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsTagProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeTagPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTagProperty, error)
	// DeserializeTargetPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsTargetProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeTargetPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTargetProperty, error)
	// DeserializeToPropertyActivityStreams returns the deserialization method
	// for the "ActivityStreamsToProperty" non-functional property in the
	// vocabulary "ActivityStreams"
	DeserializeToPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsToProperty, error)
	// DeserializeTypePropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsTypeProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeTypePropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsTypeProperty, error)
	// DeserializeUpdatedPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsUpdatedProperty" non-functional
	// property in the vocabulary "ActivityStreams"
	DeserializeUpdatedPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUpdatedProperty, error)
	// DeserializeUrlPropertyActivityStreams returns the deserialization
	// method for the "ActivityStreamsUrlProperty" non-functional property
	// in the vocabulary "ActivityStreams"
	DeserializeUrlPropertyActivityStreams() func(map[string]interface{}, map[string]string) (vocab.ActivityStreamsUrlProperty, error)
}

// jsonldContexter is a private interface to determine the JSON-LD contexts and
// aliases needed for functional and non-functional properties. It is a helper
// interface for this implementation.
type jsonldContexter interface {
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
}

// SetManager sets the manager package-global variable. For internal use only, do
// not use as part of Application behavior. Must be called at golang init time.
func SetManager(m privateManager) {
	mgr = m
}

// SetTypePropertyConstructor sets the "type" property's constructor in the
// package-global variable. For internal use only, do not use as part of
// Application behavior. Must be called at golang init time. Permits
// ActivityStreams types to correctly set their "type" property at
// construction time, so users don't have to remember to do so each time. It
// is dependency injected so other go-fed compatible implementations could
// inject their own type.
func SetTypePropertyConstructor(f func() vocab.ActivityStreamsTypeProperty) {
	typePropertyConstructor = f
}
