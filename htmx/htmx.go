package htmx

import "io"

const HxGet string = "hx-get"
const HxPost string = "hx-post"
const HxPut string = "hx-put"
const HxDelete string = "hx-delete"
const HxPushUrl string = "hx-push-url"
const HxSelect string = "hx-select"
const HxSelectOob string = "hx-select-oob"
const HxSwap string = "hx-swap"
const HxSwapOob string = "hx-swap-oob"
const HxTarget string = "hx-target"
const HxTrigger string = "hx-trigger"

// swap
const InnerHTML string = "innerHTML"
const OuterHTML string = "outerHTML"
const TextContent string = "textContent"
const Beforebegin string = "beforebegin"
const Afterbegin string = "afterbegin"
const Beforeend string = "beforeend"
const Afterend string = "afterend"
const None string = "none"

type Event string
type Js string

const (
	AfterOnLoad          Event = "hx-on:htmx:afterOnLoad"
	AfterRequest         Event = "hx-on:htmx:afterRequest"
	AfterSettle          Event = "hx-on:htmx:afterSettle"
	AfterSwap            Event = "hx-on:htmx:afterSwap"
	BeforeCleanupElement Event = "hx-on:htmx:beforeCleanupElement"
	BeforeOnLoad         Event = "hx-on:htmx:beforeOnLoad"
	BeforeRequest        Event = "hx-on:htmx:beforeRequest"
	BeforeSend           Event = "hx-on:htmx:beforeSend"
	BeforeSwap           Event = "hx-on:htmx:beforeSwap"
	BeforeTransition     Event = "hx-on:htmx:beforeTransition"
	ConfigRequest        Event = "hx-on:htmx:configRequest"
	Confirm              Event = "hx-on:htmx:confirm"
	Load                 Event = "hx-on:htmx:load"
	Click                Event = "hx-on:click"
)

type Callback struct {
	Event Event
	Code  Js
}

func On(event Event, reader io.Reader) *Callback {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return &Callback{
		Code:  Js(bytes),
		Event: event,
	}
}
