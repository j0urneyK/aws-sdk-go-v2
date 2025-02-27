// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Contains information about the contexts that a user is using in a session. You
// can configure Amazon Lex V2 to set a context when an intent is fulfilled, or you
// can set a context using the , , or operations. Use a context to indicate to
// Amazon Lex V2 intents that should be used as follow-up intents. For example, if
// the active context is order-fulfilled, only intents that have order-fulfilled
// configured as a trigger are considered for follow up.
type ActiveContext struct {

	// A list of contexts active for the request. A context can be activated when a
	// previous intent is fulfilled, or by including the context in the request. If you
	// don't specify a list of contexts, Amazon Lex V2 will use the current list of
	// contexts for the session. If you specify an empty list, all contexts for the
	// session are cleared.
	//
	// This member is required.
	ContextAttributes map[string]string

	// The name of the context.
	//
	// This member is required.
	Name *string

	// Indicates the number of turns or seconds that the context is active. Once the
	// time to live expires, the context is no longer returned in a response.
	//
	// This member is required.
	TimeToLive *ActiveContextTimeToLive

	noSmithyDocumentSerde
}

// The time that a context is active. You can specify the time to live in seconds
// or in conversation turns.
type ActiveContextTimeToLive struct {

	// The number of seconds that the context is active. You can specify between 5 and
	// 86400 seconds (24 hours).
	//
	// This member is required.
	TimeToLiveInSeconds *int32

	// The number of turns that the context is active. You can specify up to 20 turns.
	// Each request and response from the bot is a turn.
	//
	// This member is required.
	TurnsToLive *int32

	noSmithyDocumentSerde
}

// Represents a chunk of audio sent from the client application to Amazon Lex V2.
// The audio is all or part of an utterance from the user. Amazon Lex V2
// accumulates audio chunks until it recognizes a natural pause in speech before
// processing the input.
type AudioInputEvent struct {

	// The encoding used for the audio chunk. You must use 8 KHz PCM 16-bit
	// mono-channel little-endian format. The value of the field should be: audio/lpcm;
	// sample-rate=8000; sample-size-bits=16; channel-count=1; is-big-endian=false
	//
	// This member is required.
	ContentType *string

	// An encoded stream of audio.
	AudioChunk []byte

	// A timestamp set by the client of the date and time that the event was sent to
	// Amazon Lex V2.
	ClientTimestampMillis int64

	// A unique identifier that your application assigns to the event. You can use this
	// to identify events in logs.
	EventId *string

	noSmithyDocumentSerde
}

// An event sent from Amazon Lex V2 to your client application containing audio to
// play to the user.
type AudioResponseEvent struct {

	// A chunk of the audio to play.
	AudioChunk []byte

	// The encoding of the audio chunk. This is the same as the encoding configure in
	// the contentType field of the ConfigurationEvent.
	ContentType *string

	// A unique identifier of the event sent by Amazon Lex V2. The identifier is in the
	// form RESPONSE-N, where N is a number starting with one and incremented for each
	// event sent by Amazon Lex V2 in the current session.
	EventId *string

	noSmithyDocumentSerde
}

