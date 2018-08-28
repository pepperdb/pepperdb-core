#ifndef _NEBULAS_NF_NVM_V8_LIB_LOG_CALLBACK_H_
#define _NEBULAS_NF_NVM_V8_LIB_LOG_CALLBACK_H_

#include "../engine.h"

#include <v8.h>

using namespace v8;

void NewNativeLogFunction(Isolate *isolate, Local<ObjectTemplate> globalTpl);
void LogCallback(const FunctionCallbackInfo<Value> &info);

#endif // _NEBULAS_NF_NVM_V8_LIB_LOG_CALLBACK_H_
