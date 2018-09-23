#include "event.h"
#include "../engine.h"
#include "global.h"
#include "instruction_counter.h"

static EventTriggerFunc TRIGGER = NULL;

void InitializeEvent(EventTriggerFunc trigger) { TRIGGER = trigger; }

void NewNativeEventFunction(Isolate *isolate, Local<ObjectTemplate> globalTpl) {
  globalTpl->Set(String::NewFromUtf8(isolate, "_native_event_trigger"),
                 FunctionTemplate::New(isolate, EventTriggerCallback),
                 static_cast<PropertyAttribute>(PropertyAttribute::DontDelete |
                                                PropertyAttribute::ReadOnly));
}

void EventTriggerCallback(const FunctionCallbackInfo<Value> &info) {
  Isolate *isolate = info.GetIsolate();
  Local<Context> context = isolate->GetCurrentContext();

  if (info.Length() < 2) {
    isolate->ThrowException(Exception::Error(
        String::NewFromUtf8(isolate, "_native_event_trigger: mssing params")));
    return;
  }

  Local<Value> topic = info[0];
  if (!topic->IsString()) {
    isolate->ThrowException(Exception::Error(String::NewFromUtf8(
        isolate, "_native_event_trigger: topic must be string")));
    return;
  }

  Local<Value> data = info[1];
  if (!data->IsString()) {
    isolate->ThrowException(Exception::Error(String::NewFromUtf8(
        isolate, "_native_event_trigger: data must be string")));
    return;
  }

  if (TRIGGER == NULL) {
    return;
  }

  V8Engine *e = GetV8EngineInstance(context);
  String::Utf8Value sTopic(topic);
  String::Utf8Value sData(data);

  size_t cnt = 0;
  TRIGGER(e, *sTopic, *sData, &cnt);

  // record event usage.
  IncrCounter(isolate, isolate->GetCurrentContext(), cnt);
}