// A button that appears on a response card show to the user.
type Button struct {

	// The text that is displayed on the button.
	//
	// This member is required.
	Text *string

	// The value returned to Amazon Lex V2 when a user chooses the button.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Provides a score that indicates the confidence that Amazon Lex V2 has that an
// intent is the one that satisfies the user's intent.
type ConfidenceScore struct {

	// A score that indicates how confident Amazon Lex V2 is that an intent satisfies
	// the user's intent. Ranges between 0.00 and 1.00. Higher scores indicate higher
	// confidence.
	Score float64

	noSmithyDocumentSerde
}

// The initial event sent from the application to Amazon Lex V2 to configure the
// conversation, including session and request attributes and the response content
// type.
type ConfigurationEvent struct {

	// The message that Amazon Lex V2 returns in the response can be either text or
	// speech based on the responseContentType value.
	//
	// * If the value is
	// text/plain;charset=utf-8, Amazon Lex V2 returns text in the response.
	//
	// * If the
	// value begins with audio/, Amazon Lex V2 returns speech in the response. Amazon
	// Lex V2 uses Amazon Polly to generate the speech using the configuration that you
	// specified in the requestContentType parameter. For example, if you specify
	// audio/mpeg as the value, Amazon Lex V2 returns speech in the MPEG format.
	//
	// * If
	// the value is audio/pcm, the speech returned is audio/pcm in 16-bit,
	// little-endian format.
	//
	// * The following are the accepted values:
	//
	// * audio/mpeg
	//
	// *
	// audio/ogg
	//
	// * audio/pcm
	//
	// * audio/* (defaults to mpeg)
	//
	// * text/plain;
	// charset=utf-8
	//
	// This member is required.
	ResponseContentType *string

	// A timestamp set by the client of the date and time that the event was sent to
	// Amazon Lex V2.
	ClientTimestampMillis int64

	// Determines whether Amazon Lex V2 should send audio responses to the client
	// application. Set this field to false when the client is operating in a playback
	// mode where audio responses are played to the user. If the client isn't operating
	// in playback mode, such as a text chat application, set this to true so that
	// Amazon Lex V2 doesn't wait for the prompt to finish playing on the client.
	DisablePlayback bool

	// A unique identifier that your application assigns to the event. You can use this
	// to identify events in logs.
	EventId *string

	// Request-specific information passed between the client application and Amazon
	// Lex V2. The namespace x-amz-lex: is reserved for special attributes. Don't
	// create any request attributes for prefix x-amz-lex:.
	RequestAttributes map[string]string

	// The state of the user's session with Amazon Lex V2.
	SessionState *SessionState

	// A list of messages to send to the user. If you set the welcomeMessage field, you
	// must also set the DialogAction
	// (https://docs.aws.amazon.com/lexv2/latest/dg/API_runtime_DialogAction.html)
	// structure's type
	// (https://docs.aws.amazon.com/lexv2/latest/dg/API_runtime_DialogAction.html#lexv2-Type-runtime_DialogAction-type)
	// field.
	WelcomeMessages []Message

	noSmithyDocumentSerde
}

// The next action that Amazon Lex V2 should take.
type DialogAction struct {

	// The next action that the bot should take in its interaction with the user. The
	// possible values are:
	//
	// * Close - Indicates that there will not be a response from
	// the user. For example, the statement "Your order has been placed" does not
	// require a response.
	//
	// * ConfirmIntent - The next action is asking the user if the
	// intent is complete and ready to be fulfilled. This is a yes/no question such as
	// "Place the order?"
	//
	// * Delegate - The next action is determined by Amazon Lex
	// V2.
	//
	// * ElicitIntent - The next action is to elicit an intent from the user.
	//
	// *
	// ElicitSlot - The next action is to elicit a slot value from the user.
	//
	// This member is required.
	Type DialogActionType

	// Configures the slot to use spell-by-letter or spell-by-word style. When you use
	// a style on a slot, users can spell out their input to make it clear to your
	// bot.
	//
	// * Spell by letter - "b" "o" "b"
	//
	// * Spell by word - "b as in boy" "o as in
	// oscar" "b as in boy"
	//
	// For more information, see  Using spelling to enter slot
	// values  (https://docs.aws.amazon.com/lexv2/latest/dg/using-spelling.html).
	SlotElicitationStyle StyleType

	// The name of the slot that should be elicited from the user.
	SlotToElicit *string

	// The name of the constituent sub slot of the composite slot specified in
	// slotToElicit that should be elicited from the user.
	SubSlotToElicit *ElicitSubSlot

	noSmithyDocumentSerde
}

// A notification from the client that it is disconnecting from Amazon Lex V2.
// Sending a DisconnectionEvent event is optional, but can help identify a
// conversation in logs.
type DisconnectionEvent struct {

	// A timestamp set by the client of the date and time that the event was sent to
	// Amazon Lex V2.
	ClientTimestampMillis int64

	// A unique identifier that your application assigns to the event. You can use this
	// to identify events in logs.
	EventId *string

	noSmithyDocumentSerde
}

// A DTMF character sent from the client application. DTMF characters are typically
// sent from a phone keypad to represent numbers. For example, you can have Amazon
// Lex V2 process a credit card number input from a phone.
type DTMFInputEvent struct {

	// The DTMF character that the user pressed. The allowed characters are A - D, 0 -
	// 9, # and *.
	//
	// This member is required.
	InputCharacter *string

	// A timestamp set by the client of the date and time that the event was sent to
	// Amazon Lex V2.
	ClientTimestampMillis int64

	// A unique identifier that your application assigns to the event. You can use this
	// to identify events in logs.
	EventId *string

	noSmithyDocumentSerde
}

// The specific constituent sub slot of the composite slot to elicit in dialog
// action.
type ElicitSubSlot struct {

	// The name of the slot that should be elicited from the user.
	//
	// This member is required.
	Name *string

	// The field is not supported.
	SubSlotToElicit *ElicitSubSlot

	noSmithyDocumentSerde
}

// Event that Amazon Lex V2 sends to indicate that the stream is still open between
// the client application and Amazon Lex V2
type HeartbeatEvent struct {

	// A unique identifier of the event sent by Amazon Lex V2. The identifier is in the
	// form RESPONSE-N, where N is a number starting with one and incremented for each
	// event sent by Amazon Lex V2 in the current session.
	EventId *string

	noSmithyDocumentSerde
}

// A card that is shown to the user by a messaging platform. You define the
// contents of the card, the card is displayed by the platform. When you use a
// response card, the response from the user is constrained to the text associated
// with a button on the card.
type ImageResponseCard struct {

	// The title to display on the response card. The format of the title is determined
	// by the platform displaying the response card.
	//
	// This member is required.
	Title *string

	// A list of buttons that should be displayed on the response card. The arrangement
	// of the buttons is determined by the platform that displays the button.
	Buttons []Button

	// The URL of an image to display on the response card. The image URL must be
	// publicly available so that the platform displaying the response card has access
	// to the image.
	ImageUrl *string

	// The subtitle to display on the response card. The format of the subtitle is
	// determined by the platform displaying the response card.
	Subtitle *string

	noSmithyDocumentSerde
}

// The current intent that Amazon Lex V2 is attempting to fulfill.
type Intent struct {

	// The name of the intent.
	//
	// This member is required.
	Name *string

	// Contains information about whether fulfillment of the intent has been confirmed.
	ConfirmationState ConfirmationState

	// A map of all of the slots for the intent. The name of the slot maps to the value
	// of the slot. If a slot has not been filled, the value is null.
	Slots map[string]Slot

	// Contains fulfillment information for the intent.
	State IntentState

	noSmithyDocumentSerde
}

// Contains the current state of the conversation between the client application
// and Amazon Lex V2.
type IntentResultEvent struct {

	// A unique identifier of the event sent by Amazon Lex V2. The identifier is in the
	// form RESPONSE-N, where N is a number starting with one and incremented for each
	// event sent by Amazon Lex V2 in the current session.
	EventId *string

	// Indicates whether the input to the operation was text or speech.
	InputMode InputMode

	// A list of intents that Amazon Lex V2 determined might satisfy the user's
	// utterance. Each interpretation includes the intent, a score that indicates how
	// confident Amazon Lex V2 is that the interpretation is the correct one, and an
	// optional sentiment response that indicates the sentiment expressed in the
	// utterance.
	Interpretations []Interpretation

	// The bot member that is processing the intent.
	RecognizedBotMember *RecognizedBotMember

	// The attributes sent in the request.
	RequestAttributes map[string]string

	// The identifier of the session in use.
	SessionId *string

	// The state of the user's session with Amazon Lex V2.
	SessionState *SessionState

	noSmithyDocumentSerde
}

// An intent that Amazon Lex V2 determined might satisfy the user's utterance. The
// intents are ordered by the confidence score.
type Interpretation struct {

	// A list of intents that might satisfy the user's utterance. The intents are
	// ordered by the confidence score.
	Intent *Intent

	// Determines the threshold where Amazon Lex V2 will insert the
	// AMAZON.FallbackIntent, AMAZON.KendraSearchIntent, or both when returning
	// alternative intents in a response. AMAZON.FallbackIntent and
	// AMAZON.KendraSearchIntent are only inserted if they are configured for the bot.
	NluConfidence *ConfidenceScore

	// The sentiment expressed in an utterance. When the bot is configured to send
	// utterances to Amazon Comprehend for sentiment analysis, this field contains the
	// result of the analysis.
	SentimentResponse *SentimentResponse

	noSmithyDocumentSerde
}

// Container for text that is returned to the customer..
type Message struct {

	// Indicates the type of response.
	//
	// This member is required.
	ContentType MessageContentType

	// The text of the message.
	Content *string

	// A card that is shown to the user by a messaging platform. You define the
	// contents of the card, the card is displayed by the platform. When you use a
	// response card, the response from the user is constrained to the text associated
	// with a button on the card.
	ImageResponseCard *ImageResponseCard

	noSmithyDocumentSerde
}

// Event sent from the client application to Amazon Lex V2 to indicate that
// playback of audio is complete and that Amazon Lex V2 should start processing the
// user's input.
type PlaybackCompletionEvent struct {

	// A timestamp set by the client of the date and time that the event was sent to
	// Amazon Lex V2.
	ClientTimestampMillis int64

	// A unique identifier that your application assigns to the event. You can use this
	// to identify events in logs.
	EventId *string

	noSmithyDocumentSerde
}

// Event sent from Amazon Lex V2 to indicate to the client application should stop
// playback of audio. For example, if the client is playing a prompt that asks for
// the user's telephone number, the user might start to say the phone number before
// the prompt is complete. Amazon Lex V2 sends this event to the client application
// to indicate that the user is responding and that Amazon Lex V2 is processing
// their input.
type PlaybackInterruptionEvent struct {

	// The identifier of the event that contained the audio, DTMF, or text that caused
	// the interruption.
	CausedByEventId *string

	// A unique identifier of the event sent by Amazon Lex V2. The identifier is in the
	// form RESPONSE-N, where N is a number starting with one and incremented for each
	// event sent by Amazon Lex V2 in the current session.
	EventId *string

	// Indicates the type of user input that Amazon Lex V2 detected.
	EventReason PlaybackInterruptionReason

	noSmithyDocumentSerde
}

// The bot member that processes the request.
type RecognizedBotMember struct {

	// The identifier of the bot member that processes the request.
	//
	// This member is required.
	BotId *string

	// The name of the bot member that processes the request.
	BotName *string

	noSmithyDocumentSerde
}

// Provides an array of phrases that should be given preference when resolving
// values for a slot.
type RuntimeHintDetails struct {

	// One or more strings that Amazon Lex V2 should look for in the input to the bot.
	// Each phrase is given preference when deciding on slot values.
	RuntimeHintValues []RuntimeHintValue

	// A map of constituent sub slot names inside a composite slot in the intent and
	// the phrases that should be added for each sub slot. Inside each composite slot
	// hints, this structure provides a mechanism to add granular sub slot phrases.
	// Only sub slot hints are supported for composite slots. The intent name,
	// composite slot name and the constituent sub slot names must exist.
	SubSlotHints map[string]RuntimeHintDetails

	noSmithyDocumentSerde
}

// You can provide Amazon Lex V2 with hints to the phrases that a customer is
// likely to use for a slot. When a slot with hints is resolved, the phrases in the
// runtime hints are preferred in the resolution. You can provide hints for a
// maximum of 100 intents. You can provide a maximum of 100 slots. Before you can
// use runtime hints with an existing bot, you must first rebuild the bot. For more
// information, see Using runtime hints to improve recognition of slot values
// (https://docs.aws.amazon.com/lexv2/latest/dg/using-hints.html).
type RuntimeHints struct {

	// A list of the slots in the intent that should have runtime hints added, and the
	// phrases that should be added for each slot. The first level of the slotHints map
	// is the name of the intent. The second level is the name of the slot within the
	// intent. For more information, see Using hints to improve accuracy
	// (https://docs.aws.amazon.com/lexv2/latest/dg/using-hints.html). The intent name
	// and slot name must exist.
	SlotHints map[string]map[string]RuntimeHintDetails

	noSmithyDocumentSerde
}

// Provides the phrase that Amazon Lex V2 should look for in the user's input to
// the bot.
type RuntimeHintValue struct {

	// The phrase that Amazon Lex V2 should look for in the user's input to the bot.
	//
	// This member is required.
	Phrase *string

	noSmithyDocumentSerde
}

// Provides information about the sentiment expressed in a user's response in a
// conversation. Sentiments are determined using Amazon Comprehend. Sentiments are
// only returned if they are enabled for the bot. For more information, see
// Determine Sentiment
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-sentiment.html) in the
// Amazon Comprehend developer guide.
type SentimentResponse struct {

	// The overall sentiment expressed in the user's response. This is the sentiment
	// most likely expressed by the user based on the analysis by Amazon Comprehend.
	Sentiment SentimentType

	// The individual sentiment responses for the utterance.
	SentimentScore *SentimentScore

	noSmithyDocumentSerde
}

