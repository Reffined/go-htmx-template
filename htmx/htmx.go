package htmx

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
	AfterOnLoad          Event = "htmx:afterOnLoad"
	AfterRequest         Event = "htmx:afterRequest"
	AfterSettle          Event = "htmx:afterSettle"
	AfterSwap            Event = "htmx:afterSwap"
	BeforeCleanupElement Event = "htmx:beforeCleanupElement"
	BeforeOnLoad         Event = "htmx:beforeOnLoad"
	BeforeRequest        Event = "htmx:beforeRequest"
	BeforeSend           Event = "htmx:beforeSend"
	BeforeSwap           Event = "htmx:beforeSwap"
	BeforeTransition     Event = "htmx:beforeTransition"
	ConfigRequest        Event = "htmx:configRequest"
	Confirm              Event = "htmx:confirm"
	Load                 Event = "htmx:load"
)

func On(event Event, js Js) {

}
