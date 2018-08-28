#ifndef _NEBULAS_NF_NVM_V8_LIB_EVENT_H_
#define _NEBULAS_NF_NVM_V8_LIB_EVENT_H_

#include <v8.h>

using namespace v8;

void NewNativeEventFunction(Isolate *isolate, Local<ObjectTemplate> globalTpl);
void EventTriggerCallback(const FunctionCallbackInfo<Value> &info);

#endif // _NEBULAS_NF_NVM_V8_LIB_EVENT_H_