// The individual sentiment responses for the utterance.
type SentimentScore struct {

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the MIXED sentiment.
	Mixed float64

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the NEGATIVE sentiment.
	Negative float64

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the NEUTRAL sentiment.
	Neutral float64

	// The level of confidence that Amazon Comprehend has in the accuracy of its
	// detection of the POSITIVE sentiment.
	Positive float64

	noSmithyDocumentSerde
}

// The state of the user's session with Amazon Lex V2.
type SessionState struct {

	// One or more contexts that indicate to Amazon Lex V2 the context of a request.
	// When a context is active, Amazon Lex V2 considers intents with the matching
	// context as a trigger as the next intent in a session.
	ActiveContexts []ActiveContext

	// The next step that Amazon Lex V2 should take in the conversation with a user.
	DialogAction *DialogAction

	// The active intent that Amazon Lex V2 is processing.
	Intent *Intent

	// A unique identifier for a specific request.
	OriginatingRequestId *string

	// Hints for phrases that a customer is likely to use for a slot. Amazon Lex V2
	// uses the hints to help determine the correct value of a slot.
	RuntimeHints *RuntimeHints

	// Map of key/value pairs representing session-specific context information. It
	// contains application information passed between Amazon Lex V2 and a client
	// application.
	SessionAttributes map[string]string

	noSmithyDocumentSerde
}

// A value that Amazon Lex V2 uses to fulfill an intent.
type Slot struct {

	// When the shape value is List, it indicates that the values field contains a list
	// of slot values. When the value is Scalar, it indicates that the value field
	// contains a single value.
	Shape Shape

	// The constituent sub slots of a composite slot.
	SubSlots map[string]Slot

	// The current value of the slot.
	Value *Value

	// A list of one or more values that the user provided for the slot. For example,
	// if a for a slot that elicits pizza toppings, the values might be "pepperoni" and
	// "pineapple."
	Values []Slot

	noSmithyDocumentSerde
}

// Represents a stream of events between your application and Amazon Lex V2.
//
// The following types satisfy this interface:
//
//	StartConversationRequestEventStreamMemberAudioInputEvent
//	StartConversationRequestEventStreamMemberConfigurationEvent
//	StartConversationRequestEventStreamMemberDisconnectionEvent
//	StartConversationRequestEventStreamMemberDTMFInputEvent
//	StartConversationRequestEventStreamMemberPlaybackCompletionEvent
//	StartConversationRequestEventStreamMemberTextInputEvent
type StartConversationRequestEventStream interface {
	isStartConversationRequestEventStream()
}

// Speech audio sent from your client application to Amazon Lex V2. Audio starts
// accumulating when Amazon Lex V2 identifies a voice and continues until a natural
// pause in the speech is found before processing.
type StartConversationRequestEventStreamMemberAudioInputEvent struct {
	Value AudioInputEvent

	noSmithyDocumentSerde
}

func (*StartConversationRequestEventStreamMemberAudioInputEvent) isStartConversationRequestEventStream() {
}

// Configuration information sent from your client application to Amazon Lex V2
type StartConversationRequestEventStreamMemberConfigurationEvent struct {
	Value ConfigurationEvent

	noSmithyDocumentSerde
}

func (*StartConversationRequestEventStreamMemberConfigurationEvent) isStartConversationRequestEventStream() {
}

// Event sent from the client application to indicate to Amazon Lex V2 that the
// conversation is over.
type StartConversationRequestEventStreamMemberDisconnectionEvent struct {
	Value DisconnectionEvent

	noSmithyDocumentSerde
}

func (*StartConversationRequestEventStreamMemberDisconnectionEvent) isStartConversationRequestEventStream() {
}

// DTMF information sent to Amazon Lex V2 by your application. Amazon Lex V2
// accumulates the DMTF information from when the user sends the first character
// and ends
//
// * when there's a pause longer that the value configured for the end
// timeout.
//
// * when there's a digit that is the configured end character.
//
// * when
// Amazon Lex V2 accumulates characters equal to the maximum DTMF character
// configuration.
type StartConversationRequestEventStreamMemberDTMFInputEvent struct {
	Value DTMFInputEvent

	noSmithyDocumentSerde
}

func (*StartConversationRequestEventStreamMemberDTMFInputEvent) isStartConversationRequestEventStream() {
}

// Event sent from the client application to Amazon Lex V2 to indicate that it has
// finished playing audio and that Amazon Lex V2 should start listening for user
// input.
type StartConversationRequestEventStreamMemberPlaybackCompletionEvent struct {
	Value PlaybackCompletionEvent

	noSmithyDocumentSerde
}

func (*StartConversationRequestEventStreamMemberPlaybackCompletionEvent) isStartConversationRequestEventStream() {
}

// Text sent from your client application to Amazon Lex V2. Each TextInputEvent is
// processed individually.
type StartConversationRequestEventStreamMemberTextInputEvent struct {
	Value TextInputEvent

	noSmithyDocumentSerde
}

func (*StartConversationRequestEventStreamMemberTextInputEvent) isStartConversationRequestEventStream() {
}

// Represents a stream of events between Amazon Lex V2 and your application.
//
// The following types satisfy this interface:
//
//	StartConversationResponseEventStreamMemberAudioResponseEvent
//	StartConversationResponseEventStreamMemberHeartbeatEvent
//	StartConversationResponseEventStreamMemberIntentResultEvent
//	StartConversationResponseEventStreamMemberPlaybackInterruptionEvent
//	StartConversationResponseEventStreamMemberTextResponseEvent
//	StartConversationResponseEventStreamMemberTranscriptEvent
type StartConversationResponseEventStream interface {
	isStartConversationResponseEventStream()
}

// An event sent from Amazon Lex V2 to your client application containing audio to
// play to the user.
type StartConversationResponseEventStreamMemberAudioResponseEvent struct {
	Value AudioResponseEvent

	noSmithyDocumentSerde
}

func (*StartConversationResponseEventStreamMemberAudioResponseEvent) isStartConversationResponseEventStream() {
}

// Event that Amazon Lex V2 sends to indicate that the stream is still open between
// the client application and Amazon Lex V2
type StartConversationResponseEventStreamMemberHeartbeatEvent struct {
	Value HeartbeatEvent

	noSmithyDocumentSerde
}

func (*StartConversationResponseEventStreamMemberHeartbeatEvent) isStartConversationResponseEventStream() {
}

// Event sent from Amazon Lex V2 to the client application containing the current
// state of the conversation between the user and Amazon Lex V2.
type StartConversationResponseEventStreamMemberIntentResultEvent struct {
	Value IntentResultEvent

	noSmithyDocumentSerde
}

func (*StartConversationResponseEventStreamMemberIntentResultEvent) isStartConversationResponseEventStream() {
}

// Event sent from Amazon Lex V2 to indicate to the client application should stop
// playback of audio. For example, if the client is playing a prompt that asks for
// the user's telephone number, the user might start to say the phone number before
// the prompt is complete. Amazon Lex V2 sends this event to the client application
// to indicate that the user is responding and that Amazon Lex V2 is processing
// their input.
type StartConversationResponseEventStreamMemberPlaybackInterruptionEvent struct {
	Value PlaybackInterruptionEvent

	noSmithyDocumentSerde
}

func (*StartConversationResponseEventStreamMemberPlaybackInterruptionEvent) isStartConversationResponseEventStream() {
}

// The event sent from Amazon Lex V2 to your application with text to present to
// the user.
type StartConversationResponseEventStreamMemberTextResponseEvent struct {
	Value TextResponseEvent

	noSmithyDocumentSerde
}

func (*StartConversationResponseEventStreamMemberTextResponseEvent) isStartConversationResponseEventStream() {
}

// Event sent from Amazon Lex V2 to your client application that contains a
// transcript of voice audio.
type StartConversationResponseEventStreamMemberTranscriptEvent struct {
	Value TranscriptEvent

	noSmithyDocumentSerde
}

func (*StartConversationResponseEventStreamMemberTranscriptEvent) isStartConversationResponseEventStream() {
}

// The event sent from your client application to Amazon Lex V2 with text input
// from the user.
type TextInputEvent struct {

	// The text from the user. Amazon Lex V2 processes this as a complete statement.
	//
	// This member is required.
	Text *string

	// A timestamp set by the client of the date and time that the event was sent to
	// Amazon Lex V2.
	ClientTimestampMillis int64

	// A unique identifier that your application assigns to the event. You can use this
	// to identify events in logs.
	EventId *string

	noSmithyDocumentSerde
}

// The event sent from Amazon Lex V2 to your application with text to present to
// the user.
type TextResponseEvent struct {

	// A unique identifier of the event sent by Amazon Lex V2. The identifier is in the
	// form RESPONSE-N, where N is a number starting with one and incremented for each
	// event sent by Amazon Lex V2 in the current session.
	EventId *string

	// A list of messages to send to the user. Messages are ordered based on the order
	// that you returned the messages from your Lambda function or the order that the
	// messages are defined in the bot.
	Messages []Message

	noSmithyDocumentSerde
}

// Event sent from Amazon Lex V2 to your client application that contains a
// transcript of voice audio.
type TranscriptEvent struct {

	// A unique identifier of the event sent by Amazon Lex V2. The identifier is in the
	// form RESPONSE-N, where N is a number starting with one and incremented for each
	// event sent by Amazon Lex V2 in the current session.
	EventId *string

	// The transcript of the voice audio from the user.
	Transcript *string

	noSmithyDocumentSerde
}

// The value of a slot.
type Value struct {

	// The value that Amazon Lex V2 determines for the slot. The actual value depends
	// on the setting of the value selection strategy for the bot. You can choose to
	// use the value entered by the user, or you can have Amazon Lex V2 choose the
	// first value in the resolvedValues list.
	//
	// This member is required.
	InterpretedValue *string

	// The text of the utterance from the user that was entered for the slot.
	OriginalValue *string

	// A list of additional values that have been recognized for the slot.
	ResolvedValues []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isStartConversationRequestEventStream()  {}
func (*UnknownUnionMember) isStartConversationResponseEventStream() {}
